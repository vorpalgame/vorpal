package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"log"
)

// /////////////////////////////////////////////////////////////////
// /// Audio Event Processor
// /////////////////////////////////////////////////////////////////

var raylibProcessAudioEvent = func(evt bus.AudioEvent, cache MediaCache) {

	if evt != nil {
		log.Default().Println(evt.GetAudioFile())
		currentAudio := *cache.GetAudio(evt.GetAudioFile())
		switch evt.(type) {
		case bus.PlayAudioEvent:
			if !rl.IsSoundPlaying(currentAudio) {

				for !rl.IsSoundReady(currentAudio) {
				} //Loop until resource is ready
				rl.PlaySound(currentAudio)

			}
		case bus.StopAudioEvent:
			rl.StopSound(currentAudio)

		}
	}

}
