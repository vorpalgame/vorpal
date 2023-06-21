package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type idleZombie struct {
	zombieData
	attackZombie  AttackZombie
	deadZombie    DeadZombie
	walkingZombie WalkingZombie
}

// TODO Some of the methods for setting shoudl be private to the package.
type IdleZombie interface {
	ZombieSprite
	SetWalkingZombie(zombie WalkingZombie) IdleZombie
	SetDeadZombie(zombie DeadZombie) IdleZombie
	SetAttackZombie(zombie AttackZombie) IdleZombie
}

func newIdleZombie() IdleZombie {
	return &idleZombie{newZombieData(15, 3, 200, 300, "idle"), nil, nil, nil}
}

func (s *idleZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = s.doTransition(s.attackZombie)
	} else {
		s.DoSendAudio()
		point := s.calculateMove(mouseEvent)

		if s.updateIdleCount(point) < 250 {
			s.currentLocation.Add(point)
			s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.IncrementFrame()
			s.Loop()
		} else {
			zReturn = s.doTransition(s.deadZombie)
			s.framesIdle = 0
		}

	}
	return zReturn
}

// SetDeadZombie implements IdleZombie.
func (z *idleZombie) SetDeadZombie(zombie DeadZombie) IdleZombie {
	z.deadZombie = zombie
	return z
}

// SetWalkingZombie implements IdleZombie.
func (z *idleZombie) SetWalkingZombie(zombie WalkingZombie) IdleZombie {
	z.walkingZombie = zombie
	return z
}

// setAttackZombie implements IdleZombie.
func (z *idleZombie) SetAttackZombie(zombie AttackZombie) IdleZombie {
	z.attackZombie = zombie
	return z
}
