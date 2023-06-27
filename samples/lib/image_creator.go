package lib

import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

type imageCreatorData struct {
	fileName string
	pipeline ImageRendererPipeline
}

type ImageCreator interface {
	CreateImage(mouseEvent bus.MouseEvent, currentLocation CurrentLocation, frameData FrameData) bus.ImageMetadata
	AddRenderer(ImageDataRenderer) ImageCreator
}

func (icd *imageCreatorData) CreateImage(mouseEvent bus.MouseEvent, currentLocation CurrentLocation, frameData FrameData) bus.ImageMetadata {
	return NewImageMetadata(icd.fileName, currentLocation.GetX(), currentLocation.GetY())
}

func (icd *imageCreatorData) AddRenderer(renderer ImageDataRenderer) ImageCreator {
	icd.pipeline.AddDataRenderer(renderer)
	return icd
}

type sequenceData struct {
	imageCreatorData
}

func (icd *sequenceData) CreateImage(mouseEvent bus.MouseEvent, currentLocation CurrentLocation, frameData FrameData) bus.ImageMetadata {
	img := NewImageMetadata(fmt.Sprintf(icd.imageCreatorData.fileName+"%d", frameData.GetCurrentFrame()), currentLocation.GetX(), currentLocation.GetY())
	icd.pipeline.Render(img, mouseEvent, currentLocation)
	return img
}

func NewImageSequenceCreator(fileName string) ImageCreator {
	icd := imageCreatorData{fileName, &imgageRenderPipelineData{}}
	return &sequenceData{icd}
}
