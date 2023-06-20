package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type deadZombie struct {
	zombieData
	walkingZombie ZombieSprite
}

//TODO Some of the methods for setting shoudl be private to the package.
type DeadZombie interface {
	ZombieSprite
	SetWalkingZombie(zombie WalkingZombie) DeadZombie
}

func newDeadZombie() DeadZombie {
	return &deadZombie{newZombieData(12, 3, 300, 300, "dead"), nil}
}

func (s *deadZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	s.doSendAudio()
	point := s.calculateMove(mouseEvent)

	if s.updateIdleCount(point) > 0 {
		s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
		s.incrementFrame()
		s.noLoop()
	} else {
		zReturn = s.doTransition(s.walkingZombie)
	}
	return zReturn
}

func (s *deadZombie) SetWalkingZombie(zombie WalkingZombie) DeadZombie {
	s.walkingZombie = zombie
	return s
}
