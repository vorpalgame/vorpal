package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

type idleZombie struct {
	zombieData
}

type IdleZombie interface {
	ZombieSprite
}

func newIdleZombie(sprites ZombieSprites) IdleZombie {
	zombie := &idleZombie{newZombieData(15, 3, 200, 300, "idle", sprites)}
	zombie.GetFrameData().SetToLoop(false)
	return zombie
}

func (currentZombie *idleZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.sprites.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		if currentZombie.GetFrameData().UpdateIdleFrames(point) < 250 {
			currentZombie.GetCurrentLocation().Add(point)
			return currentZombie
		} else {
			return currentZombie.sprites.GetDeadZombie()
		}

	}
}
