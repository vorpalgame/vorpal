package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type deadZombie struct {
	lib.SpriteData
}

type DeadZombie interface {
	lib.Sprite
	ZombieState
}

func newDeadZombie() DeadZombie {
	zombie := &deadZombie{lib.NewSprite()}
	zombie.SetAudioFile(getZombieAudioTemplate("dead")).SetImageFileName(getZombieImageTemplate("dead")).SetToLoop(false).SetMaxFrame(10).SetRepeatFrame(5).SetImageScale(25)
	return zombie
}

func (currentZombie *deadZombie) GetState(mouseEvent bus.MouseEvent, states ZombieStates) ZombieState {

	point := currentZombie.CalculateMove(mouseEvent)
	if currentZombie.UpdateIdleFrames(point) > 0 {
		return currentZombie
	} else {
		return states.GetWalkingZombie()
	}
}
