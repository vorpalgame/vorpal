package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
)

// Tests are about state trnasitions...
// Some configuration data like idle should be settable for testing.
// More  worked out with JSON configuration files in the future.
type mockData struct {
	sprites    ZombieSprites
	mouseEvent bus.MouseEvent
}

func mockZombieData() mockData {
	mockData := mockData{}
	mockData.sprites = NewZombieSprites()
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
	z := data.sprites.GetWalkingZombie()
	assertWalkingState(t, z)
	mouseEvent := createMouseEvent(z.GetCurrentLocation().GetX(), z.GetCurrentLocation().GetY())

	//Sit idle for enough frames without triggering transition.
	z = doNoMouseMove(mouseEvent, 49, z)
	assert.IsType(t, &walkingZombie{}, z, "Should still be in walking state")

	assert.True(t, z.GetFrameData().GetIdleFrames() == 49)
	assert.NotNil(t, data.sprites.GetIdleZombie(), "Mock not available for idle transition")
	//Now tranasition
	z = z.GetState(mouseEvent)
	assertIdleState(t, z)

}

func TestToAttackZombie(t *testing.T) {

	data := mockZombieData()

	verifyAttackTransition(data.sprites.GetWalkingZombie(), t)
	verifyAttackTransition(data.sprites.GetIdleZombie(), t)

}

func TestIdleToDeadZombie(t *testing.T) {
	data := mockZombieData()
	z := data.sprites.GetIdleZombie()
	assertIdleState(t, z)
	mouseEvent := createMouseEvent(z.GetCurrentLocation().GetX(), z.GetCurrentLocation().GetY())
	z = doNoMouseMove(mouseEvent, 249, z)
	assertIdleState(t, z)

	//Transition
	z = z.GetState(mouseEvent)
	assertDeadState(t, z)

}

func TestDeadZombie(t *testing.T) {
	data := mockZombieData()
	z := data.sprites.GetDeadZombie()
	assert.NotNil(t, z, "Dead zombie not constructed properly")
}

func TestAttackZombie(t *testing.T) {
	data := mockZombieData()
	z := data.sprites.GetAttackZombie()
	assert.NotNil(t, z, "Attack zombie not constructed properly")
}

func verifyAttackTransition(currentState ZombieSprite, t *testing.T) {
	mouseEvent := createLeftDownMouseEvent()
	nextState := currentState.GetState(mouseEvent)
	assert.NotNil(t, nextState, "Failed to get state.")
	assert.IsType(t, &attackZombie{}, nextState, "Didn't get attack type back")

}
func doNoMouseMove(mouseEvent bus.MouseEvent, count int, currentState ZombieSprite) ZombieSprite {
	for i := 0; i < count; i++ {
		currentState = currentState.GetState(mouseEvent)
	}
	return currentState
}

// Coudld modify to pass error string from concreete assert types..
func assertAttackState(t *testing.T, actual ZombieSprite) {
	assertType(t, actual, &attackZombie{})
}
func assertWalkingState(t *testing.T, actual ZombieSprite) {
	assertType(t, actual, &walkingZombie{})
}
func assertIdleState(t *testing.T, actual ZombieSprite) {
	assertType(t, actual, &idleZombie{})
}
func assertDeadState(t *testing.T, actual ZombieSprite) {
	assertType(t, actual, &deadZombie{})
}

func assertType(t *testing.T, actual ZombieSprite, expected ZombieSprite) {
	assert.NotNil(t, actual, "Failed to get state.")
	assert.IsType(t, expected, actual, "Didn't get attack type back")
}
