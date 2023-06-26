package zombiecide

import (
	//"log"

	"log"

	"github.com/vorpalgame/vorpal/bus"
)

// Outter wrapper that keeps track of states and transtion logic.
func NewZombie(percentScale int32) ZombieSprite {
	zs := NewZmobieStates(percentScale)
	return &zombieData{zs, zs.GetWalkingZombie()}
}

type ZombieSprite interface {
	Execute(drawEvent bus.DrawEvent, keyEvent bus.KeyEvent, mouseEvent bus.MouseEvent)
}

type zombieData struct {
	sprites ZombieStates
	current ZombieState
}

func (zs *zombieData) Execute(drawEvent bus.DrawEvent, keyEvent bus.KeyEvent, evt bus.MouseEvent) {
	previousState := zs.current
	zs.current = zs.current.GetState(evt)
	if keyEvent != nil {
		log.Default().Println(keyEvent.GetKey().ToString())
	}

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
