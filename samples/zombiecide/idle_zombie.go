package zombiecide

import (
	"log"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type idleZombie struct {
	lib.SpriteData
}

type IdleZombie interface {
	lib.Sprite
	ZombieState
}

func newIdleZombie() IdleZombie {
	zombie := &idleZombie{lib.NewSprite()}
	zombie.SetAudioFile(getZombieAudioTemplate("idle")).SetImageFileName(getZombieImageTemplate("idle")).SetToLoop(false).SetMaxFrame(15).SetRepeatFrame(5).SetImageScale(25)
	return zombie
}

func (currentZombie *idleZombie) GetState(mouseEvent bus.MouseEvent, states ZombieSprites) ZombieState {

	log.Default().Println("Idle")

	if mouseEvent.LeftButton().IsDown() {
		return states.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		if currentZombie.UpdateIdleFrames(point) < 150 {
			return currentZombie
		} else {
			return states.GetDeadZombie()
		}

	}
}
