package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"log"
)

// /////////////////////////////////////////////////////////////////
// /// Audio Event Processor
// /////////////////////////////////////////////////////////////////

type audioData struct {
	MediaCache
}

func NewAudioEventProcessor(mediaCache MediaCache) bus.AudioEventProcessor {
	return &audioData{mediaCache}
}

func (dep *audioData) ProcessAudioEvent(evt bus.AudioEvent) {

	if evt != nil {
		log.Default().Println(evt.GetAudioFile())
		currentAudio := *dep.GetAudio(evt.GetAudioFile())
		switch evt := evt.(type) {
		case bus.PlayAudioEvent:
			if !rl.IsSoundPlaying(currentAudio) {
				if evt.IncrementCount() == 1 || evt.IsLoop() {
					for !rl.IsSoundReady(currentAudio) {
					} //Loop until resource is ready
					rl.PlaySound(currentAudio)
				}
			}
		case bus.StopAudioEvent:
			rl.StopSound(currentAudio)
			evt.ResetCount()
			for rl.IsSoundPlaying(currentAudio) {
			} //Wait until stopped
		}
	}

}
