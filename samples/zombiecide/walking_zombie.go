package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type walkingZombie struct {
	zombieData
}

type WalkingZombie interface {
	ZombieSprite
}

func newWalkingZombie(sprites ZombieSprites) WalkingZombie {
	zombie := &walkingZombie{newZombieData(10, 3, 200, 300, "walk", sprites)}
	zombie.GetFrameData().SetToLoop(true)
	return zombie
}
func (currentZombie *walkingZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	point := currentZombie.CalculateMove(mouseEvent)
	currentZombie.GetCurrentLocation().Add(point)

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.sprites.GetAttackZombie()
	} else {
		if currentZombie.GetFrameData().UpdateIdleFrames(point) < 50 {
			return currentZombie
		} else {
			return currentZombie.sprites.GetIdleZombie()
		}
	}

}
