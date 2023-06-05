package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

var e = engine{}

type engine struct {
	controller   bus.StandardGameController
	bus          bus.VorpalBus
	imageLibrary map[string]*rl.Image
}

//Need to disambiguate the controller and bus uses as one call the controller as if its't the bus and that's undesirable.

func NewEngine() bus.LifeCycle {
	log.Println("Init'd")

	e.imageLibrary = make(map[string]*rl.Image)
	e.controller = bus.NewGameController()
	e.bus = bus.GetVorpalBus()

	return &e
}
func (e *engine) Start() {
	rl.InitWindow(1920, 1080, "Get Window Title from Event!")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		e.sendMouseEvents()
		e.sendKeyEvents()
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		//rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		e.drawImages()
		rl.EndDrawing()
	}
}

// TODO This needs to change to use multiple images...
// TODO This eliminated preloading of images in favor of
// of lazy loading, scaling and then storing for further use
// so can be better consolidated later along with multiple layered
// drawing and background color.
func (e *engine) drawImages() {

	evt := e.controller.GetImageDrawEvents()
	if evt != nil {
		var img = e.imageLibrary[evt.GetImage()]
		if img == nil {
			e.imageLibrary[evt.GetImage()] = rl.LoadImage(evt.GetImage())
			img = e.imageLibrary[evt.GetImage()]
			rl.ImageResize(img, evt.GetWidth(), evt.GetHeight())
		} else { //Give it time to load the image...
			texture := rl.LoadTextureFromImage(img)
			rl.DrawTexture(texture, evt.GetX(), evt.GetY(), rl.White)
		}
	}

}

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
