package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type walkingZombie struct {
	zombieData
}

type WalkingZombie interface {
	ZombieSprite
}

func newWalkingZombie(sprites ZombieSprites) WalkingZombie {
	return &walkingZombie{newZombieData(10, 3, 200, 300, "walk", sprites)}
}

func (s *walkingZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		zReturn = s.doTransition(s.sprites.GetAttackZombie())
	} else {
		s.DoSendAudio()
		point := s.calculateMove(mouseEvent)

		if s.updateIdleCount(point) < 50 {
			s.currentLocation.Add(point)
			s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
			s.IncrementFrame()
			s.Loop()
		} else {
			zReturn = s.doTransition(s.sprites.GetIdleZombie())
		}
	}
	return zReturn
}
