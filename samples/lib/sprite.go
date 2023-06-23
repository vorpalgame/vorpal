package lib

//TODO Put sprit in common library.
import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

func NewSpriteData(maxFrames, repeatPerFrame, width, height int32, imageTemplate, audioFile string) SpriteData {
	return SpriteData{width, height, imageTemplate, audioFile, &point{600, 600}, false, false, &frameData{1, maxFrames, repeatPerFrame, 0, 0, false}}
}

// Sprite interface/structure can be used for any type.
type Sprite interface {
	Start() Sprite
	Stop() Sprite
	IsStarted() bool

	GetFrameData() FrameData

	SetAudioFile(fileName string) Sprite
	GetAudioFile() string
	GetPlayAudioEvent() bus.AudioEvent
	GetStopAudioEvent() bus.AudioEvent
	SetImageTemplate(fileTemplate string) Sprite
	SetCurrentLocation(point Point) Sprite

	GetCurrentLocation() Point
	CalculateMove(evt bus.MouseEvent) Point
	FlipHorizontal(evt bus.MouseEvent) bool
	CreateImageLayer(mouseEvent bus.MouseEvent) *bus.ImageLayer
}

type SpriteData struct {
	width, height   int32
	fileTemplate    string
	audioFile       string
	currentLocation Point
	started         bool
	audioRunning    bool
	frameData       FrameData
}

type FrameData interface {
	GetCurrentFrame() int32
	GetMaxFrame() int32
	SetToLoop(bool)
	UpdateIdleFrames(point Point) int32
	GetIdleFrames() int32
	IsLoop() bool
	Increment()
	Reset()
}
type frameData struct {
	currentFrame, maxFrame, repeatPerFrame, currentFrameRepeats, idleFrames int32
	loop                                                                    bool
}

func (fd *frameData) Increment() {
	fd.currentFrameRepeats++
	if fd.currentFrameRepeats > fd.repeatPerFrame {
		fd.currentFrameRepeats = 0
		fd.currentFrame++
		if fd.currentFrame > fd.maxFrame {

			if fd.IsLoop() {
				fd.currentFrame = 1
			} else {
				fd.currentFrame = fd.maxFrame
			}
		}
	}
}

func (fd *frameData) UpdateIdleFrames(point Point) int32 {
	if point.GetY() == 0 && point.GetX() == 0 {
		fd.idleFrames++
	} else {
		fd.idleFrames = 0
	}
	return fd.idleFrames
}

func (fd *frameData) GetIdleFrames() int32 {
	return fd.idleFrames
}
func (fd *frameData) SetToLoop(repeat bool) {
	fd.loop = repeat
}

func (fd *frameData) IsLoop() bool {
	return fd.loop
}

func (fd *frameData) Reset() {
	fd.currentFrame = 1
	fd.idleFrames = 0
}

func (fd *frameData) GetCurrentFrame() int32 {
	return fd.currentFrame
}

func (fd *frameData) GetMaxFrame() int32 {
	return fd.maxFrame
}

type Point interface {
	GetX() int32
	GetY() int32
	Add(Point)
}

type point struct {
	x, y int32
}

func (p *point) GetX() int32 {
	return p.x
}

func (p *point) GetY() int32 {
	return p.y
}

func (p *point) Add(addPoint Point) {
	p.x += addPoint.GetX()
	p.y += addPoint.GetY()
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
func (s *SpriteData) SetCurrentLocation(point Point) Sprite {
	s.currentLocation = point
	return s
}
func (s *SpriteData) GetCurrentLocation() Point {
	return s.currentLocation
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
	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplate, s.frameData.GetCurrentFrame()), s.currentLocation.GetX(), s.currentLocation.GetY(), s.width, s.height)
	layer.SetFlipHorizontal(s.FlipHorizontal(mouseEvent))
	s.GetFrameData().Increment()
	return &layer
}

// TODO The calcs are using the upper left for location relative to image and that probably isn't desired.
// values of x,y, and the width of window should be in configuration data in struct so they can be varied...
// Clean up required...
func (z *SpriteData) CalculateMove(evt bus.MouseEvent) Point {
	x := int32(-4)
	y := int32(-2)
	point := point{x, y}
	//abs math function is floating point so just -1 multiple
	if evt.GetX() > z.currentLocation.GetX() {
		point.x = point.x * -1
	}
	if evt.GetY() > z.currentLocation.GetY() {
		point.y = point.y * -1
	}

	var xOffset = evt.GetX() - z.currentLocation.GetX()
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < 5 {
		point.x = 0
	}
	yOffset := evt.GetY() - z.currentLocation.GetY()
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < 5 {
		point.y = 0
	}
	return &point
}

func (z *SpriteData) FlipHorizontal(mouseEvent bus.MouseEvent) bool {
	return mouseEvent.GetX() < z.currentLocation.GetX()
}
