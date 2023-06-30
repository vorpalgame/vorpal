package lib

import "github.com/vorpalgame/vorpal/bus"

func NewAudioController() AudioController {
	return &audioControllerData{}
}

type AudioController interface {
	SetAudioFile(fileName string) AudioController
	GetAudioFile() string
	GetPlayAudioEvent() bus.AudioEvent
	GetStopAudioEvent() bus.AudioEvent
}

type audioControllerData struct {
	audioFile string
}

func (s *audioControllerData) GetAudioFile() string {
	return s.audioFile
}
func (s *audioControllerData) SetAudioFile(fileName string) AudioController {
	s.audioFile = fileName
	return s
}

func (s *audioControllerData) GetPlayAudioEvent() bus.AudioEvent {
	return bus.NewAudioEvent(s.GetAudioFile()).Play()
}

func (s *audioControllerData) GetStopAudioEvent() bus.AudioEvent {
	return bus.NewAudioEvent(s.GetAudioFile()).Stop()
}
