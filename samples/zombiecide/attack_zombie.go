package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type attackZombie struct {
	zombieData
}

type AttackZombie interface {
	ZombieSprite
}

func newAttackZombie(sprites ZombieSprites) AttackZombie {
	return &attackZombie{newZombieData(7, 3, 200, 300, "attack", sprites)}

}
func (s *attackZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		s.DoSendAudio()
		s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
		s.IncrementFrame()
		s.NoLoop()
	} else {
		zReturn = s.doTransition(s.sprites.GetWalkingZombie())
	}
	return zReturn
}
