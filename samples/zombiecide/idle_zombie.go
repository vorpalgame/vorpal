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
		return currentZombie.Transition(currentZombie.sprites.GetAttackZombie())
	} else {
		point := currentZombie.calculateMove(mouseEvent)
		if currentZombie.GetFrameData().UpdateIdleFrames(point) < 250 {
			currentZombie.currentLocation.Add(point)
			return currentZombie
		} else {
			return currentZombie.Transition(currentZombie.sprites.GetDeadZombie())
		}

	}
}

func (currentZombie *idleZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {

	currentZombie.RunAudio()

	currentZombie.SendDrawEvent(drawEvent, currentZombie.currentLocation, currentZombie.flipHorizontal(mouseEvent))
	currentZombie.GetFrameData().Increment()

}
