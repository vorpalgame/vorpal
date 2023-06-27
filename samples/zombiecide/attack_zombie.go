package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type attackZombie struct {
	zombieStateData
}

type AttackZombie interface {
	ZombieState
}

func (currentZombie *attackZombie) GetState(mouseEvent bus.MouseEvent) ZombieState {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie
	} else {
		return currentZombie.GetWalkingZombie()
	}
}
