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
	return &deadZombie{newZombieData(12, 3, 300, 300, "dead", sprites)}
}

func (currentZombie *deadZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	point := currentZombie.calculateMove(mouseEvent)
	currentZombie.currentLocation.Add(point)
	if currentZombie.updateIdleCount(point) > 0 {
		return currentZombie
	} else {
		return currentZombie.doTransition(currentZombie.sprites.GetWalkingZombie())
	}
}

func (currentZombie *deadZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {
	currentZombie.DoSendAudio()
	currentZombie.SendDrawEvent(drawEvent, currentZombie.currentLocation, currentZombie.flipHorizontal(mouseEvent))
	currentZombie.IncrementFrame()
	currentZombie.NoLoop()

}
