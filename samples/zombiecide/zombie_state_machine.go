package zombiecide

import (
	//"log"

	"fmt"
	"log"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

const (
	Walk   string = "walk"
	Idle          = "idle"
	Dead          = "dead"
	Attack        = "attack"
)

func NewZombieStateMachine() ZombieStateMachine {
	//Clean this up with private constructors...
	zs := zombieStateMachineData{}
	locator := lib.NewCurrentLocation(lib.NewPoint(600, 600), -4, -2, 5, 5)
	attack := newAttackState(locator)
	walk := newWalkState(locator)
	dead := newDeadState(locator)
	idle := newIdleState(locator)
	//Preeviously kept all states in sperarate shared struct and may go back to that after refactor.
	attack.walk = &walk
	idle.dead = &dead
	idle.walk = &walk
	idle.attack = &attack
	walk.attack = &attack
	walk.idle = &idle
	dead.walk = &walk

	zs.current = &walk
	zs.current.start()
	return &zs
}

// DoState is unique to current state but DoRender is comming.
type ZombieStateMachine interface {
	Execute(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent)
}

type zState interface {
	doState(mouseEvent bus.MouseEvent) zState
	doRender(layer bus.ImageLayer, mouseEvent bus.MouseEvent)
	start()
	stop()
}
type zombieStateMachineData struct {
	current zState
}

func (z *zombieStateMachineData) Execute(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) {
	layer := bus.NewImageLayer()
	drawEvent.AddImageLayer(layer)

	state := z.current.doState(mouseEvent)
	if z.current != state {
		z.current.stop()
		state.start()
		z.current = state
	}
	z.current.doRender(layer, mouseEvent)
}

type zStateData struct {
	name    string
	scale   int32
	started bool
	locator lib.Navigator
	lib.FrameTracker
}

func (z *zStateData) doRender(imageLayer bus.ImageLayer, mouseEvent bus.MouseEvent) {
	imageLayer.AddLayerData(render(z.name, z.scale, z.locator, z.FrameTracker, mouseEvent))
}

func (z *zStateData) start() {
	if !z.started {
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(getZombieAudio(z.name)).Play())
		z.started = true
	}
}
func (z *zStateData) stop() {
	if z.started {
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(getZombieAudio(z.name)).Stop())
		z.started = false
		z.FrameTracker.Reset()
	}
}

// Current scale of 30%
func newStateData(name string, locator lib.Navigator) zStateData {
	return zStateData{name, 30, false, locator, lib.NewFrameData()}
}

type zAttackData struct {
	zStateData
	walk zState
}

func newAttackState(locator lib.Navigator) zAttackData {
	zStateData := newStateData(Attack, locator)
	zStateData.FrameTracker.SetMaxFrame(7).SetRepeatFrame(3)
	attack := zAttackData{zStateData, nil}
	return attack
}

func (z *zAttackData) doState(mouseEvent bus.MouseEvent) zState {
	log.Default().Println(z.name)
	z.FrameTracker.Increment()

	if mouseEvent.LeftButton().IsDown() {
		return z
	} else {
		return z.walk
	}

}

type zIdleData struct {
	zStateData
	dead, walk, attack zState
}

func newIdleState(locator lib.Navigator) zIdleData {
	zStateData := newStateData(Idle, locator)
	zStateData.FrameTracker.SetMaxFrame(15).SetRepeatFrame(5).SetToLoop(false)
	idle := zIdleData{zStateData, nil, nil, nil}
	return idle
}

func (z *zIdleData) doState(mouseEvent bus.MouseEvent) zState {
	log.Default().Println(z.name)
	z.FrameTracker.Increment()
	if mouseEvent.LeftButton().IsDown() {
		return z.attack
	} else {
		point := z.locator.CalculateMove(mouseEvent)
		idleFrames := z.UpdateIdleFrames(point)
		if idleFrames == 0 {
			return z.walk
		} else if idleFrames >= 150 {
			return z.dead
		}
	}
	return z
}

type zWalkData struct {
	zStateData
	attack, idle zState
}

func newWalkState(locator lib.Navigator) zWalkData {
	zStateData := newStateData(Walk, locator)
	zStateData.FrameTracker.SetMaxFrame(10).SetRepeatFrame(3)
	walk := zWalkData{zStateData, nil, nil}
	return walk
}

func (z *zWalkData) doState(mouseEvent bus.MouseEvent) zState {
	log.Default().Println(z.name)
	z.FrameTracker.Increment()
	if mouseEvent.LeftButton().IsDown() {
		return z.attack
	} else {
		point := z.locator.CalculateMove(mouseEvent)
		if z.UpdateIdleFrames(point) >= 50 {
			return z.idle
		}

		z.locator.Move(point)
		return z
	}
}

type zDeadData struct {
	zStateData
	walk zState
}

func newDeadState(locator lib.Navigator) zDeadData {
	zStateData := newStateData(Dead, locator)
	zStateData.FrameTracker.SetMaxFrame(10).SetRepeatFrame(5).SetToLoop(false)
	dead := zDeadData{zStateData, nil}
	return dead
}

func (z *zDeadData) doState(mouseEvent bus.MouseEvent) zState {
	log.Default().Println(z.name)
	z.FrameTracker.Increment()
	point := z.locator.CalculateMove(mouseEvent)
	if z.UpdateIdleFrames(point) <= 0 {
		return z.walk
	}

	return z

}

func render(fileName string, scale int32, locator lib.Navigator, frameData lib.FrameTracker, mouseEvent bus.MouseEvent) bus.ImageMetadata {
	return createImageData(fileName, locator.GetX(), locator.GetY(), frameData.GetCurrentFrame(), scale, flipHorizontal(mouseEvent, locator))
}

func createImageData(fileName string, x, y, frameNumber, scale int32, flipHorizontal bool) bus.ImageMetadata {
	return bus.NewImageMetadata(getZombieImage(fileName, frameNumber), x, y, scale).SetFlipHorizontal(flipHorizontal)
}

func flipHorizontal(mouseEvent bus.MouseEvent, locator lib.Navigator) bool {
	return mouseEvent.GetX() < locator.GetX()
}

// Helper methods for the states...
func getZombieImage(name string, frameNumber int32) string {
	return fmt.Sprintf("samples/resources/zombiecide/%s (%d).png", name, frameNumber)
}

func getZombieAudio(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
