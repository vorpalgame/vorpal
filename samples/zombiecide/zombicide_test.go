package zombiecide

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
)

// Tests are about state trnasitions...
// Some configuration data like idle should be settable for testing.
// More  worked out with JSON configuration files in the future.
type mockData struct {
	zombieStates ZombieStates
	mouseEvent   bus.MouseEvent
}

func mockZombieData() mockData {
	mockData := mockData{}
	mockData.zombieStates = NewZombieStates(50) //TODO Sccale needs to be removed from constructor.
	mockData.mouseEvent = createMouseEvent(0, 0)
	return mockData
}
func createLeftDownMouseEvent() bus.MouseEvent {
	return bus.NewMouseEvent(bus.NewMouseButtonState("left", true), bus.NewMouseButtonState("right", false), bus.NewMouseButtonState("middle", false), 0, 0)

}
func createMouseEvent(x, y int32) bus.MouseEvent {
	return bus.NewMouseEvent(bus.NewMouseButtonState("left", false), bus.NewMouseButtonState("right", false), bus.NewMouseButtonState("middle", false), x, y)

}

func TestWalkingToIdleZombie(t *testing.T) {
	data := mockZombieData()
	z := data.zombieStates.GetCurrent()
	log.Default().Println(data.zombieStates.GetCurrent())
	assertWalkingState(t, z)
	mouseEvent := createMouseEvent(z.GetCurrentLocation().GetX(), z.GetCurrentLocation().GetY())

	//Sit idle for enough frames without triggering transition.
	z = doNoMouseMove(mouseEvent, 49, z)
	assertWalkingState(t, z)

	//assert.True(t, z.GetIdleFrames() == 49)
	assert.NotNil(t, data.zombieStates.GetIdleZombie(), "Mock not available for idle transition")
	//Now tranasition
	z = z.GetState(mouseEvent)
	assertIdleState(t, z)

}

func TestToAttackZombie(t *testing.T) {

	data := mockZombieData()

	verifyAttackTransition(data.zombieStates.GetWalkingZombie(), t)
	verifyAttackTransition(data.zombieStates.GetIdleZombie(), t)

}

func TestIdleToDeadZombie(t *testing.T) {
	data := mockZombieData()
	z := data.zombieStates.GetIdleZombie()
	assertIdleState(t, z)
	mouseEvent := createMouseEvent(z.GetCurrentLocation().GetX(), z.GetCurrentLocation().GetY())
	z = doNoMouseMove(mouseEvent, 149, z)
	assertIdleState(t, z)

	//Transition
	z = z.GetState(mouseEvent)
	assertDeadState(t, z)

}

func TestDeadZombie(t *testing.T) {
	data := mockZombieData()
	z := data.zombieStates.GetDeadZombie()
	assert.NotNil(t, z, "Dead zombie not constructed properly")
}

func TestAttackZombie(t *testing.T) {
	data := mockZombieData()
	z := data.zombieStates.GetAttackZombie()
	assert.NotNil(t, z, "Attack zombie not constructed properly")
}

func verifyAttackTransition(currentState ZombieState, t *testing.T) {
	mouseEvent := createLeftDownMouseEvent()
	nextState := currentState.GetState(mouseEvent)
	assertAttackState(t, nextState)

}
func doNoMouseMove(mouseEvent bus.MouseEvent, count int, currentState ZombieState) ZombieState {
	for i := 0; i < count; i++ {
		currentState = currentState.GetState(mouseEvent)
	}
	return currentState
}

// Coudld modify to pass error string from concreete assert types..
func assertAttackState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &attackZombie{})
}
func assertWalkingState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &walkingZombie{})
}
func assertIdleState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &idleZombie{})
}
func assertDeadState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &deadZombie{})
}

func assertType(t *testing.T, actual ZombieState, expected ZombieState) {
	assert.NotNil(t, actual, "Failed to get state.")
	assert.IsType(t, expected, actual, "Didn't get expected type back")
}
