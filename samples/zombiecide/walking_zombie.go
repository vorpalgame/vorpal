package zombiecide

import (
	"log"

	"github.com/vorpalgame/vorpal/bus"
)

type walkingZombie struct {
	zombieData
}

type WalkingZombie interface {
	ZombieSprite
}

func newWalkingZombie(sprites ZombieSprites) WalkingZombie {
	zombie := &walkingZombie{NewZombieData(10, 5, "walk", sprites)}
	zombie.SetToLoop(true)
	return zombie
}
func (currentZombie *walkingZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {
	point := currentZombie.CalculateMove(mouseEvent)
	if mouseEvent.LeftButton().IsDown() {
		log.Default().Println(currentZombie.sprites)
		return currentZombie.sprites.GetAttackZombie()
	} else {
		if currentZombie.UpdateIdleFrames(point) < 50 {
			currentZombie.Move(point)
			return currentZombie
		} else {
			return currentZombie.sprites.GetIdleZombie()
		}
	}

}
