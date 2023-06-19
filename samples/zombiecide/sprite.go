package zombiecide

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

type SpriteController interface {
	RunSprite(drawEvent bus.DrawEvent, p Point, flipHorizontal bool)
	StopSprite() SpriteController
	SetAudio(fileName string) SpriteController
	SetToLoop(repeat bool) SpriteController
}

type spriteControllerData struct {
	currentFrame, maxFrame, repeatFrame int
	width, height                       int32
	fileTemplateName                    string
	repeat                              bool
	audioName                           string
	bus                                 bus.VorpalBus
}

// TODO regularize this constructor to either take the audio event or to remove the file and use builder pattern only.
func NewSpriteController(maxFrame, repeatFrame int, width, height int32, fileTemplateName string) SpriteController {
	return &spriteControllerData{1, maxFrame, repeatFrame, width, height, fileTemplateName, true, "", bus.GetVorpalBus()} //Loop by default//no audio by default
}

func (s *spriteControllerData) SetAudio(fileName string) SpriteController {
	s.audioName = fileName
	return s
}
func (s *spriteControllerData) SetToLoop(repeat bool) SpriteController {
	s.repeat = repeat
	return s
}
func (s *spriteControllerData) StopSprite() SpriteController {
	if s.audioName != "" {
		s.bus.SendAudioEvent(bus.NewAudioEvent(s.audioName, false))
	}
	return s
}
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
	//TODO Need a better template file name mechanism.
	layer := bus.NewImageLayer(fmt.Sprintf(s.fileTemplateName, s.currentFrame), p.GetX(), p.GetY(), s.width, s.height)
	layer.SetFlipHorizontal(flipHorizontal)
	drawEvent.AddImageLayer(layer)
	s.bus.SendDrawEvent(drawEvent)
	//Can probably cache the audio event between calls.
	if s.audioName != "" {
		s.bus.SendAudioEvent(bus.NewAudioEvent(s.audioName, true))
	}

}
