package lib

//TODO Put sprit in common library.
import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

// TODO Break into composable elements now that we are in lib...
func NewSpriteData(maxFrames, repeatPerFrame, scale int32, imageTemplate, audioFile string) SpriteData {
	return SpriteData{NewCurrentLocation(), scale, imageTemplate, audioFile, false, false, &frameData{1, maxFrames, repeatPerFrame, 0, 0, false}}
}

// Sprite interface/structure can be used for any type.
type Sprite interface {
	Start() Sprite
	Stop() Sprite
	IsStarted() bool

	GetFrameData() FrameData
	GetCurrentLocation() CurrentLocation
	SetCurrentLocation(currentLocation CurrentLocation)

	SetAudioFile(fileName string) Sprite
	GetAudioFile() string
	GetPlayAudioEvent() bus.AudioEvent
	GetStopAudioEvent() bus.AudioEvent
	SetImageTemplate(fileTemplate string) Sprite

	GetScale() int32

	FlipHorizontal(evt bus.MouseEvent) bool
	CreateImageLayer(mouseEvent bus.MouseEvent) *bus.ImageLayer
}

type SpriteData struct {
	currentLocation CurrentLocation
	scale           int32
	fileTemplate    string
	audioFile       string
	started         bool
	audioRunning    bool
	frameData       FrameData
}

func (s *SpriteData) GetFrameData() FrameData {
	return s.frameData
}
func (s *SpriteData) IsStarted() bool {
	return s.started
}
func (s *SpriteData) SetImageTemplate(fileTemplate string) Sprite {
	s.fileTemplate = fileTemplate
	return s
}
func (s *SpriteData) GetAudioFile() string {
	return s.audioFile
}
func (s *SpriteData) SetAudioFile(fileName string) Sprite {
	s.audioFile = fileName
	return s
}
func (s *SpriteData) SetCurrentLocation(location CurrentLocation) {
	s.currentLocation = location
}
func (s *SpriteData) GetCurrentLocation() CurrentLocation {
	return s.currentLocation
}
func (s *SpriteData) GetScale() int32 {
	return s.scale
}

func (s *SpriteData) GetPlayAudioEvent() bus.AudioEvent {
	return bus.NewAudioEvent(s.GetAudioFile()).Play()
}

func (s *SpriteData) GetStopAudioEvent() bus.AudioEvent {
	return bus.NewAudioEvent(s.GetAudioFile()).Stop()
}

func (s *SpriteData) Start() Sprite {
	if !s.IsStarted() {
		s.frameData.Reset()
		s.started = true
	}
	return s
}

func (s *SpriteData) Stop() Sprite {
	s.frameData.Reset()
	s.started = false
	return s
}

// TODO The flip horizontal may belong at higher abstraction level and then
// this would not require the mouse event.
func (s *SpriteData) CreateImageLayer(mouseEvent bus.MouseEvent) *bus.ImageLayer {
	imgData := bus.NewImageMetadata(fmt.Sprintf(s.fileTemplate, s.frameData.GetCurrentFrame()), s.currentLocation.GetX(), s.currentLocation.GetY(), s.GetScale())
	imgData.SetFlipHorizontal(s.FlipHorizontal(mouseEvent))
	layer := bus.NewImageLayer().AddLayerData(imgData)

	s.GetFrameData().Increment()
	return &layer
}

func (z *SpriteData) FlipHorizontal(mouseEvent bus.MouseEvent) bool {
	return mouseEvent.GetX() < z.currentLocation.GetX()
}
