package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type idleZombie struct {
	zombieData
}

type IdleZombie interface {
	ZombieSprite
}

func newIdleZombie(sprites ZombieSprites) IdleZombie {
	return &idleZombie{newZombieData(15, 3, 200, 300, "idle", sprites)}
}

func (s *idleZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = s.doTransition(s.sprites.GetAttackZombie())
	} else {
		s.DoSendAudio()
		point := s.calculateMove(mouseEvent)

		if s.updateIdleCount(point) < 250 {
			s.currentLocation.Add(point)
			s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.IncrementFrame()
			s.Loop()
		} else {
			zReturn = s.doTransition(s.sprites.GetDeadZombie())
			s.framesIdle = 0
		}

	}
	return zReturn
}
