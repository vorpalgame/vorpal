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
	zombie := &idleZombie{NewZombieData(15, 3, "idle", sprites)}
	zombie.SetToLoop(false)
	return zombie
}

func (currentZombie *idleZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.sprites.GetAttackZombie()
	} else {
		point := currentZombie.CalculateMove(mouseEvent)
		if currentZombie.UpdateIdleFrames(point) < 250 {
			currentZombie.Move(point)
			return currentZombie
		} else {
			return currentZombie.sprites.GetDeadZombie()
		}

	}
}
