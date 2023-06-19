package zombiecide

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

type SpriteController interface {
	RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool)
	StopSprite() SpriteController
	SetAudio(fileName string) SpriteController
	SetImageTemplate(fileTemplate string) SpriteController
	SetToLoop(repeat bool) SpriteController
}

type spriteControllerData struct {
	currentFrame, maxFrame, repeatFrame int
	width, height                       int32
	fileTemplate                        string
	audioFile                           string
	repeat                              bool
}

// TODO Redesign to remove nil checks for audio event
// TODO regularize this constructor to either take the audio event or to remove the file and use builder pattern only.
func NewSpriteController(maxFrame, repeatFrame int, width, height int32) SpriteController {
	return &spriteControllerData{1, maxFrame, repeatFrame, width, height, "", "", true} //Loop by default//no audio by default
}

func (s *spriteControllerData) SetImageTemplate(fileTemplate string) SpriteController {
	s.fileTemplate = fileTemplate
	return s
}

func (s *spriteControllerData) SetAudio(fileName string) SpriteController {
	s.audioFile = fileName
	return s
}
func (s *spriteControllerData) SetToLoop(repeat bool) SpriteController {
	s.repeat = repeat
	return s
}
func (s *spriteControllerData) StopSprite() SpriteController {
	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audioFile).Stop())
	s.currentFrame = 1
	return s
}

// TODO Think this control structure through more.
func (s *spriteControllerData) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {

	if s.currentFrame >= s.maxFrame && !s.repeat {
		s.currentFrame = s.maxFrame
		s.doSprite(drawEvent, p, flipHorizontal)
	} else {
		if s.currentFrame == 1 {
			bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audioFile).Play())
		}
		//We repeat frames to prevent blur and jitters and make it smoother.
		s.repeatFrame++
		if s.repeatFrame > 4 {
			s.currentFrame++
			s.repeatFrame = 0
		}

		if s.currentFrame >= s.maxFrame && s.repeat {
			s.currentFrame = 1
		}
		s.doSprite(drawEvent, p, flipHorizontal)
	}

}

func (s *spriteControllerData) doSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {
	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplate, s.currentFrame), p.GetX(), p.GetY(), s.width, s.height)

	layer.SetFlipHorizontal(flipHorizontal)
	drawEvent.AddImageLayer(layer)
	bus.GetVorpalBus().SendDrawEvent(drawEvent)
}
