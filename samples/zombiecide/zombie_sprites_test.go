package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sprites ZombieSprites

func mockZombieData() {
	data := zombieSprites{}
	data.walking = newWalkingZombie(sprites)
	data.idle = newIdleZombie(sprites)
	data.attack = newAttackZombie(sprites)
	data.dead = newDeadZombie(sprites)
	sprites = data
}

func TestWalkingZombie(t *testing.T) {
	mockZombieData()
	z := sprites.GetWalkingZombie()
	assert.NotNil(t, z, "Walking zombie not constructed properly")
}
func TestIdleZombie(t *testing.T) {
	mockZombieData()
	z := sprites.GetIdleZombie()
	assert.NotNil(t, z, "Idle zombie not constructed properly")
}

func TestDeadZombie(t *testing.T) {
	mockZombieData()
	z := sprites.GetDeadZombie()
	assert.NotNil(t, z, "Dead zombie not constructed properly")
}

func TestAttackZombie(t *testing.T) {
	mockZombieData()
	z := sprites.GetAttackZombie()
	assert.NotNil(t, z, "Attack zombie not constructed properly")
}
