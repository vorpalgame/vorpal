package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

// Tests are about state trnasitions...
// Some configuration data like idle should be settable for testing.
// More  worked out with JSON configuration files in the future.
// Revamping unit test to match new structures...
type mockData struct {
	ZombieStateMachine
	mouseEvent      bus.MouseEvent
	currentLocation lib.Point
}

func mockZombieData() mockData {
	mockData := mockData{NewZombieStateMachine(), createMouseEvent(0, 0), lib.NewPoint(600, 600)}
	//log.Default().Println("Start test...")
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
	z := data.getCurrent()

	assertWalkingState(t, z)
	mouseEvent := createMouseEvent(data.currentLocation.GetX(), data.currentLocation.GetY())

	//Sit idle for enough frames without triggering transition.
	z = doNoMouseMove(mouseEvent, 49, z)
	assertWalkingState(t, z)

	assert.NotNil(t, data.getIdleZombie(), "Mock not available for idle transition")
	//Now tranasition
	z = z.doState(mouseEvent)
	assertIdleState(t, z)

}

func TestToAttackZombie(t *testing.T) {

	data := mockZombieData()

	verifyAttackTransition(data.getWalkingZombie(), t)
	verifyAttackTransition(data.getIdleZombie(), t)

}

func TestIdleToDeadZombie(t *testing.T) {
	data := mockZombieData()
	z := data.getIdleZombie()
	assertIdleState(t, z)
	mouseEvent := createMouseEvent(data.currentLocation.GetX(), data.currentLocation.GetY())
	z = doNoMouseMove(mouseEvent, 149, z)
	assertIdleState(t, z)

	//Transition
	z = z.doState(mouseEvent)
	assertDeadState(t, z)

}

func TestDeadZombie(t *testing.T) {
	data := mockZombieData()
	z := data.getDeadZombie()
	assert.NotNil(t, z, "Dead zombie not constructed properly")
}

func TestAttackZombie(t *testing.T) {
	data := mockZombieData()
	z := data.getAttackZombie()
	assert.NotNil(t, z, "Attack zombie not constructed properly")
}

func verifyAttackTransition(currentState ZombieState, t *testing.T) {
	mouseEvent := createLeftDownMouseEvent()
	nextState := currentState.doState(mouseEvent)
	assertAttackState(t, nextState)

}
func doNoMouseMove(mouseEvent bus.MouseEvent, count int, currentState ZombieState) ZombieState {
	for i := 0; i < count; i++ {
		currentState = currentState.doState(mouseEvent)
	}
	return currentState
}

// // Coudld modify to pass error string from concreete assert types..
func assertAttackState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &zAttackData{})
}
func assertWalkingState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &zWalkData{})
}
func assertIdleState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &zIdleData{})
}
func assertDeadState(t *testing.T, actual ZombieState) {
	assertType(t, actual, &zDeadData{})
}

func assertType(t *testing.T, actual ZombieState, expected ZombieState) {
	assert.NotNil(t, actual, "Failed to get state.")
	assert.IsType(t, expected, actual, "Didn't get expected type back")
}
