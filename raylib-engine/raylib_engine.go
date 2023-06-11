package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

//T

type engine struct {
	controller    bus.StandardMediaPeerController
	bus           bus.VorpalBus
	cache         MediaCache
	currentEvtId  int32
	renderTexture rl.RenderTexture2D
}

//Need to disambiguate the controller and bus uses as one call the controller as if its't the bus and that's undesirable.

func NewEngine() bus.Engine {
	log.Println("Init'd")
	var e = engine{}
	e.cache = NewMediaCache()
	e.controller = bus.NewGameController()
	e.bus = bus.GetVorpalBus()
	e.currentEvtId = -1

	return &e
}
func (e *engine) Start() {
	rl.InitWindow(1920, 1080, "Get Window Title from Event!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	rl.InitAudioDevice()
	for !rl.WindowShouldClose() {
		e.cacheImages()
		e.sendMouseEvents()
		e.sendKeyEvents()
		rl.BeginDrawing()
		// if e.controller.GetDrawEvent() != nil && e.currentEvtId != e.controller.GetDrawEvent().GetId() {

		// 	e.currentEvtId = e.controller.GetDrawEvent().GetId()
		// 	e.drawImages()
		// }
		// rl.DrawTexture(e.renderTexture.Texture, 0, 0, rl.RayWhite)
		for _, img := range e.controller.GetDrawEvent().GetImageLayers() {
			if img != nil {
				currentImg := e.cache.GetImage(img.GetImage())
				rl.DrawTexture(rl.LoadTextureFromImage(currentImg), img.GetX(), img.GetY(), rl.RayWhite)

			}
		}
		e.showText()
		rl.EndDrawing()
		e.playAudio()
	}
}

func (e *engine) drawImages() {
	imageLayers := e.controller.GetDrawEvent().GetImageLayers()
	rl.UnloadRenderTexture(e.renderTexture)
	e.renderTexture = rl.LoadRenderTexture(1920, 1080)
	rl.BeginTextureMode(e.renderTexture)
	for _, img := range imageLayers {
		if img != nil {
			currentImg := e.cache.GetImage(img.GetImage())
			rl.DrawTexture(rl.LoadTextureFromImage(currentImg), img.GetX(), img.GetY(), rl.RayWhite)

		}
	}
	rl.EndTextureMode()
}

// TODO Add font and color. Color must haave a controller/Volrpal version that
// is translated here to the raylib.
func (e *engine) showText() {
	evt := e.controller.GetTextEvent()

	if evt != nil {
		rl.DrawText(evt.GetText(), evt.GetX(), evt.GetY(), 20, rl.Black)
	}
}

func (e *engine) playAudio() {
	evt := e.controller.GetAudioEvent()
	if evt != nil {
		var audio = rl.LoadSound(evt.GetAudio())
		rl.PlaySound(audio)

	}
}
func (e *engine) cacheImages() {
	e.cache.CacheImages(e.controller.GetDrawEvent())

}

// TODO This needs to change to use multiple images...
// TODO This eliminated preloading of images in favor of
// of lazy loading, scaling and then storing for further use
// so can be better consolidated later along with multiple layered
// drawing and background color.
// func (e *engine) drawImages() {
// 	imageLayers := e.controller.GetDrawEvent().GetImageLayers()
// 	rl.UnloadRenderTexture(e.renderTexture)
// 	e.renderTexture = rl.LoadRenderTexture(1920, 1080)
// 	rl.BeginTextureMode(e.renderTexture)
// 	rl.BeginDrawing()
// 	for _, img := range imageLayers {
// 		if img != nil {
// 			texture := rl.LoadTextureFromImage(e.cache.GetImage(img.GetImage())) //not sure why this can't be done inside the append???
// 			rl.DrawTexture(texture, img.GetX(), img.GetY(), rl.RayWhite)

// 		}
// 	}
// 	rl.EndTextureMode()
// 	rl.EndDrawing()

// }

// TODO Rethink the mouse event as it probably should be static...
// Need enums for values...
func (e *engine) sendMouseEvents() {
	e.bus.SendMouseEvent(bus.NewMouseEvent(getMouseButton(rl.MouseLeftButton, "Left"), getMouseButton(rl.MouseMiddleButton, "Center"), getMouseButton(rl.MouseRightButton, "Right"), int32(rl.GetMouseX()), int32(rl.GetMouseY())))

}

func getMouseButton(button int32, buttonName string) bus.MouseButtonState {
	return bus.NewMouseButtonState(buttonName, rl.IsMouseButtonDown(button), rl.IsMouseButtonUp(button))
}

// TODO Need registration mechanism to tell which keys to listen for...
func (e *engine) sendKeyEvents() {
	if rl.IsKeyReleased(rl.KeyS) {
		e.bus.SendKeyEvent(bus.NewKeyEvent(bus.GetKeyByString("s")))
	} else if rl.IsKeyReleased(rl.KeyN) {
		e.bus.SendKeyEvent(bus.NewKeyEvent(bus.GetKeyByString("n")))
	}
}
