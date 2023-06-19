package bus

//First cut of audio..

type AudioEventListener interface {
	OnAudioEvent(audioChannel <-chan AudioEvent)
}

type AudioEvent interface {
	GetAudio() string
	SetAudio(string) AudioEvent
	Play() bool
	SetPlay(play bool) AudioEvent
}

type audioEvent struct {
	audio string
	play  bool
}

func NewAudioEvent(audio string, play bool) AudioEvent {
	return &audioEvent{audio, play}

}

func (e *audioEvent) SetAudio(fileName string) AudioEvent {
	e.audio = fileName
	return e
}

func (e *audioEvent) GetAudio() string {
	return e.audio
}

func (e *audioEvent) Play() bool {
	return e.play
}

func (e *audioEvent) SetPlay(play bool) AudioEvent {
	e.play = play
	return e
}
