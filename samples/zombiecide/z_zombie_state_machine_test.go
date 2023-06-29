package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
)

// States wont' be determeined externally like this test. This is just to get the
// basic set up working.
func TestZombieStateMachine(t *testing.T) {
	z := NewZombieStateMachine()
	assert.NotNil(t, z)
	//This shoudl be automated...
	bus.InitKeys()
	var mouseEvent bus.MouseEvent = bus.NewMouseEvent(nil, nil, nil, 800, 800)
	walkEvent := bus.NewKeyEvent(bus.GetKeyByString("W"))
	idleEvent := bus.NewKeyEvent(bus.GetKeyByString("I"))
	deadEvent := bus.NewKeyEvent(bus.GetKeyByString("D"))
	attackEvent := bus.NewKeyEvent(bus.GetKeyByString("A"))

	//Start in walk so we can transtion to attack or idle
	z.DoState(mouseEvent, attackEvent)
	//From attack we can only go back to walk so could match on anything.
	z.DoState(mouseEvent, walkEvent)
	//Now we'll go to idle
	z.DoState(mouseEvent, idleEvent)
	//From idle we can die or walk.
	z.DoState(mouseEvent, deadEvent)
	//You can only start walking again..
	z.DoState(mouseEvent, walkEvent)

}
