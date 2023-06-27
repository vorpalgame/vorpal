package lib

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

func NewImageRendererPipeline() ImageRendererPipeline {
	return &imgageRenderPipelineData{}
}

type ImageDataRenderer interface {
	Render(imgData bus.ImageMetadata, mouseEvent bus.MouseEvent, cl CurrentLocation)
}
type ImageRendererPipeline interface {
	ImageDataRenderer
	AddDataRenderer(ImageDataRenderer)
}

type imgageRenderPipelineData struct {
	ir   ImageDataRenderer
	next ImageRendererPipeline
}

func (irp *imgageRenderPipelineData) AddDataRenderer(ir ImageDataRenderer) {
	if irp.ir == nil {
		irp.ir = ir
	} else if irp.next == nil {
		irp.next = &imgageRenderPipelineData{ir, nil}
	} else {
		irp.next.AddDataRenderer(ir)
	}
}
func (irp *imgageRenderPipelineData) Render(imgData bus.ImageMetadata, mouseEvent bus.MouseEvent, cl CurrentLocation) {
	if irp.next != nil {
		irp.next.Render(imgData, mouseEvent, cl)
	}
	irp.ir.Render(imgData, mouseEvent, cl)
}

type FlipHorizontalTransformer interface {
	ImageDataRenderer
}

// Anyway to create stateless instances of interfaces???
type flipData struct {
}

func NewFlipHorizontalTransformer() FlipHorizontalTransformer {
	return &flipData{}
}
func (fd *flipData) Render(imgData bus.ImageMetadata, mouseEvent bus.MouseEvent, cl CurrentLocation) {
	imgData.SetFlipHorizontal(mouseEvent.GetX() < cl.GetX())
}

// Scale set to 1 as default...
func NewImageMetadataSequence(imageFileName string, sequenceNumber, x, y int32) bus.ImageMetadata {
	return NewImageMetadata(fmt.Sprintf(imageFileName, sequenceNumber), x, y)
}

func NewImageMetadata(imageFileName string, x, y int32) bus.ImageMetadata {
	return bus.NewImageMetadata(imageFileName, x, y, int32(1))
}
