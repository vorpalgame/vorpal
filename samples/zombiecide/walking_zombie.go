package zombiecide

import (
	"log"

	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/samples/lib"
)

type walkingZombie struct {
	lib.SpriteData
}

type WalkingZombie interface {
	lib.Sprite
	ZombieState
}

func newWalkingZombie() WalkingZombie {
	zombie := &walkingZombie{lib.NewSprite()}
	zombie.SetAudioFile(getZombieAudioTemplate("walk")).SetImageFileName(getZombieImageTemplate("walk")).SetToLoop(true).SetMaxFrame(10).SetRepeatFrame(5).SetImageScale(25)
	return zombie
}
func (currentZombie *walkingZombie) GetState(mouseEvent bus.MouseEvent, states ZombieStates) ZombieState {
	log.Default().Println("Walking")

	if mouseEvent.LeftButton().IsDown() {
		log.Default().Println("Returning attack")
		return states.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		if currentZombie.UpdateIdleFrames(point) < 10 {
			currentZombie.Move(point)
			return currentZombie
		} else {
			log.Default().Println("Returning idle")
			return states.GetIdleZombie()
		}
	}

}
