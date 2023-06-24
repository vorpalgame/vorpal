package lib

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

func NewImageRenderer() ImageRenderer {
	return &imageControllerData{}
}

type ImageRenderer interface {
	CreateImageLayer(mouseEvent bus.MouseEvent, fd FrameData, cl CurrentLocation) bus.ImageLayer
	SetImageName(imageName string) ImageRenderer
	SetScale(percent int32) ImageRenderer //Could use floats instead of whole number percent....

}

type imageControllerData struct {
	imageName string
	scale     int32
}

func (s *imageControllerData) SetImageName(imageName string) ImageRenderer {
	s.imageName = imageName
	return s
}

func (s *imageControllerData) SetScale(percent int32) ImageRenderer {
	s.scale = percent
	return s
}

func (s *imageControllerData) CreateImageLayer(mouseEvent bus.MouseEvent, fd FrameData, cl CurrentLocation) bus.ImageLayer {
	imgData := bus.NewImageMetadata(fmt.Sprintf(s.imageName, fd.GetCurrentFrame()), cl.GetX(), cl.GetY(), s.scale)
	imgData.SetFlipHorizontal(s.flipHorizontal(mouseEvent, cl))
	layer := bus.NewImageLayer().AddLayerData(imgData)

	fd.Increment()
	return layer
}

func (s *imageControllerData) flipHorizontal(mouseEvent bus.MouseEvent, cl CurrentLocation) bool {
	return mouseEvent.GetX() < cl.GetX()
}
func (s *imageControllerData) getScale() int32 {
	return s.scale
}
