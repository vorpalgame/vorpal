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

func (s *idleZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	if mouseEvent.LeftButton().IsDown() {
		return s.doTransition(s.sprites.GetAttackZombie())
	} else {
		point := s.calculateMove(mouseEvent)

		if s.updateIdleCount(point) < 250 {
			s.currentLocation.Add(point)
			return s
		} else {
			return s.doTransition(s.sprites.GetDeadZombie())
		}

	}
}

func (s *idleZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {

	s.DoSendAudio()
	s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
	s.IncrementFrame()
	s.Loop()

}
