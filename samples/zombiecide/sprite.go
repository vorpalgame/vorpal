package zombiecide

import (
	"strconv"

	"github.com/vorpalgame/vorpal/bus"
)

type Sprite interface {
	RenderNext(p Point, flipHorizontal bool) *bus.ImageLayer
	SetToLoop(repeat bool)
}

type sprite struct {
	currentFrame, maxFrame, repeatFrame int
	width, height                       int32
	fileBaseName                        string
	repeat                              bool
}

func NewSprite(maxFrame, repeatFrame int, width, height int32, fileBaseName string) Sprite {
	return &sprite{1, maxFrame, repeatFrame, width, height, fileBaseName, true} //Loop by default
}
func (s *sprite) SetToLoop(repeat bool) {
	s.repeat = repeat
}
func (s *sprite) RenderNext(p Point, flipHorizontal bool) *bus.ImageLayer {

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
	//TODO Need a substitution mechanism for this...
	layer := bus.NewImageLayer("samples/resources/zombiecide/"+s.fileBaseName+" ("+strconv.Itoa(s.currentFrame)+").png", p.GetX(), p.GetY(), s.width, s.height)
	layer.SetFlipHorizontal(flipHorizontal)
	return &layer
}
