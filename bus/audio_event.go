package bus

//First cut of audio..

type AudioEventListener interface {
	OnAudioEvent(audioChannel <-chan AudioEvent)
}

type AudioEvent interface {
	GetAudio() string
}

type audioEvent struct {
	audio string
}

func NewAudioEvent(audio string) AudioEvent {
	return &audioEvent{audio}

}

func (e *audioEvent) GetAudio() string {
	return e.audio
}
