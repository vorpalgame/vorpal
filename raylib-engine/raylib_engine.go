package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
)

func NewEngine() lib.Engine {

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
	evt := bus.NewMouseEvent(getMouseEvent())
	e.SendMouseEvent(evt)

}

// TODO We need to complete the mapping from rl mouse to mouse.Event
func getMouseEvent() mouse.Event {
	evt := mouse.Event{}
	evt.X = float32(rl.GetMouseX())
	evt.Y = float32(rl.GetMouseY())
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		evt.Button = mouse.ButtonLeft
		evt.Direction = mouse.DirPress
	}
	return evt
}

// TODO Revamp
func (e *engine) sendKeyEvents() {

	if e.GetKeysRegistrationEvent() != nil {
		for _, r := range e.GetKeysRegistrationEvent().GetRunes() {
			//We currently pass as "string" so have to fish out the first letter. Refactor later.

			if rl.IsKeyPressed(r) {
				evt := key.Event{}
				evt.Rune = r
				evt.Direction = key.DirPress
				evt.Code = key.Code(r)
				e.SendKeyEvent(bus.NewKeyEvent(evt))
			}
		}
	}
}
