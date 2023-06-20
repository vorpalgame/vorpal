package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO Now that prototyping PoC is done we need unit tests
// and to move to TDD.
func TestWalkingZombie(t *testing.T) {
	z := newWalkingZombie()
	assert.NotNil(t, z, "Walking zombie not constructed properly")
}
func TestIdleZombie(t *testing.T) {
	z := newIdleZombie()
	assert.NotNil(t, z, "Idle zombie not constructed properly")
}

func TestDeadZombie(t *testing.T) {
	z := newDeadZombie()
	assert.NotNil(t, z, "Dead zombie not constructed properly")
}

func TestAttackZombie(t *testing.T) {
	z := newAttackZombie()
	assert.NotNil(t, z, "Attack zombie not constructed properly")
}
