package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

type zombieData struct {
	SpriteData
	framesIdle int32
	sprites    ZombieSprites
}

type ZombieSprite interface {
	Sprite
	GetState(mouseEvent bus.MouseEvent) ZombieSprite
	RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent)
	getIdleFrames() int32
	updateIdleCount(p Point) int32
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
func newZombieData(x, y, width, height int32, name string, sprites ZombieSprites) zombieData {
	return zombieData{NewSpriteData(x, y, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name)), 0, sprites}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}

func (s *zombieData) doTransition(nextState ZombieSprite) ZombieSprite {
	s.Stop()
	s.framesIdle = 0
	nextState.SetCurrentLocation(s.GetCurrentLocation())
	return nextState
}

func (s *zombieData) getIdleFrames() int32 {
	return s.framesIdle
}

func (s *zombieData) updateIdleCount(point Point) int32 {
	if point.GetY() == 0 && point.GetX() == 0 {
		s.framesIdle++
	} else {
		s.framesIdle = 0
	}
	return s.framesIdle
}
