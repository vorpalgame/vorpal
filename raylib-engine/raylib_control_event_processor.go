package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// /////////////////////////////////////////////////////////////////
// /// Control Event Processor
// /////////////////////////////////////////////////////////////////
type ControlEventProcessor interface {
	processControlEvent(evt bus.ControlEvent)
}
type controlData struct {
}

func NewControlEventProcessor() ControlEventProcessor {
	return &controlData{}
}

func (dep *controlData) processControlEvent(evt bus.ControlEvent) {

	if evt != nil {
		switch evt := evt.(type) {
		case bus.WindowTitleEvent:
			rl.SetWindowTitle(evt.GetTitle())
		}

	}
}
