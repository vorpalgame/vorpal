package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

func NewEngine() bus.Engine {
	log.Println("Init'd")
	e := engine{}
	e.MediaCache = NewMediaCache()
	e.StandardMediaPeerController = bus.NewGameController()
	e.currentlyPlayingAudio = make(map[string]*rl.Sound)
	e.VorpalBus = bus.GetVorpalBus()
	e.DrawEventProcessor = NewDrawEventProcessor(e.MediaCache)
	e.ControlEventProcessor = NewControlEventProcessor()
	e.TextEventProcessor = NewTextEventProcessor(e.MediaCache)
	e.AudioEventProcessor = NewAudioEventProcessor(e.MediaCache)
	return &e
}

type engine struct {
	bus.VorpalBus
	MediaCache
	bus.StandardMediaPeerController
	bus.DrawEventProcessor
	bus.ControlEventProcessor
	bus.TextEventProcessor
	bus.AudioEventProcessor
	currentTexture        rl.Texture2D
	currentlyPlayingAudio map[string]*rl.Sound
}

// TODO window init, size, title should come from control event...
func (e *engine) Start() {
	rl.InitWindow(1920, 1080, "Get Window Title from Event!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.InitAudioDevice()

	for !rl.WindowShouldClose() {
		e.sendMouseEvents()
		e.sendKeyEvents()
		e.ProcessControlEvent(e.GetControlEvent())
		e.ProcessDrawEvent(e.GetDrawEvent())
		e.ProcessTextEvent(e.GetTextEvent())
		e.ProcessAudioEvent(e.GetAudioEvent())
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		e.renderTexture()
		rl.DrawTexture(e.currentTexture, 0, 0, rl.RayWhite)
		rl.EndDrawing()

	}
}

func (e *engine) renderTexture() {
	//Make sure we don't get a race on nil check.
	renderImg := e.GetCurrentRenderImage()
	if renderImg != nil {
		previousTexture := e.currentTexture
		e.currentTexture = rl.LoadTextureFromImage(renderImg)
		rl.UnloadTexture(previousTexture)
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
