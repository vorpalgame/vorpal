package bus

//First cut of audio..

type AudioEventListener interface {
	OnAudioEvent(audioChannel <-chan AudioEvent)
}

type AudioEvent interface {
	GetAudio() string
	SetAudio(string) AudioEvent
	IsStop() bool
	IsPlay() bool
	Play() AudioEvent
	Stop() AudioEvent
}

type audioEvent struct {
	audio string
	play  bool
}

func NewAudioEvent(audio string) AudioEvent {
	return &audioEvent{audio, false}

}

func (e *audioEvent) SetAudio(fileName string) AudioEvent {
	e.audio = fileName
	return e
}

func (e *audioEvent) GetAudio() string {
	return e.audio
}

func (e *audioEvent) IsPlay() bool {
	return e.play
}

func (e *audioEvent) IsStop() bool {
	return !e.play
}

func (e *audioEvent) Play() AudioEvent {
	e.play = true
	return e
}
func (e *audioEvent) Stop() AudioEvent {
	e.play = false
	return e
}
