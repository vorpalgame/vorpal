package zombiecide

import "github.com/vorpalgame/vorpal/bus"

type walkingZombie struct {
	zombieData
}

type WalkingZombie interface {
	ZombieSprite
}

func newWalkingZombie(sprites ZombieSprites) WalkingZombie {
	return &walkingZombie{newZombieData(10, 3, 200, 300, "walk", sprites)}
}
func (currentZombie *walkingZombie) GetState(mouseEvent bus.MouseEvent) ZombieSprite {

	point := currentZombie.calculateMove(mouseEvent)
	currentZombie.currentLocation.Add(point)
	if mouseEvent.LeftButton().IsDown() {
		return currentZombie.doTransition(currentZombie.sprites.GetAttackZombie())
	} else {
		if currentZombie.updateIdleCount(point) < 50 {
			return currentZombie
		} else {
			return currentZombie.doTransition(currentZombie.sprites.GetIdleZombie())
		}
	}

}

func (s *walkingZombie) RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {

	s.DoSendAudio()
	s.SendDrawEvent(drawEvent, s.currentLocation, s.flipHorizontal(mouseEvent))
	s.IncrementFrame()
	s.Loop()

}
