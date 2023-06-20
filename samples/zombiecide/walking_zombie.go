package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type walkingZombie struct {
	spriteControllerData
	idleZombie   ZombieSprite
	attackZombie ZombieSprite
	framesIdle   int32
}

type WalkingZombie interface {
	ZombieSprite
	SetAttackZombie(attack AttackZombie) WalkingZombie
	SetIdleZombie(idle ZombieSprite) WalkingZombie
}

func newWalkingZombie() WalkingZombie {
	return &walkingZombie{newSpriteControllerData(10, 3, 200, 300, "walk"), nil, nil, 0}
}

func (s *walkingZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = s.doTransition(s.attackZombie)
	} else {
		s.doSendAudio()
		point := s.calculateMove(mouseEvent)
		s.framesIdle = doIdleCount(s.framesIdle, point)

		if s.framesIdle < 50 {
			s.currentLocation.Add(point)
			s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.incrementFrame()
			s.loop()
		} else {
			zReturn = s.doTransition(s.idleZombie)
			s.framesIdle = 0
		}
	}
	return zReturn
}

func (s *walkingZombie) SetAttackZombie(zombie AttackZombie) WalkingZombie {
	s.attackZombie = zombie
	return s
}

func (s *walkingZombie) SetIdleZombie(idle ZombieSprite) WalkingZombie {
	s.idleZombie = idle
	return s
}
