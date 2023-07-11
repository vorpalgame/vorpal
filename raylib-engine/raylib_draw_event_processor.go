package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

// /////////////////////////////////////////////////////////////////
// /// Draw Event Processor
// /////////////////////////////////////////////////////////////////

var raylibProcessDrawEvent = func(evt bus.DrawEvent, cache MediaCache) {

	if evt != nil {
		switch evt := evt.(type) {
		case bus.DrawLayersEvent:
			cache.CacheImages(evt)
			renderImageLayers(evt, cache)
		}
	}
}

// /////////////////////////////////////////////////////////////////
// /// Draw Layer Event rendering
// /////////////////////////////////////////////////////////////////
func renderImageLayers(evt bus.DrawLayersEvent, cache MediaCache) {

	for _, layer := range evt.GetImageLayers() {
		if !isReady(layer, cache) {
			return
		}
	}
	var renderImg *rl.Image

	for _, layer := range evt.GetImageLayers() {
		renderImg = renderLayer(renderImg, layer, cache)
	}

	if renderImg != nil {
		previousImg := cache.GetCurrentRenderImage()
		cache.SetCurrentRenderImage(renderImg)
		if previousImg != nil {
			rl.UnloadImage(previousImg)
		}
	}
}
func isReady(layer lib.ImageLayerData, cache MediaCache) bool {

	for _, imgData := range layer.LayerMetadata {
		if cache.GetImage(imgData.ImageFileName) == nil {
			return false
		}
	}
	return true
}
func renderLayer(baseImg *rl.Image, layer lib.ImageLayerData, cache MediaCache) *rl.Image {
	for _, img := range layer.LayerMetadata {
		originalImg := cache.GetImage(img.ImageFileName)

		//If at any point an image is not loaded and ready, we bail out for this frame.
		if originalImg != nil {

			clonedImage := rl.ImageCopy(originalImg)

			if baseImg == nil {
				baseImg = clonedImage
				//For testing
				//rl.ImageClearBackground(baseImg, rl.White)
			} else {

				if img.HorizontalFlip {
					rl.ImageFlipHorizontal(clonedImage)
				}
				//Create generic

				destRect := rl.NewRectangle(float32(img.X), float32(img.Y), float32(img.Width), float32(img.Height))
				rl.ImageDraw(baseImg, clonedImage, rl.NewRectangle(0, 0, float32(clonedImage.Width), float32(clonedImage.Height)), destRect, rl.RayWhite)
				rl.UnloadImage(clonedImage)
			}

		}

	}
	return baseImg
}
