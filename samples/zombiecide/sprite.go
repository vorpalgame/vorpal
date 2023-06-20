package zombiecide

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

// Sprite interface/structure can be used for any type.
type SpriteController interface {
	SetAudioFile(fileName string) SpriteController
	GetAudioFile() string
	SetImageTemplate(fileTemplate string) SpriteController
	SetCurrentLocation(point Point) SpriteController
	GetCurrentLocation() Point
	Start() SpriteController
	Stop() SpriteController
	IsStarted() bool
}

type spriteControllerData struct {
	currentFrame, maxFrame, repeatFrame, width, height int32
	fileTemplate                                       string
	audioFile                                          string
	currentLocation                                    Point
	started                                            bool
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

func (s *spriteControllerData) IsStarted() bool {
	return s.started
}
func (s *spriteControllerData) SetImageTemplate(fileTemplate string) SpriteController {
	s.fileTemplate = fileTemplate
	return s
}
func (s *spriteControllerData) GetAudioFile() string {
	return s.audioFile
}
func (s *spriteControllerData) SetAudioFile(fileName string) SpriteController {
	s.audioFile = fileName
	return s
}
func (s *spriteControllerData) SetCurrentLocation(point Point) SpriteController {
	s.currentLocation = point
	return s
}
func (s *spriteControllerData) GetCurrentLocation() Point {
	return s.currentLocation
}

func (s *spriteControllerData) doSendAudio() {
	if !s.IsStarted() {
		bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.GetAudioFile()).Play())
		s.Start()
	}
}

// Default behavior...
func (s *spriteControllerData) Start() SpriteController {
	s.started = true
	return s
}

func (s *spriteControllerData) Stop() SpriteController {
	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audioFile).Stop())
	s.currentFrame = 1
	s.started = false
	return s
}

func (s *spriteControllerData) loop() {
	if s.currentFrame+1 >= s.maxFrame {
		s.currentFrame = 1
	}
}

func (s *spriteControllerData) noLoop() {
	if s.currentFrame+1 >= s.maxFrame {
		s.currentFrame = s.maxFrame
	}
}

func (s *spriteControllerData) incrementFrame() {
	s.repeatFrame++
	if s.repeatFrame > 4 {
		s.currentFrame++
		s.repeatFrame = 0
	}

}
func (s *spriteControllerData) sendDrawEvent(drawEvent bus.DrawEvent, location Point, flip bool) {

	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplate, s.currentFrame), location.GetX(), location.GetY(), s.width, s.height)

	layer.SetFlipHorizontal(flip)
	drawEvent.AddImageLayer(layer)
	bus.GetVorpalBus().SendDrawEvent(drawEvent)
}

// TODO The calcs are using the upper left for location relative to image and that probably isn't desired.
func (z *spriteControllerData) calculateMove(evt bus.MouseEvent) Point {
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

func (z *spriteControllerData) flipHorizontal(mouseEvent bus.MouseEvent) bool {
	return mouseEvent.GetX() < z.currentLocation.GetX()
}
