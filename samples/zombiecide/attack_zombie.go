package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type attackZombie struct {
	zombieStateData
}

type AttackZombie interface {
	lib.Sprite
	ZombieState
}

func (currentZombie *attackZombie) GetState(mouseEvent bus.MouseEvent) ZombieState {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie
	} else {
		return currentZombie.GetWalkingZombie()
	}
}
