package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

//T

var e = engine{}

type engine struct {
	controller   bus.StandardMediaPeerController
	bus          bus.VorpalBus
	imageLibrary map[string]*rl.Image
	textures     []rl.Texture2D
}

//Need to disambiguate the controller and bus uses as one call the controller as if its't the bus and that's undesirable.

func NewEngine() bus.Engine {
	log.Println("Init'd")

	e.imageLibrary = make(map[string]*rl.Image)
	e.textures = make([]rl.Texture2D, 0, 100)
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

		rl.BeginDrawing()
		e.loadImages()
		e.drawImages()
		e.showText()
		rl.EndDrawing()
		e.playAudio()
	}
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
func (e *engine) loadImages() {
	//Continue looping and loading as necessary.
	//TODO need mechanism to load and unload.

	drawEvents := e.controller.GetImageDrawEvents()

	for _, evt := range drawEvents {
		img := e.imageLibrary[evt.GetImage()]
		if img == nil {
			e.imageLibrary[evt.GetImage()] = rl.LoadImage(evt.GetImage())
			img = e.imageLibrary[evt.GetImage()]
			rl.ImageResize(img, evt.GetWidth(), evt.GetHeight())

			e.textures = append(e.textures, rl.LoadTextureFromImage(img))
		}
	}

}

// TODO This needs to change to use multiple images...
// TODO This eliminated preloading of images in favor of
// of lazy loading, scaling and then storing for further use
// so can be better consolidated later along with multiple layered
// drawing and background color.
func (e *engine) drawImages() {
	//Need to store coordinate and other information along with the texture as necessary.

	for _, texture := range e.textures {
		log.Default().Print(texture.Height)
		rl.DrawTexture(texture, 0, 0, rl.LightGray)

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
