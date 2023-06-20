package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

type ZombieSprite interface {
	SpriteController
	RunSprite(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) ZombieSprite
}

// Probably a better factory pattern for this in idiomatic Golang
func NewZombie() ZombieSprite {

	walking := newWalkingZombie()
	dead := newDeadZombie()
	idle := newIdleZombie()
	attack := newAttackZombie()

	attack.SetWalkingZombie(walking)
	dead.SetWalkingZombie(walking)
	walking.SetAttackZombie(attack).SetIdleZombie(idle)
	idle.SetDeadZombie(dead).SetWalkingZombie(walking).SetAttackZombie(attack)

	//Start walking...
	return walking
}

func newSpriteControllerData(x, y, width, height int32, name string) spriteControllerData {
	return spriteControllerData{1, x, y, width, height, getZombieImageTemplate(name), getZombieAudioTemplate(name), &point{600, 600}, false}
}

func getZombieImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getZombieAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
}

func doTransition(currentState, nextState ZombieSprite) ZombieSprite {
	currentState.Stop()
	nextState.SetCurrentLocation(currentState.GetCurrentLocation())
	return nextState
}

func doSendAudio(currentState ZombieSprite) {
	if !currentState.IsStarted() {
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(currentState.GetAudioFile()).Play())
		currentState.Start()
	}
}

func doIdleCount(idleCount int32, point Point) int32 {
	if point.GetY() == 0 && point.GetX() == 0 {
		idleCount++
	} else {
		idleCount = 0
	}
	return idleCount
}
