package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type walkingZombie struct {
	zombieStateData
}

type WalkingZombie interface {
	ZombieState
}

// Only GetState is really unique to a given state to determine behavior
func (currentZombie *walkingZombie) GetState(mouseEvent bus.MouseEvent) ZombieState {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		if currentZombie.UpdateIdleFrames(point) < 50 {
			currentZombie.Move(point)
			return currentZombie
		} else {
			return currentZombie.GetIdleZombie()
		}
	}

}
