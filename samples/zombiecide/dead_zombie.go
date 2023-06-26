package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type deadZombie struct {
	zombieStateData
}

type DeadZombie interface {
	lib.Sprite
	ZombieState
}

func (currentZombie *deadZombie) GetState(mouseEvent bus.MouseEvent) ZombieState {

	point := currentZombie.CalculateMove(mouseEvent)
	if currentZombie.UpdateIdleFrames(point) > 0 {
		return currentZombie
	} else {
		return currentZombie.GetWalkingZombie()
	}
}
