package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type deadZombie struct {
	zombieData
}

//TODO Some of the methods for setting shoudl be private to the package.
type DeadZombie interface {
	ZombieSprite
}

func newDeadZombie(sprites ZombieSprites) DeadZombie {
	return &deadZombie{newZombieData(12, 3, 300, 300, "dead", sprites)}
}

func (s *deadZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	s.DoSendAudio()
	point := s.calculateMove(mouseEvent)

	if s.updateIdleCount(point) > 0 {
		s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
		s.IncrementFrame()
		s.NoLoop()
	} else {
		zReturn = s.doTransition(s.sprites.GetWalkingZombie())
	}
	return zReturn
}
