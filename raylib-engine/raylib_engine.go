package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

func NewEngine() bus.Engine {

	e := engine{}

	e.StandardMediaPeerController = bus.NewGameController()
	e.VorpalBus = bus.GetVorpalBus()
	return &e
}

var cache = NewMediaCache()

type engine struct {
	bus.VorpalBus
	bus.StandardMediaPeerController
	CurrentRenderImage *rl.Image
	CurrentTexture     *rl.Texture2D
}

func (e *engine) Start() {
	//These need to be set from events...
	rl.InitWindow(1920, 1080, "")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	//When the interface is set for render transaction we can stop using the pointer.

	pipeline := NewRendererPipeline()

	for !rl.WindowShouldClose() {
		//async safe.
		go e.sendMouseEvents()
		go e.sendKeyEvents()
		go raylibProcessControlEvent(e.GetControlEvents())
		go raylibProcessAudioEvent(e.GetAudioEvent(), cache)
		//Baton pass between threads...

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		tx := NewRenderTransaction(e.GetDrawEvent(), e.GetTextEvent(), cache, e.CurrentTexture)
		pipeline.Execute(&tx)
		if tx.RenderTexture != nil {
			e.CurrentTexture = tx.RenderTexture
		}
		if e.CurrentTexture != nil {
			rl.DrawTexture(*e.CurrentTexture, 0, 0, rl.RayWhite)
		}
		rl.EndDrawing()

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
