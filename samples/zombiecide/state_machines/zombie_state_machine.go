package state_machines

import (
	"fmt"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"log"
)

///////////////////////////////////////////////////////////////////////////////////////
/////      Zombie Data and Functions

type Zombie interface {
	Execute(drawEvent bus.DrawLayersEvent, mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent)
	getNavigator() lib.Navigator
}

func (z *ZombieData) Execute(drawEvent bus.DrawLayersEvent, mouseEvent bus.MouseEvent, keyEvent bus.KeyEvent) {
	state := z.StateMap[z.CurrentStateName]

	tx := NewBehaviorTransaction(z.CurrentStateName, mouseEvent, z.getNavigator(), state.getFrameTracker())
	state.getBehaviors().ExecuteBehaviors(tx)
	state.render(drawEvent, mouseEvent, z.getNavigator())
	if tx.IsTransition() {
		z.CurrentStateName = tx.GetNextStateName()
	}
}

func (z *ZombieData) getNavigator() lib.Navigator {
	return z.Navigator
}

///////////////////////////////////////////////////////////////////////////////////////
/////          State Data and Functions

type ZombieState interface {
	getBehaviors() BehaviorsExecutor
	getFrameTracker() lib.FrameTracker
	render(drawEvent bus.DrawLayersEvent, mouseEvent bus.MouseEvent, navigator lib.Navigator)
}

func (z *ZombieStateData) getFrameTracker() lib.FrameTracker {
	log.Default().Println(z)
	return z.FrameTracker
}

func (z *ZombieStateData) getBehaviors() BehaviorsExecutor {
	if z.behaviorsExecutor == nil {
		z.behaviorsExecutor = NewBehaviorExecutor(z.BehaviorNames)
	}
	return z.behaviorsExecutor
}

func (z *ZombieStateData) render(drawEvent bus.DrawLayersEvent, mouseEvent bus.MouseEvent, navigator lib.Navigator) {
	img := getZombieImage(z.Spec, z.Name, z.FrameTracker.GetCurrentFrame())
	metadata := bus.NewImageMetadata(img, navigator.GetX(), navigator.GetY(), z.Scale).SetFlipHorizontal(flipHorizontal(mouseEvent, navigator))
	layer := bus.NewImageLayer()
	layer.AddLayerData(metadata)
	drawEvent.AddImageLayer(layer)
}

// Helper functions...
func flipHorizontal(mouseEvent bus.MouseEvent, locator lib.Navigator) bool {
	return mouseEvent.GetX() < locator.GetX()
}
func getZombieImage(spec, name string, frameNumber int32) string {
	return fmt.Sprintf(spec, name, frameNumber)
}

func getZombieAudio(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
