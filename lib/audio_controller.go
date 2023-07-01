package lib

//TODO Add tracking states like play, pause, stopped, loop, etc.
//Currently an issue with cyclic dependency with bus.
func NewAudioController() AudioController {
	return &audioControllerData{}
}

type AudioController interface {
	SetAudioFile(fileName string) AudioController
	GetAudioFile() string
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
