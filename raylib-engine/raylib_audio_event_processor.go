package raylibengine

import (
	//rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// TODO Refactor audio events and this processor...
// /////////////////////////////////////////////////////////////////
// /// Audio Event Processor
// /////////////////////////////////////////////////////////////////
type AudioEventProcessor interface {
	processAudioEvent(evt bus.AudioEvent)
}
type audioData struct {
}

func NewAudioEventProcessor() AudioEventProcessor {
	return &audioData{}
}

func (dep *audioData) processAudioEvent(evt bus.AudioEvent) {

}
