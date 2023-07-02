package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

//TODO We now have a draw event with multiple layers and each layer can have multiple
//renderable components. It may be desirable to render all the components in a layer onto
//a single blank image layer that can be used. That would imply we can cache that rendered
//layer image based on unique tag and not rerender it every time.
//Currently we pre-scale all the images but we wil need to implement the cache flush
//implementation or enable flushing specific elements. TBD

type engine struct {
	controller            bus.StandardMediaPeerController
	bus                   bus.VorpalBus
	cache                 MediaCache
	renderedImg           *rl.Image
	currentTexture        rl.Texture2D
	currentlyPlayingAudio map[string]*rl.Sound
}

//

func NewEngine() bus.Engine {
	log.Println("Init'd")
	var e = engine{}
	e.currentlyPlayingAudio = make(map[string]*rl.Sound)
	e.cache = NewMediaCache()
	e.controller = bus.NewGameController()
	e.bus = bus.GetVorpalBus()

	return &e
}
func (e *engine) Start() {
	rl.InitWindow(1920, 1080, "Get Window Title from Event!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	for !rl.WindowShouldClose() {
		e.sendMouseEvents()
		e.sendKeyEvents()
		e.runAudio()
		e.processControlEvent()

		drawEvt := e.controller.GetDrawEvent()
		textEvt := e.controller.GetTextEvent()

		if drawEvt != nil {
			e.cache.CacheImages(drawEvt)
			e.renderImages(drawEvt)
			//We draw text only if image is actually being rednered...
			if textEvt != nil {
				e.cache.CacheFonts(textEvt)
				e.renderText(textEvt)
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		e.renderTexture()

		rl.DrawTexture(e.currentTexture, 0, 0, rl.RayWhite)
		rl.EndDrawing()

	}
}

func (e *engine) processControlEvent() {
	controlEvt := e.controller.GetControlEvent()
	if controlEvt != nil {
		switch controlEvt := controlEvt.(type) {
		case bus.WindowTitleEvent:
			rl.SetWindowTitle(controlEvt.GetTitle())
		}

	}
}
func (e *engine) renderTexture() {
	if e.renderedImg != nil {
		previousTexture := e.currentTexture
		e.currentTexture = rl.LoadTextureFromImage(e.renderedImg)
		rl.UnloadTexture(previousTexture)
	}
}

func (e *engine) renderImages(evt bus.DrawEvent) {

	//Get each layer and render the 1...N entries of content on the layer
	//Then render the next layer on top of it.

	//Verify everything is ready to render or return.
	//This ensures images are loaded from file before we
	//attempt to use them.
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
		previousImg := e.renderedImg
		e.renderedImg = renderImg
		if previousImg != nil {
			rl.UnloadImage(previousImg)
		}
	}
}
func (e *engine) isReady(layer bus.ImageLayer) bool {

	for _, imgData := range layer.GetLayerData() {
		if e.cache.GetImage(imgData.GetImage()) == nil {
			return false
		}
	}
	return true
}
func (e *engine) renderLayer(baseImg *rl.Image, layer bus.ImageLayer) *rl.Image {
	for _, img := range layer.GetLayerData() {
		originalImg := e.cache.GetImage(img.GetImage())

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
				rl.ImageDraw(baseImg, clonedImage, rl.NewRectangle(0, 0, float32(clonedImage.Width), float32(clonedImage.Height)), rl.NewRectangle(float32(img.GetX()), float32(img.GetY()), float32(clonedImage.Width), float32(clonedImage.Height)), rl.RayWhite)
				rl.UnloadImage(clonedImage)
			}

		} else {
			return nil
		}

	}
	return baseImg
}

func (e *engine) renderText(txtEvt bus.TextEvent) {

	if e.renderedImg != nil {

		//TODO The lines will not be wrapped here so this is temporary
		//The next step is to send each presplit line from the other side of the bus
		//and then iterate over it here.
		//
		x := float32(txtEvt.GetX())
		var y = float32(txtEvt.GetY())
		for _, txt := range txtEvt.GetText() {
			rl.ImageDrawTextEx(e.renderedImg, rl.Vector2{x, y}, *e.cache.GetFont(txt.GetFont()), txt.GetText(), float32(txt.GetFontSize()), 0, rl.Black)
			//How to do line spacing????
			y += float32(txt.GetFontSize()) * float32(1.1) //Extra space..
		}

	}

}

func (e *engine) runAudio() {
	evt := e.controller.GetAudioEvent()

	if evt != nil {
		currentAudio := e.cache.GetAudio(evt.GetAudio())
		if evt.IsStop() {
			rl.StopSound(*currentAudio)
			//rl.UnloadSound(*currentAudio)
		}
		if evt.IsPlay() && !rl.IsSoundPlaying(*currentAudio) {
			for !rl.IsSoundReady(*currentAudio) {
			}
			if evt.IncrementCount() == 1 || evt.IsLoop() {
				rl.PlaySound(*currentAudio)
			}
		}

	}

}

func (e *engine) sendMouseEvents() {

	evt := bus.NewMouseEvent(getMouseButton(rl.MouseLeftButton, "Left"), getMouseButton(rl.MouseMiddleButton, "Center"), getMouseButton(rl.MouseRightButton, "Right"), int32(rl.GetMouseX()), int32(rl.GetMouseY()))
	e.bus.SendMouseEvent(evt)

}

func getMouseButton(button int32, buttonName string) lib.MouseButtonState {
	return lib.NewMouseButtonState(buttonName, rl.IsMouseButtonDown(button))
}

func (e *engine) sendKeyEvents() {
	if e.controller.GetKeysRegistrationEvent() != nil {
		for _, key := range e.controller.GetKeysRegistrationEvent().GetKeys() {
			if rl.IsKeyPressed(key.ToAscii()) {
				e.bus.SendKeyEvent(bus.NewKeyEvent(key))
			}
		}
	}
}
