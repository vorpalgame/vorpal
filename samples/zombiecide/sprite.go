package zombiecide

import (
	"fmt"
	"log"

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
	audio                               string
	repeat                              bool
}

// TODO regularize this constructor to either take the audio event or to remove the file and use builder pattern only.
func NewSpriteController(maxFrame, repeatFrame int, width, height int32) SpriteController {
	return &spriteControllerData{1, maxFrame, repeatFrame, width, height, "", "", true} //Loop by default//no audio by default
}

func (s *spriteControllerData) SetImageTemplate(fileTemplate string) SpriteController {
	s.fileTemplate = fileTemplate
	return s
}

func (s *spriteControllerData) SetAudio(fileName string) SpriteController {
	s.audio = fileName
	return s
}
func (s *spriteControllerData) SetToLoop(repeat bool) SpriteController {
	s.repeat = repeat
	return s
}
func (s *spriteControllerData) StopSprite() SpriteController {
	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audio, false))

	return s
}

// TODO Need sanity checks on empty string names for audio and image. Perhaps on receiver side as well.
func (s *spriteControllerData) RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool) {

	s.repeatFrame++
	if s.repeatFrame > 4 {
		s.currentFrame++
		s.repeatFrame = 0
	}
	if s.currentFrame >= s.maxFrame && s.repeat {
		s.currentFrame = 1
	} else if s.currentFrame >= s.maxFrame && !s.repeat {
		s.currentFrame = s.maxFrame
	}

	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplate, s.currentFrame), p.GetX(), p.GetY(), s.width, s.height)
	log.Default().Println(layer.GetImage())
	layer.SetFlipHorizontal(flipHorizontal)
	drawEvent.AddImageLayer(layer)
	bus.GetVorpalBus().SendDrawEvent(drawEvent)

	bus.GetVorpalBus().SendAudioEvent(bus.NewAudioEvent(s.audio, true))

}
