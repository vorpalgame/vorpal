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
	zombie := &idleZombie{newZombieData(15, 3, "idle", sprites)}
	zombie.GetFrameData().SetToLoop(false)
	return zombie
}

func (currentZombie *idleZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.sprites.GetAttackZombie()
	} else {
		point := currentZombie.GetCurrentLocation().CalculateMove(mouseEvent)
		if currentZombie.GetFrameData().UpdateIdleFrames(point) < 250 {
			currentZombie.GetCurrentLocation().Move(point)
			return currentZombie
		} else {
			return currentZombie.sprites.GetDeadZombie()
		}

	}
}
