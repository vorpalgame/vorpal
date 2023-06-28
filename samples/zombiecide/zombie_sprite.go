package zombiecide

import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

// Outter wrapper that keeps track of states and transtion logic.
// TODO We need to refactor to get ric of Sprite and use more composable elements.

func NewZombie(percentScale int32) ZombieSprite {
	zs := NewZombieStates(percentScale)
	return &zombieData{zs}
}

type ZombieSprite interface {
	Execute(drawEvent bus.DrawEvent, keyEvent bus.KeyEvent, mouseEvent bus.MouseEvent)
	Stop()
}

type zombieData struct {
	sprites ZombieStates
}

func (zs *zombieData) Stop() {
	zs.sprites.GetCurrent().Stop()
	bus.GetVorpalBus().SendAudioEvent(zs.sprites.GetCurrent().GetStopAudioEvent())
}

func (zs *zombieData) Execute(drawEvent bus.DrawEvent, keyEvent bus.KeyEvent, evt bus.MouseEvent) {
	previousState := zs.sprites.GetCurrent()
	current := previousState.GetState(evt)

	if previousState != current {
		current.SetCurrentLocation(previousState.GetCurrentLocation())
		zs.sprites.SetCurrent(current)
		bus.GetVorpalBus().SendAudioEvent(previousState.GetStopAudioEvent())
		previousState.Stop()

	}
	if !current.IsStarted() {
		bus.GetVorpalBus().SendAudioEvent(current.GetPlayAudioEvent())
		current.Start()
	}
	drawEvent.AddImageLayer(current.CreateImage(evt))

}
