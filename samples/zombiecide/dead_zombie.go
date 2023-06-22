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
	zombie := &deadZombie{newZombieData(12, 3, 300, 300, "dead", sprites)}
	zombie.GetFrameData().SetToLoop(false)
	return zombie
}

func (currentZombie *deadZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	point := currentZombie.calculateMove(mouseEvent)
	currentZombie.currentLocation.Add(point)
	if currentZombie.GetFrameData().UpdateIdleFrames(point) > 0 {
		return currentZombie
	} else {
		return currentZombie.Transition(currentZombie.sprites.GetWalkingZombie())
	}
}

func (currentZombie *deadZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {
	currentZombie.RunAudio()

	currentZombie.SendDrawEvent(drawEvent, currentZombie.currentLocation, currentZombie.flipHorizontal(mouseEvent))
	currentZombie.GetFrameData().Increment()

}
