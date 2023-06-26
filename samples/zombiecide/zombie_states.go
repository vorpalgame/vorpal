package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

// TODO Would be better with a type keyed map but
// that appears problematic in Golang...TBD...

type ZombieState interface {
	lib.Sprite
	GetState(mouseEvent bus.MouseEvent, states ZombieStates) ZombieState
}

type zombieStateData struct {
	walking WalkingZombie
	dead    DeadZombie
	idle    IdleZombie
	attack  AttackZombie
}
type ZombieStates interface {
	GetAttackZombie() ZombieState
	GetDeadZombie() ZombieState
	GetIdleZombie() ZombieState
	GetWalkingZombie() ZombieState
}

func NewZombieSprites() ZombieStates {
	var sprites = zombieStateData{}

	sprites.dead = newDeadZombie()
	sprites.idle = newIdleZombie()
	sprites.attack = newAttackZombie()
	sprites.walking = newWalkingZombie()
	return sprites
}

func (zs zombieStateData) GetAttackZombie() ZombieState {
	return zs.attack
}
func (zs zombieStateData) GetDeadZombie() ZombieState {
	return zs.dead
}
func (zs zombieStateData) GetIdleZombie() ZombieState {
	return zs.idle
}
func (zs zombieStateData) GetWalkingZombie() ZombieState {
	return zs.walking
}

// Helper methods for the states...
func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
