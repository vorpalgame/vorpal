package bus

import "github.com/vorpalgame/vorpal/lib"

/// Constructors

func NewPlayAudioEvent(state lib.AudioState) PlayAudioEvent {
	state.ResetCount() //Make sure it indicates it is ready to play.
	return &playAudioEventData{state}
}

func NewStopAudioEvent(state lib.AudioState) StopAudioEvent {
	return &stopAudioEventData{state}
}

type AudioEventListener interface {
	OnAudioEvent(audioChannel <-chan AudioEvent)
}
type AudioEventProcessor interface {
	ProcessAudioEvent(event AudioEvent)
}

//////Basic AudioEvent //////

type AudioEvent interface {
	lib.AudioState
}
type audioEventData struct {
	lib.AudioState
}

func (e *audioEventData) IncrementCount() int32 {
	return e.IncrementCount()
}
func (e *audioEventData) ResetCount() {
	e.ResetCount()
}

/////PlayAudioEvent

type PlayAudioEvent interface {
	AudioEvent
}
type playAudioEventData struct {
	AudioEvent
}

/////StopAudioEvent

type StopAudioEvent interface {
	AudioEvent
}
type stopAudioEventData struct {
	AudioEvent
}
