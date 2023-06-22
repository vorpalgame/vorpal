package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

// TODO More general purpose functionality and duplicates can be pushed
// to sprite. TBD
type zombieData struct {
	SpriteData
	sprites ZombieSprites
}

type ZombieSprite interface {
	Sprite
	GetState(mouseEvent bus.MouseEvent) ZombieSprite
	Transition(sprite ZombieSprite) ZombieSprite
	GetSprites() ZombieSprites //Mainly for testing...
}

func (s *zombieData) Transition(nextState ZombieSprite) ZombieSprite {
	s.Stop()
	nextState.SetCurrentLocation(s.GetCurrentLocation())
	nextState.Init()
	return nextState
}
func (s *zombieData) GetSprites() ZombieSprites {
	return s.sprites
}

// TODO Would be better with a type keyed map but
// that appears problematic in Golang...TBD...
type zombieSprites struct {
	walking WalkingZombie
	dead    DeadZombie
	idle    IdleZombie
	attack  AttackZombie
}
type ZombieSprites interface {
	GetAttackZombie() ZombieSprite
	GetDeadZombie() ZombieSprite
	GetIdleZombie() ZombieSprite
	GetWalkingZombie() ZombieSprite
}

func NewZombieSprites() ZombieSprites {
	var sprites = zombieSprites{}
	sprites.walking = newWalkingZombie(&sprites)
	sprites.dead = newDeadZombie(&sprites)
	sprites.idle = newIdleZombie(&sprites)
	sprites.attack = newAttackZombie(&sprites)
	return sprites
}

// Probably a better factory pattern for this in idiomatic Golang
func NewZombie() ZombieSprite {
	return NewZombieSprites().GetWalkingZombie()
}

func (zs zombieSprites) GetAttackZombie() ZombieSprite {
	return zs.attack
}
func (zs zombieSprites) GetDeadZombie() ZombieSprite {
	return zs.dead
}
func (zs zombieSprites) GetIdleZombie() ZombieSprite {
	return zs.idle
}
func (zs zombieSprites) GetWalkingZombie() ZombieSprite {
	return zs.walking
}
func newZombieData(maxFrames, repeatPerFrame, width, height int32, name string, sprites ZombieSprites) zombieData {
	return zombieData{NewSpriteData(maxFrames, repeatPerFrame, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name)), sprites}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
