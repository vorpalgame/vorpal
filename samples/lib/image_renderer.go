package lib

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

// TODO We need to refactor to decompose into pipeline segments that are composable.
func NewImageRenderer() ImageRenderer {
	return &imageRendererData{}
}

type ImageRenderer interface {
	CreateImageLayer(mouseEvent bus.MouseEvent, fd FrameData, cl CurrentLocation) bus.ImageLayer
	SetImageName(imageName string) ImageRenderer
	SetScale(percent int32) ImageRenderer //Could use floats instead of whole number percent....
	IncrementScale(percent int32) ImageRenderer
}

type imageRendererData struct {
	imageName string
	scale     int32
}

func (s *imageRendererData) SetImageName(imageName string) ImageRenderer {
	s.imageName = imageName
	return s
}
func (s *imageRendererData) IncrementScale(percent int32) ImageRenderer {
	s.scale += percent
	return s
}
func (s *imageRendererData) SetScale(percent int32) ImageRenderer {
	s.scale = percent
	return s
}

// TODO This is coupled to the frame data flip book style animation and shouldn't be.
func (s *imageRendererData) CreateImageLayer(mouseEvent bus.MouseEvent, fd FrameData, cl CurrentLocation) bus.ImageLayer {
	imgData := bus.NewImageMetadata(fmt.Sprintf(s.imageName, fd.GetCurrentFrame()), cl.GetX(), cl.GetY(), s.scale)
	imgData.SetFlipHorizontal(s.flipHorizontal(mouseEvent, cl))
	layer := bus.NewImageLayer().AddLayerData(imgData)

	fd.Increment()
	return layer
}

func (s *imageRendererData) flipHorizontal(mouseEvent bus.MouseEvent, cl CurrentLocation) bool {
	return mouseEvent.GetX() < cl.GetX()
}
func (s *imageRendererData) getScale() int32 {
	return s.scale
}
