package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

type RenderTransaction interface {
	GetDrawEvent() bus.DrawEvent
	SetDrawEvent(event bus.DrawEvent)
	GetTextEvent() bus.TextEvent
	SetTextEvent(evt bus.TextEvent)
	GetMediaCache() MediaCache
	SetMediaCache(MediaCache)
	GetRenderImage() *rl.Image
	SetRenderImage(*rl.Image) //Mot used currently but will be in future instead of cache.
	SetRenderTexture(*rl.Texture2D)
	GetRenderTexture() *rl.Texture2D
	SetPreviousTexture(*rl.Texture2D)
	GetPreviousTexture() *rl.Texture2D
}
type renderData struct {
	DrawEvent                      bus.DrawEvent
	TextEvent                      bus.TextEvent
	MediaCache                     MediaCache
	RenderImage                    *rl.Image //Mot used currently but will be in future instead of cache.
	RenderTexture, PreviousTexture *rl.Texture2D
}

func NewRenderTransaction(drawEvent bus.DrawEvent, textEvent bus.TextEvent, mediaCache MediaCache, previousTexture *rl.Texture2D) renderData {
	return renderData{drawEvent, textEvent, mediaCache, nil, nil, previousTexture}
}

type renderingPipelineFunction = func(tx *renderData)

func NewRendererPipeline() RendererPipeline {
	pipeline := RenderPipelineData{}
	pipeline.Add(drawRenderer).Add(raylibProcessTextEvent).Add(rayLibTextureProcessor)
	return &pipeline
}

type RenderPipelineData struct {
	pipeline []renderingPipelineFunction
}

type RendererPipeline interface {
	Add(function renderingPipelineFunction) RendererPipeline
	Execute(transaction *renderData)
}

func (r *RenderPipelineData) Add(function renderingPipelineFunction) RendererPipeline {
	r.pipeline = append(r.pipeline, function)
	return r
}

func (r *RenderPipelineData) Execute(transaction *renderData) {
	//log.Println(r.pipeline)
	for _, function := range r.pipeline {
		function(transaction)
	}
}

// /////////////////////////////////////////////////////////////////
// /// Texture  Processor
// /////////////////////////////////////////////////////////////////
var rayLibTextureProcessor = func(tx *renderData) {

	renderImg := tx.RenderImage

	if renderImg != nil {
		newTexture := rl.LoadTextureFromImage(renderImg)
		tx.RenderTexture = &newTexture

		if tx.PreviousTexture != nil {
			rl.UnloadTexture(*tx.PreviousTexture)
		}

	}
	//log.Default().Println("tx.RenderTexture")
	//log.Default().Println(tx.RenderTexture)
}

// /////////////////////////////////////////////////////////////////
// /// Text Event Processor
// /////////////////////////////////////////////////////////////////

var raylibProcessTextEvent = func(tx *renderData) {

	//Currently the switch isn't differentiating on different
	//subtypes of TextEvent but it will.
	if tx.TextEvent != nil {
		switch evt := tx.TextEvent.(type) {
		case bus.MultilineTextEvent:
			renderText(evt, tx)
		}
	}
}

func renderText(evt bus.MultilineTextEvent, tx *renderData) {
	renderImg := tx.RenderImage

	if renderImg != nil {
		cache.CacheFonts(evt)
		//TODO The lines will not be wrapped here so this is temporary
		//The next step is to send each presplit line from the other side of the bus
		//and then iterate over it here.
		//
		intx, inty := evt.GetLocation()
		var x = float32(intx)
		var y = float32(inty)
		for _, txt := range evt.GetText() {
			cache.CacheFonts(txt)
			rl.ImageDrawTextEx(renderImg, rl.Vector2{x, y}, *cache.GetFont(txt.GetFont()), txt.GetText(), float32(txt.GetFontSize()), 0, rl.Black)
			//How to do line spacing????
			y += float32(txt.GetFontSize()) * float32(1.1) //Extra space..
		}

	}

}

// /////////////////////////////////////////////////////////////////
// /// Draw Layer Event rendering
// /////////////////////////////////////////////////////////////////
var drawRenderer = func(tx *renderData) {

	if tx.DrawEvent != nil {
		switch evt := tx.DrawEvent.(type) {
		case bus.DrawLayersEvent:
			cache.CacheImages(evt)
			renderImageLayers(evt, tx)
		}
	}
}

func renderImageLayers(evt bus.DrawLayersEvent, tx *renderData) {

	for _, layer := range evt.GetImageLayers() {
		if !isReady(layer, tx) {
			return
		}
	}
	var renderImg *rl.Image

	for _, layer := range evt.GetImageLayers() {
		renderImg = renderLayer(renderImg, layer, tx)
	}

	tx.RenderImage = renderImg
	//log.Println("Set TX RenderImage")
	//log.Println(tx.RenderImage)
}
func isReady(layer lib.ImageLayerData, tx *renderData) bool {

	for _, imgData := range layer.LayerMetadata {
		if tx.MediaCache.GetImage(imgData.ImageFileName) == nil {
			return false
		}
	}
	return true
}
func renderLayer(baseImg *rl.Image, layer lib.ImageLayerData, tx *renderData) *rl.Image {
	for _, img := range layer.LayerMetadata {
		originalImg := tx.MediaCache.GetImage(img.ImageFileName)

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
