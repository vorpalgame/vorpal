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
	zombie := &attackZombie{NewZombieData(7, 3, "attack", sprites)}
	zombie.SetToLoop(true)
	return zombie
}

func (currentZombie *attackZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie
	} else {
		return currentZombie.sprites.GetWalkingZombie()
	}
}
