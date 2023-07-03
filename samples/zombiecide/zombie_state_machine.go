package zombiecide

import (
	//"log"

	"fmt"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

// Temporary until we get external configuration mechanism complete....
const (
	karen  = "samples/resources/zombiecide/karen/animation/%s%d.png"
	george = "samples/resources/zombiecide/george/animation/%s%d.png"
	albert = "samples/resources/zombiecide/albert/animation/%s%d.png"
	henry  = "samples/resources/zombiecide/henry/animation/%s%d.png"
)

// TODO States should come from yaml as some have more states than others....
// Would be great to have mapping of state transitions there as well.
const (
	Walk   = "Walk"
	Idle   = "Idle"
	Dead   = "Dead"
	Attack = "Attack"
)

// TODO use this to load yaml data...Ongoing work...

type ZombieConfigurationData struct {
	Name            string
	ImageTemplate   string
	ImageScale      int32
	CurrentLocation lib.Navigator
}

// TODO Karen, Henry and Albert have different number of frames and states.
// Temporary mechanism...
// TODO a configuration mechanism and struct for scale, directory, states, etc. is needed.
var animationSpec = henry
var currentScale = int32(30)

func NewZombieStateMachine() ZombieStateMachine {
	//Clean this up with private constructors...
	zs := newZombieStates()
	sm := &zombieStateMachineData{zs}

	locator := lib.NewNavigatorOffset(lib.NewPoint(600, 600), -4, -2, 5, 5)
	newAttackState(zs, locator)
	newDeadState(zs, locator)
	newIdleState(zs, locator)
	zs.setCurrent(newWalkState(zs, locator))
	zs.getCurrent().start()
	return sm
}

type ZombieStateMachine interface {
	ZombieStates
	Execute(drawEvent bus.DrawLayersEvent, mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent)
}

type zombieStateMachineData struct {
	ZombieStates
}

func (z *zombieStateMachineData) Execute(drawEvent bus.DrawLayersEvent, mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) {
	layer := bus.NewImageLayer()
	drawEvent.AddImageLayer(layer)
	current := z.getCurrent()
	state := current.doState(mouseEvent)
	if current != state {
		current.stop()
		state.start()
		z.setCurrent(state)
	}
	z.getCurrent().doRender(layer, mouseEvent)
}

// ================================================================================
// STATE COLLECTION
// ================================================================================
func newZombieStates() ZombieStates {
	return &zombieStatesData{nil, make(map[string]ZombieState)}
}

type ZombieStates interface {
	getCurrent() ZombieState
	setCurrent(currentState ZombieState) ZombieStates
	getState(stateName string) ZombieState
	addState(stateName string, state ZombieState) ZombieStates

	getAll() map[string]ZombieState
	//Convenience methods will probably go away in the future
	//in favor of using the getState(string) to make it a bit more flexible.
	getAttackZombie() ZombieState
	getDeadZombie() ZombieState
	getIdleZombie() ZombieState
	getWalkingZombie() ZombieState
}
type zombieStatesData struct {
	current  ZombieState
	stateMap map[string]ZombieState
}

func (zs *zombieStatesData) addState(state string, zombie ZombieState) ZombieStates {
	zs.stateMap[state] = zombie
	return zs
}
func (zs *zombieStatesData) setCurrent(zombie ZombieState) ZombieStates {
	zs.current = zombie
	return zs
}
func (zs *zombieStatesData) getCurrent() ZombieState {
	return zs.current
}
func (zs *zombieStatesData) getState(stateName string) ZombieState {
	return zs.stateMap[stateName]
}
func (zs *zombieStatesData) getDeadZombie() ZombieState {
	return zs.getState(Dead)
}
func (zs *zombieStatesData) getIdleZombie() ZombieState {
	return zs.getState(Idle)
}
func (zs *zombieStatesData) getWalkingZombie() ZombieState {
	return zs.getState(Walk)
}
func (zs *zombieStatesData) getAttackZombie() ZombieState {
	return zs.getState(Attack)
}
func (zs *zombieStatesData) getAll() map[string]ZombieState {
	return zs.stateMap
}

// ================================================================================
// General State
// ================================================================================

type ZombieState interface {
	doState(mouseEvent bus.MouseEvent) ZombieState
	doRender(layer bus.ImageLayer, mouseEvent bus.MouseEvent)
	start()
	stop()
}

type zStateData struct {
	name    string
	scale   int32
	started bool
	locator lib.Navigator
	lib.FrameTracker
	lib.AudioState
	ZombieStates
}

func (z *zStateData) doRender(imageLayer bus.ImageLayer, mouseEvent bus.MouseEvent) {
	imageLayer.AddLayerData(render(z.name, z.scale, z.locator, z.FrameTracker, mouseEvent))
}

func (z *zStateData) start() {
	if !z.started {
		bus.GetVorpalBus().SendAudioEvent(bus.NewPlayAudioEvent(z.AudioState))
		z.started = true
	}
}
func (z *zStateData) stop() {
	if z.started {
		bus.GetVorpalBus().SendAudioEvent(bus.NewStopAudioEvent(z.AudioState))
		z.started = false
		z.FrameTracker.Reset()
	}
}

// Current scale of 30%
func newStateData(name string, locator lib.Navigator, zs ZombieStates) zStateData {
	return zStateData{name, currentScale, false, locator, lib.NewFrameData(), lib.NewAudioState(getZombieAudio(name), false), zs}
}

// ================================================================================
// ATTACK
// ================================================================================
type zAttackData struct {
	zStateData
}

func newAttackState(states ZombieStates, locator lib.Navigator) ZombieState {
	zStateData := newStateData(Attack, locator, states)
	zStateData.FrameTracker.SetMaxFrame(6).SetRepeatFrame(5)
	state := zAttackData{zStateData}
	states.addState(zStateData.name, &state)
	return &state
}

func (z *zAttackData) doState(mouseEvent bus.MouseEvent) ZombieState {

	z.FrameTracker.Increment()
	if mouseEvent.LeftButton().IsDown() {
		return z
	} else {
		return z.getWalkingZombie()
	}

}

// ================================================================================
// IDLE
// ================================================================================
type zIdleData struct {
	zStateData
}

func newIdleState(states ZombieStates, locator lib.Navigator) ZombieState {
	zStateData := newStateData(Idle, locator, states)
	zStateData.FrameTracker.SetMaxFrame(6).SetRepeatFrame(15).SetToLoop(false)
	state := zIdleData{zStateData}
	states.addState(zStateData.name, &state)
	return &state
}

func (z *zIdleData) doState(mouseEvent bus.MouseEvent) ZombieState {

	z.FrameTracker.Increment()
	if mouseEvent.LeftButton().IsDown() {
		return z.getAttackZombie()
	} else {
		point := z.locator.CalculateMove(mouseEvent.GetCursorPoint())
		idleFrames := z.UpdateIdleFrames(point)
		if idleFrames == 0 {
			return z.getWalkingZombie()
		} else if idleFrames >= 150 {
			return z.getDeadZombie()
		}
	}
	return z
}

// ================================================================================
// WALK
// ================================================================================
type zWalkData struct {
	zStateData
}

func newWalkState(states ZombieStates, locator lib.Navigator) ZombieState {
	zStateData := newStateData(Walk, locator, states)
	zStateData.FrameTracker.SetMaxFrame(10).SetRepeatFrame(3)
	state := zWalkData{zStateData}
	states.addState(zStateData.name, &state)
	return &state
}

func (z *zWalkData) doState(mouseEvent bus.MouseEvent) ZombieState {

	z.FrameTracker.Increment()
	if mouseEvent.LeftButton().IsDown() {
		return z.getAttackZombie()
	} else {
		point := z.locator.CalculateMove(mouseEvent.GetCursorPoint())
		if z.UpdateIdleFrames(point) >= 50 {
			return z.getIdleZombie()
		}

		z.locator.Move(point)
		return z
	}
}

// ================================================================================
// DEAD
// ================================================================================
type zDeadData struct {
	zStateData
}

func newDeadState(states ZombieStates, locator lib.Navigator) ZombieState {
	zStateData := newStateData(Dead, locator, states)
	zStateData.FrameTracker.SetMaxFrame(10).SetRepeatFrame(5).SetToLoop(false)
	state := zDeadData{zStateData}
	states.addState(zStateData.name, &state)
	return &state
}

func (z *zDeadData) doState(mouseEvent bus.MouseEvent) ZombieState {

	z.FrameTracker.Increment()
	point := z.locator.CalculateMove(mouseEvent.GetCursorPoint())
	if z.UpdateIdleFrames(point) <= 0 {
		return z.getWalkingZombie()
	}

	return z

}

// ================================================================================
// Helper functions...
// ================================================================================
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

	return fmt.Sprintf(animationSpec, name, frameNumber)
}

func getZombieAudio(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
