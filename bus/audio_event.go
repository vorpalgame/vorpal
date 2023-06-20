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
	IsLoop() bool
	Play() AudioEvent
	Stop() AudioEvent
	Loop(bool) AudioEvent
	IncrementCount() int32
	ResetCount()
}

type audioEvent struct {
	audio      string
	play, loop bool
	count      int32
}

func NewAudioEvent(audio string) AudioEvent {
	return &audioEvent{audio, false, false, 0}

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

func (e *audioEvent) IsLoop() bool {
	return e.loop
}
func (e *audioEvent) Loop(loop bool) AudioEvent {
	e.loop = loop
	return e
}
func (e *audioEvent) Play() AudioEvent {
	e.play = true
	return e
}
func (e *audioEvent) Stop() AudioEvent {
	e.play = false
	return e
}
func (e *audioEvent) IncrementCount() int32 {
	e.count++
	return e.count
}
func (e *audioEvent) ResetCount() {
	e.count = 0
}
