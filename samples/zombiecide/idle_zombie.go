package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type idleZombie struct {
	zombieStateData
}

type IdleZombie interface {
	lib.Sprite
	ZombieState
}

func (currentZombie *idleZombie) GetState(mouseEvent bus.MouseEvent) ZombieState {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		if currentZombie.UpdateIdleFrames(point) < 150 {
			return currentZombie
		} else {
			return currentZombie.GetDeadZombie()
		}

	}
}
