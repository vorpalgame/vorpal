package lib

//Having to return the struct instead of the interface is an unfortunate
//side effect of the YAML marshaling.

func NewAudioState(fileName string, loop bool) AudioState {
	return &AudioStateData{fileName, loop, 0}
}

type AudioState interface {
	SetAudioFile(fileName string) AudioState
	SetLoop(shouldLoop bool) AudioState
	IsLoop() bool
	GetAudioFile() string
	IncrementCount() int32
	ResetCount() AudioState
}

// UC so we don't end up with marshaling issues..
type AudioStateData struct {
	AudioFile string `yaml:"AudioFile"`
	LoopAudio bool   `yaml:"LoopAudio"`
	count     int32  `yaml:"-"`
}

func (s *AudioStateData) IncrementCount() int32 {
	s.count++
	return s.count
}

func (s *AudioStateData) ResetCount() AudioState {
	s.count = 0
	return s
}

func (s *AudioStateData) SetLoop(shouldLoop bool) AudioState {
	s.LoopAudio = shouldLoop
	return s
}

func (s *AudioStateData) IsLoop() bool {
	return s.LoopAudio
}

func (s *AudioStateData) GetAudioFile() string {
	return s.AudioFile
}
func (s *AudioStateData) SetAudioFile(fileName string) AudioState {
	s.AudioFile = fileName
	return s
}
