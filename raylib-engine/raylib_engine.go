package raylibengine

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

//T

type engine struct {
	controller     bus.StandardMediaPeerController
	bus            bus.VorpalBus
	cache          MediaCache
	currentEvtId   int32
	currentTextId  int32
	renderedImg    *rl.Image
	currentTexture rl.Texture2D
}

//

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

	//e.currentFont = rl.LoadFontEx(fontName, 12, make([]rune, 95))
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.RayWhite)
		e.cacheResources()
		e.sendMouseEvents()
		e.sendKeyEvents()
		rl.BeginDrawing()
		e.renderImages()
		e.renderText()
		rl.DrawTexture(e.currentTexture, 0, 0, rl.RayWhite)
		rl.EndDrawing()
		e.playAudio()
	}
}

func (e *engine) renderImages() {
	if e.controller.GetDrawEvent() != nil && e.controller.GetDrawEvent().GetId() != e.currentEvtId {

		if e.renderedImg != nil {
			rl.UnloadImage(e.renderedImg)
		}

		e.currentEvtId = e.controller.GetDrawEvent().GetId()
		var baseImg *rl.Image
		for _, img := range e.controller.GetDrawEvent().GetImageLayers() {
			if img != nil {
				currentImg := rl.ImageCopy(e.cache.GetImage(img.GetImage()))
				if baseImg == nil {
					baseImg = currentImg
				} else {
					rl.ImageDraw(baseImg, currentImg, rl.NewRectangle(0, 0, float32(currentImg.Width), float32(currentImg.Height)), rl.NewRectangle(float32(img.GetX()), float32(img.GetY()), float32(currentImg.Width), float32(currentImg.Height)), rl.RayWhite)

				}

			}

		}
		e.renderedImg = baseImg
		rl.UnloadTexture(e.currentTexture)
		e.currentTexture = rl.LoadTextureFromImage(baseImg)
	}
}

func (e *engine) renderText() {
	txtEvt := e.controller.GetTextEvent()
	if txtEvt != nil && txtEvt.GetId() != e.currentTextId {
		e.currentTextId = txtEvt.GetId()
		baseImg := rl.ImageCopy(e.renderedImg)
		//TODO The lines will not be wrapped here so this is temporary
		//The next step is to send each presplit line from the other side of the bus
		//and then iterate over it here.
		//
		var tempTxt string
		for _, txt := range txtEvt.GetText() {
			log.Default().Println(tempTxt)
			tempTxt = tempTxt + txt.GetText() + "\n"
		}
		log.Default().Println(tempTxt)
		rl.ImageDrawTextEx(baseImg, rl.Vector2{float32(txtEvt.GetX()), float32(txtEvt.GetY())}, *e.cache.GetFont(txtEvt.GetFont()), tempTxt, float32(txtEvt.GetFontSize()), 0, rl.Black)
		rl.UnloadTexture(e.currentTexture)
		e.currentTexture = rl.LoadTextureFromImage(baseImg)
	}
}

func (e *engine) playAudio() {
	evt := e.controller.GetAudioEvent()
	if evt != nil {
		var audio = rl.LoadSound(evt.GetAudio())
		rl.PlaySound(audio)

	}
}
func (e *engine) cacheResources() {
	e.cache.CacheFonts(e.controller.GetTextEvent())
	e.cache.CacheImages(e.controller.GetDrawEvent())

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
