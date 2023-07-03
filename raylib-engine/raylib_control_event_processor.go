package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// /////////////////////////////////////////////////////////////////
// /// Raylib Control Event Processor
// /////////////////////////////////////////////////////////////////

type controlData struct {
}

func NewControlEventProcessor() bus.ControlEventProcessor {
	return &controlData{}
}

func (dep *controlData) ProcessControlEvent(evt bus.ControlEvent) {

	if evt != nil {
		switch evt := evt.(type) {
		case bus.WindowTitleEvent:
			rl.SetWindowTitle(evt.GetTitle())
		}

	}
}
