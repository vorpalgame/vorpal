package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

// Sprite interface/structure can be used for any type.
type SpriteController interface {
	RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool)
	StopSprite() SpriteController
	SetAudio(fileName string) SpriteController
	SetImageTemplate(fileTemplate string) SpriteController
}

type spriteControllerData struct {
	currentFrame, maxFrame, repeatFrame, width, height int32
	fileTemplate                                       string
	audioFile                                          string
}

// Abstract/template function
func (s *spriteControllerData) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {}

func (s *spriteControllerData) SetImageTemplate(fileTemplate string) SpriteController {
	s.fileTemplate = fileTemplate
	return s
}

func (s *spriteControllerData) SetAudio(fileName string) SpriteController {
	s.audioFile = fileName
	return s
}

// Default behavior...
func (s *spriteControllerData) StopSprite() SpriteController {
	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audioFile).Stop())
	s.currentFrame = 1
	return s
}
