package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// /////////////////////////////////////////////////////////////////
// /// Raylib Control Event Processor
// /////////////////////////////////////////////////////////////////

var raylibProcessControlEvent = func(evts []bus.ControlEvent) {

	if evts != nil {
		for _, evt := range evts {
			switch evt := evt.(type) {
			case bus.WindowTitleEvent:
				rl.SetWindowTitle(evt.GetTitle())

			case bus.WindowSizeEvent:
				rl.SetWindowSize(evt.GetWindowSize())
			}

		}
	}
}
