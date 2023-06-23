package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

//T

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
		drawEvt := e.controller.GetDrawEvent()
		textEvt := e.controller.GetTextEvent()
		rl.BeginDrawing()
		if drawEvt != nil {
			e.cache.CacheImages(drawEvt)
			e.renderImages(drawEvt)
		}

		if textEvt != nil {
			e.cache.CacheFonts(textEvt)
			e.renderText(textEvt)
		}
		e.renderTexture()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawTexture(e.currentTexture, 0, 0, rl.RayWhite)
		rl.EndDrawing()

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
	var isReady bool
	for _, layer := range evt.GetImageLayers() {
		isReady = e.isReady(layer)
		if !isReady {
			break
		}
	}
	var baseImg *rl.Image
	if isReady {
		for _, layer := range evt.GetImageLayers() {
			baseImg = e.renderLayer(baseImg, layer)
		}
	}

	if baseImg != nil {
		previousImg := e.renderedImg
		e.renderedImg = baseImg
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

func getMouseButton(button int32, buttonName string) bus.MouseButtonState {
	//We have an error on occasion wher the mouse button is nil and it isn't clear why or how...May be start up race or bug.
	return bus.NewMouseButtonState(buttonName, rl.IsMouseButtonDown(button))
}

func (e *engine) sendKeyEvents() {
	if e.controller.GetKeysRegistrationEvent() != nil {
		for _, key := range e.controller.GetKeysRegistrationEvent().GetKeys() {
			if rl.IsKeyReleased(int32(key.ToAscii())) {
				e.bus.SendKeyEvent(bus.NewKeyEvent(key))
			}
		}
	}
}
