package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

// TODO Work out engine initizliation sequence with
// window size, title, location, etc. to eliminate hard coded values.
func NewEngine() bus.Engine {

	e := engine{}
	e.MediaCacheData = NewMediaCache()
	e.StandardMediaPeerController = bus.NewGameController()
	e.VorpalBus = bus.GetVorpalBus()
	return &e
}

type engine struct {
	bus.VorpalBus
	MediaCacheData
	bus.StandardMediaPeerController
}

func (e *engine) Start() {
	//These need to be set from events...
	rl.InitWindow(1920, 1080, "")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	for !rl.WindowShouldClose() {
		//async safe.
		go e.sendMouseEvents()
		go e.sendKeyEvents()
		go raylibProcessControlEvent(e.GetControlEvents())
		go raylibProcessAudioEvent(e.GetAudioEvent(), &e.MediaCacheData)

		//Process these on thread...
		raylibProcessDrawEvent(e.GetDrawEvent(), &e.MediaCacheData)
		raylibProcessTextEvent(e.GetTextEvent(), &e.MediaCacheData)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		e.renderTexture()
		if e.GetCurrentTexture() != nil {
			rl.DrawTexture(*e.GetCurrentTexture(), 0, 0, rl.RayWhite)
		} //TODO not sure about the white background...
		rl.EndDrawing()

	}
}

func (e *engine) renderTexture() {
	renderImg := e.GetCurrentRenderImage()
	if renderImg != nil {
		previousTexture := e.CurrentTexture
		newTexture := rl.LoadTextureFromImage(renderImg)
		e.SetCurrentTexture(&newTexture)
		if previousTexture != nil {
			rl.UnloadTexture(*previousTexture)
		}
	}
}

func (e *engine) sendMouseEvents() {
	evt := bus.NewMouseEvent(getMouseButton(rl.MouseLeftButton, "Left"), getMouseButton(rl.MouseMiddleButton, "Center"), getMouseButton(rl.MouseRightButton, "Right"), int32(rl.GetMouseX()), int32(rl.GetMouseY()))
	e.SendMouseEvent(evt)

}

func getMouseButton(button int32, buttonName string) lib.MouseButtonState {
	return lib.NewMouseButtonState(buttonName, rl.IsMouseButtonDown(button))
}

func (e *engine) sendKeyEvents() {
	if e.GetKeysRegistrationEvent() != nil {
		for _, key := range e.GetKeysRegistrationEvent().GetKeys() {
			if rl.IsKeyPressed(key.ToAscii()) {
				e.SendKeyEvent(bus.NewKeyEvent(key))
			}
		}
	}
}
