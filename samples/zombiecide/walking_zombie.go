package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type walkingZombie struct {
	zombieData
	idleZombie   ZombieSprite
	attackZombie ZombieSprite
}

type WalkingZombie interface {
	ZombieSprite
	SetAttackZombie(attack AttackZombie) WalkingZombie
	SetIdleZombie(idle ZombieSprite) WalkingZombie
}

func newWalkingZombie() WalkingZombie {
	return &walkingZombie{newZombieData(10, 3, 200, 300, "walk"), nil, nil}
}

func (s *walkingZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = s.doTransition(s.attackZombie)
	} else {
		s.DoSendAudio()
		point := s.calculateMove(mouseEvent)

		if s.updateIdleCount(point) < 50 {
			s.currentLocation.Add(point)
			s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.IncrementFrame()
			s.Loop()
		} else {
			zReturn = s.doTransition(s.idleZombie)
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
