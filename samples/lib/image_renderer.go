package lib

import (
	"github.com/vorpalgame/vorpal/bus"
)

// TODO We need to refactor to decompose into pipeline segments that are composable.
// Flip horizontal is the first but more to come...
func NewImageRenderer() ImageRenderer {
	ir := &imageRendererData{}
	ir.pipeline = NewImageRendererPipeline()
	ir.pipeline.AddDataRenderer(NewFlipHorizontalTransformer())
	return ir
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
	pipeline  ImageRendererPipeline
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
	//TODO The imgage creation should be in pipleine as well.
	imgData := NewImageMetadataSequence(s.imageName, fd.GetCurrentFrame(), cl.GetX(), cl.GetY()).SetScale(s.scale)
	s.pipeline.Render(imgData, mouseEvent, cl)
	layer := bus.NewImageLayer().AddLayerData(imgData)

	fd.Increment()
	return layer
}

func (s *imageRendererData) getScale() int32 {
	return s.scale
}
