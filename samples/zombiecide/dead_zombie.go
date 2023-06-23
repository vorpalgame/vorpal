package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type deadZombie struct {
	zombieData
}

//TODO Some of the methods for setting shoudl be private to the package.
type DeadZombie interface {
	ZombieSprite
}

func newDeadZombie(sprites ZombieSprites) DeadZombie {
	zombie := &deadZombie{newZombieData(12, 3, "dead", sprites)}
	zombie.GetFrameData().SetToLoop(false)
	return zombie
}

func (currentZombie *deadZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	point := currentZombie.CalculateMove(mouseEvent)
	currentZombie.GetCurrentLocation().Add(point)
	if currentZombie.GetFrameData().UpdateIdleFrames(point) > 0 {
		return currentZombie
	} else {
		return currentZombie.sprites.GetWalkingZombie()
	}
}
