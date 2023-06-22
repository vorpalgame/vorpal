package zombiecide

//TODO Put sprit in common library.
import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

// Sprite interface/structure can be used for any type.
type Sprite interface {
	SetAudioFile(fileName string) Sprite
	GetAudioFile() string
	DoSendAudio()
	SendDrawEvent(drawEvent bus.DrawEvent, location Point, flip bool)
	SetImageTemplate(fileTemplate string) Sprite
	SetCurrentLocation(point Point) Sprite
	GetCurrentLocation() Point
	Start() Sprite
	Stop() Sprite
	IsStarted() bool
}

type SpriteData struct {
	currentFrame, maxFrame, repeatFrame, width, height int32
	fileTemplate                                       string
	audioFile                                          string
	currentLocation                                    Point
	started                                            bool
}

func NewSpriteData(x, y, width, height int32, imageTemplate, audioFile string) SpriteData {
	return SpriteData{1, x, y, width, height, imageTemplate, audioFile, &point{600, 600}, false}
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

func (s *SpriteData) DoSendAudio() {
	if !s.IsStarted() {
		s.Start()
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.GetAudioFile()).Play())
	}
}

// Default behavior...
func (s *SpriteData) Start() Sprite {
	s.started = true
	return s
}

func (s *SpriteData) Stop() Sprite {
	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.GetAudioFile()).Stop())
	s.currentFrame = 1
	s.started = false
	return s
}

func (s *SpriteData) Loop() {
	if s.currentFrame+1 >= s.maxFrame {
		s.currentFrame = 1
	}
}

func (s *SpriteData) NoLoop() {
	if s.currentFrame+1 >= s.maxFrame {
		s.currentFrame = s.maxFrame
	}
}

func (s *SpriteData) IncrementFrame() {
	s.repeatFrame++
	if s.repeatFrame > 4 {
		s.currentFrame++
		s.repeatFrame = 0
	}

}
func (s *SpriteData) SendDrawEvent(drawEvent bus.DrawEvent, location Point, flip bool) {

	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplate, s.currentFrame), location.GetX(), location.GetY(), s.width, s.height)

	layer.SetFlipHorizontal(flip)
	drawEvent.AddImageLayer(layer)
	bus.GetVorpalBus().SendDrawEvent(drawEvent)
}

// TODO The calcs are using the upper left for location relative to image and that probably isn't desired.
func (z *SpriteData) calculateMove(evt bus.MouseEvent) Point {
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

func (z *SpriteData) flipHorizontal(mouseEvent bus.MouseEvent) bool {
	return mouseEvent.GetX() < z.currentLocation.GetX()
}
