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

func (currentZombie *attackZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie
	} else {
		return currentZombie.doTransition(currentZombie.sprites.GetWalkingZombie())
	}
}

func (currentZombie *attackZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {

	currentZombie.DoSendAudio()
	currentZombie.SendDrawEvent(drawEvent, currentZombie.currentLocation, currentZombie.flipHorizontal(mouseEvent))
	currentZombie.IncrementFrame()
	currentZombie.NoLoop()

}
