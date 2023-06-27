package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type idleZombie struct {
	zombieStateData
}

type IdleZombie interface {
	ZombieState
}

func (currentZombie *idleZombie) GetState(mouseEvent bus.MouseEvent) ZombieState {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		idleFrames := currentZombie.UpdateIdleFrames(point)
		if idleFrames == 0 {
			return currentZombie.GetWalkingZombie()
		} else if idleFrames < 150 {
			return currentZombie
		} else {
			return currentZombie.GetDeadZombie()
		}

	}
}
