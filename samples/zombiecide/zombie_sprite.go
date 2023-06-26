package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

// Outter wrapper that keeps track of states and transtion logic.
func NewZombie() ZombieSprite {
	zs := NewZombieSprites()
	return &zombieData{zs, zs.GetWalkingZombie()}
}

type ZombieSprite interface {
	Execute(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent)
}

type zombieData struct {
	sprites ZombieStates
	current ZombieState
}

func (zs *zombieData) Execute(drawEvent bus.DrawEvent, evt bus.MouseEvent) {
	previousState := zs.current
	zs.current = zs.current.GetState(evt, zs.sprites)

	if previousState != zs.current {
		zs.current.SetCurrentLocation(previousState.GetCurrentLocation())
		bus.GetVorpalBus().SendAudioEvent(previousState.GetStopAudioEvent())
		previousState.Stop()

	}
	if !zs.current.IsStarted() {
		bus.GetVorpalBus().SendAudioEvent(zs.current.GetPlayAudioEvent())
		zs.current.Start()
	}
	drawEvent.AddImageLayer(zs.current.CreateImage(evt))

}
