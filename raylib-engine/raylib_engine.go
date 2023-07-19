package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
	"log"
	"unicode"
)

func NewEngine() lib.Engine {

	e := engine{}

	e.VorpalBus = bus.GetVorpalBus()
	e.RaylibPeerController = NewRaylibPeerController()
	e.MediaCache = NewMediaCache()
	e.VorpalBus.AddDrawEventListener(e.RaylibPeerController)
	e.VorpalBus.AddControlEventListener(e.RaylibPeerController)
	e.VorpalBus.AddAudioEventListener(e.RaylibPeerController)
	e.VorpalBus.AddKeysRegistrationEventListener(e.RaylibPeerController)
	e.VorpalBus.AddTextEventListener(e.RaylibPeerController)

	return &e
}

type engine struct {
	bus.VorpalBus
	RaylibPeerController
	MediaCache
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
		go raylibProcessAudioEvent(e.GetAudioEvent(), e.MediaCache)
		//Baton pass between threads...

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		tx := NewRenderTransaction(e.GetDrawEvent(), e.GetTextEvent(), e.MediaCache, e.CurrentTexture)
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

func (e *engine) sendKeyEvents() {

	if e.GetKeysRegistrationEvent() != nil {
		log.Println(e.GetKeysRegistrationEvent())
		for _, r := range e.GetKeysRegistrationEvent().GetRunes() {
			//TODO Let's wire a better mechanism for this...and pull it from the registration event.
			if rl.IsKeyPressed(r) || rl.IsKeyPressed(unicode.ToUpper(r)) {
				evt := key.Event{}
				evt.Rune = r
				evt.Direction = key.DirPress
				evt.Code = key.Code(r)
				e.SendKeyEvent(bus.NewKeyEvent(evt))
			}
		}
	}
}
