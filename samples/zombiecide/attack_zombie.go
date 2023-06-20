package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type attackZombie struct {
	spriteControllerData
	walkingZombie ZombieSprite
}

type AttackZombie interface {
	ZombieSprite
	SetWalkingZombie(zombie WalkingZombie)
}

func newAttackZombie() AttackZombie {
	return &attackZombie{newSpriteControllerData(7, 3, 200, 300, "attack"), nil}

}
func (s *attackZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite {
	var zReturn ZombieSprite = s
	if mouseEvent.LeftButton().IsDown() {
		s.doSendAudio()
		s.sendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
		s.incrementFrame()
		s.noLoop()
	} else {
		zReturn = s.doTransition(s.walkingZombie)
	}
	return zReturn
}

func (s *attackZombie) SetWalkingZombie(zombie WalkingZombie) {
	s.walkingZombie = zombie
}
