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
	RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent)
	Transition(sprite ZombieSprite) ZombieSprite
}

func (s *zombieData) Transition(nextState ZombieSprite) ZombieSprite {
	s.Stop()
	nextState.SetCurrentLocation(s.GetCurrentLocation())
	nextState.Init()
	return nextState
}

// TODO Would be better with a type keyed map but
// that appears problematic in Golang...TBD...
type zombieSprites struct {
	walkingZombie WalkingZombie
	deadZombie    DeadZombie
	idleZombie    IdleZombie
	attackZombie  AttackZombie
}
type ZombieSprites interface {
	GetAttackZombie() ZombieSprite
	GetDeadZombie() ZombieSprite
	GetIdleZombie() ZombieSprite
	GetWalkingZombie() ZombieSprite
}

// Probably a better factory pattern for this in idiomatic Golang
func NewZombie() ZombieSprite {

	var sprites = zombieSprites{}
	sprites.walkingZombie = newWalkingZombie(&sprites)
	sprites.deadZombie = newDeadZombie(&sprites)
	sprites.idleZombie = newIdleZombie(&sprites)
	sprites.attackZombie = newAttackZombie(&sprites)

	//Start walking...
	return sprites.walkingZombie
}

func (zs zombieSprites) GetAttackZombie() ZombieSprite {
	return zs.attackZombie
}

func (zs zombieSprites) GetDeadZombie() ZombieSprite {
	return zs.deadZombie
}
func (zs zombieSprites) GetIdleZombie() ZombieSprite {
	return zs.idleZombie
}
func (zs zombieSprites) GetWalkingZombie() ZombieSprite {
	return zs.walkingZombie
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
