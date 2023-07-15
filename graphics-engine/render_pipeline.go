package util

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/image/draw"
	"image"
)

type pipelineData struct {
	drawEventChannel chan bus.DrawEvent
	outputChannel    chan *image.RGBA
	MediaCache
}

func (data *pipelineData) OnDrawEvent(inputChannel <-chan bus.DrawEvent) {
	for evt := range inputChannel {
		data.drawEventChannel <- evt
	}
}

func NewRenderPipeline(outputChannel chan *image.RGBA) {
	data := pipelineData{make(chan bus.DrawEvent, 1), outputChannel, NewMediaCache()}
	bus.GetVorpalBus().AddDrawEventListener(&data)
	go renderPipelineFunc(&data, data.drawEventChannel)
}

var renderPipelineFunc = func(data *pipelineData, inputChannel <-chan bus.DrawEvent) {
	cacheChan := make(chan bus.DrawLayersEvent, 1)
	go cacheImagesFunc(data, cacheChan)
	for evt := range inputChannel {
		switch evt := evt.(type) {
		case bus.DrawLayersEvent:
			cacheChan <- evt
		}

	}
}

var cacheImagesFunc = func(data *pipelineData, inputChannel <-chan bus.DrawLayersEvent) {
	renderChan := make(chan bus.DrawLayersEvent, 1)
	go renderImageLayersFunc(data, renderChan)
	for evt := range inputChannel {
		data.CacheImages(evt)
		renderChan <- evt
	}

}

var renderImageLayersFunc = func(data *pipelineData, inputChannel <-chan bus.DrawLayersEvent) {

	for evt := range inputChannel {
		var buffer *image.RGBA
		for _, layer := range evt.GetImageLayers() {
			for _, imgData := range layer.LayerMetadata {
				img := *data.GetImage(imgData.ImageFileName)
				if buffer == nil {
					buffer = image.NewRGBA(img.Bounds())
				}
				if imgData.HorizontalFlip {
					img = flip(img)
				}
				draw.Draw(buffer, getRect(imgData), img, *getPoint(0, 0), draw.Over)
			}

		}
		data.outputChannel <- buffer

	}
}

func flip(img image.Image) image.Image {
	size := img.Bounds().Size()
	flipImg := image.NewNRGBA(img.Bounds())
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			xp := size.X - x - 1
			flipImg.Set(x, y, img.At(xp, y))
		}
	}
	return flipImg
}

// ////////////////////////////////////////////////////////////////////////////
// TODO We need these on ImageMetadta interfaces...
func getRect(imgData *lib.ImageMetadata) image.Rectangle {
	return image.Rectangle{*getPoint(imgData.X, imgData.Y), *getPoint(imgData.X+imgData.Width, imgData.Y+imgData.Height)}
}

func getPoint(x, y int32) *image.Point {
	return &image.Point{int(x), int(y)}
}
