package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

type zombieData struct {
	SpriteData
	framesIdle int32
}

type ZombieSprite interface {
	Sprite
	RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite
	getIdleFrames() int32
	updateIdleCount(p Point) int32
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

// Probably a better factory pattern for this in idiomatic Golang
func NewZombie() ZombieSprite {

	walking := newWalkingZombie()
	dead := newDeadZombie()
	idle := newIdleZombie()
	attack := newAttackZombie()

	attack.SetWalkingZombie(walking)
	dead.SetWalkingZombie(walking)
	walking.SetAttackZombie(attack).SetIdleZombie(idle)
	idle.SetDeadZombie(dead).SetWalkingZombie(walking).SetAttackZombie(attack)

	//Start walking...
	return walking
}

func newZombieData(x, y, width, height int32, name string) zombieData {
	return zombieData{NewSpriteData(x, y, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name)), 0}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}
