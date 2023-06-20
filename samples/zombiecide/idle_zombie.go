package zombiecide

import (
	"log"

	"github.com/vorpalgame/vorpal/bus"
)

type idleZombie struct {
	spriteControllerData
	attackZombie  AttackZombie
	deadZombie    DeadZombie
	walkingZombie WalkingZombie
	framesIdle    int32
}

// TODO Some of the methods for setting shoudl be private to the package.
type IdleZombie interface {
	ZombieSprite
	SetWalkingZombie(zombie WalkingZombie) IdleZombie
	SetDeadZombie(zombie DeadZombie) IdleZombie
	SetAttackZombie(zombie AttackZombie) IdleZombie
}

func newIdleZombie() IdleZombie {
	return &idleZombie{newSpriteControllerData(15, 3, 200, 300, "idle"), nil, nil, nil, 0}
}

func (s *idleZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = doTransition(s, s.attackZombie)
	} else {
		doSendAudio(s)
		point := s.calculateMove(mouseEvent)
		s.framesIdle = doIdleCount(s.framesIdle, point)
		log.Default().Println(s.framesIdle)
		if s.framesIdle < 150 {
			s.currentLocation.Add(point)
			s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.incrementFrame()
			s.loop()
		} else {
			zReturn = doTransition(s, s.deadZombie)
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
