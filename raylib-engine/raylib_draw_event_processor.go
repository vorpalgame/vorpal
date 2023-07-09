package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

// /////////////////////////////////////////////////////////////////
// /// Draw Event Processor
// /////////////////////////////////////////////////////////////////

type drawData struct {
	MediaCache
}

func NewDrawEventProcessor(mediaCache MediaCache) bus.DrawEventProcessor {
	return &drawData{mediaCache}
}

func (dep *drawData) ProcessDrawEvent(evt bus.DrawEvent) {

	if evt != nil {
		switch evt := evt.(type) {
		case bus.DrawLayersEvent:
			dep.CacheImages(evt)
			dep.renderImageLayers(evt)
		}
	}
}

// /////////////////////////////////////////////////////////////////
// /// Draw Layer Event rendering
// /////////////////////////////////////////////////////////////////
func (e *drawData) renderImageLayers(evt bus.DrawLayersEvent) {

	for _, layer := range evt.GetImageLayers() {
		if !e.isReady(layer) {
			return
		}
	}
	var renderImg *rl.Image

	for _, layer := range evt.GetImageLayers() {
		renderImg = e.renderLayer(renderImg, layer)
	}

	if renderImg != nil {
		previousImg := e.GetCurrentRenderImage()
		e.SetCurrentRenderImage(renderImg)
		if previousImg != nil {
			rl.UnloadImage(previousImg)
		}
	}
}
func (e *drawData) isReady(layer lib.ImageLayer) bool {

	for _, imgData := range layer.GetLayerData() {
		if e.GetImage(imgData.GetFileName()) == nil {
			return false
		}
	}
	return true
}
func (e *drawData) renderLayer(baseImg *rl.Image, layer lib.ImageLayer) *rl.Image {
	for _, img := range layer.GetLayerData() {
		originalImg := e.GetImage(img.GetFileName())

		//If at any point an image is not loaded and ready, we bail out for this frame.
		if originalImg != nil {

			clonedImage := rl.ImageCopy(originalImg)

			if baseImg == nil {
				baseImg = clonedImage
				//For testing
				//rl.ImageClearBackground(baseImg, rl.White)
			} else {

				if img.IsFlipHorizontal() {
					rl.ImageFlipHorizontal(clonedImage)
				}
				//Create generic
				x, y := img.GetPoint()
				destRect := rl.NewRectangle(float32(x), float32(y), float32(clonedImage.Width), float32(clonedImage.Height))
				rl.ImageDraw(baseImg, clonedImage, rl.NewRectangle(0, 0, float32(clonedImage.Width), float32(clonedImage.Height)), destRect, rl.RayWhite)
				rl.UnloadImage(clonedImage)
			}

		} else {
			return nil
		}

	}
	return baseImg
}
