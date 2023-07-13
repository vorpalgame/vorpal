package native_engine

import (
	"fmt"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"image"
	"image/png"
	"log"
	"os"
)

type engine struct {
	bus.VorpalBus
	bus.StandardMediaPeerController
	currentRenderImage, currentDisplayImage image.Image
	window                                  screen.Window
	screen                                  screen.Screen
}

func NewEngine() lib.Engine {

	log.Println("New native engine...")
	e := engine{}

	e.StandardMediaPeerController = bus.NewGameController()
	e.VorpalBus = bus.GetVorpalBus()
	go e.initWindow()

	return &e
}
func (e *engine) Start() {
	log.Println("Start Native Engine...")
	e.renderLoop()

}
func (e *engine) renderLoop() {
	for {
		e.VorpalBus.SendMouseEvent(bus.NewDefaultMouseEvent(800, 800))

		evt := e.GetDrawEvent()
		if evt != nil {
			log.Println("Loading...")
			log.Println(evt)
			switch evt := evt.(type) {
			case bus.DrawLayersEvent:
				e.renderImageLayers(evt)
				//log.Println(evt)
			}
		}
	}
}
func (e *engine) renderImageLayers(evt bus.DrawLayersEvent) {
	log.Println("Render image...")

	var buffer *image.RGBA
	for _, layer := range evt.GetImageLayers() {
		for _, imgData := range layer.LayerMetadata {
			img := *loadImage(imgData.ImageFileName)
			if buffer == nil {
				buffer = image.NewRGBA(img.Bounds())
			} else {
				toRect := image.Rect(0, 0, int(imgData.Width), int(imgData.Height))
				resizedImage := image.NewRGBA(toRect)
				draw.NearestNeighbor.Scale(resizedImage, resizedImage.Rect, img, img.Bounds(), draw.Over, nil)
				//log.Default().Println(img.Bounds())
				img = resizedImage

			}

			draw.Draw(buffer, getRect(imgData), img, *getPoint(0, 0), draw.Over)

		}
	}
	//write(buffer)

	//os.Exit(1)
	e.blit(buffer)

}

// TODO We need thes on ImageMetadta interfaces...
func getRect(imgData *lib.ImageMetadata) image.Rectangle {
	return image.Rectangle{*getPoint(imgData.X, imgData.Y), *getPoint(imgData.X+imgData.Width, imgData.Y+imgData.Height)}
}

func (e *engine) blit(buffer *image.RGBA) {

	b, _ := e.screen.NewBuffer(buffer.Bounds().Max)
	draw.Draw(b.RGBA(), b.Bounds(), buffer, *getPoint(0, 0), draw.Over)
	log.Println("Updload image...")
	e.window.Upload(image.Point{0, 0}, b, buffer.Bounds())
	e.window.Publish()
}
func getPoint(x, y int32) *image.Point {
	return &image.Point{int(x), int(y)}
}

func write(buffer *image.RGBA) {
	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(out, buffer)
}
func loadImage(imageFileName string) *image.Image {
	log.Println(imageFileName)
	f, err := os.Open(imageFileName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	return &img
}
func (e *engine) doEventLoop() {

	for {
		switch evt := e.window.NextEvent().(type) {
		case lifecycle.Event:
			if evt.To == lifecycle.StageDead {
				return
			}
		case mouse.Event:
			e.VorpalBus.SendMouseEvent(bus.NewDefaultMouseEvent(int32(evt.X), int32(evt.Y)))
		}
	}
}

func (e *engine) initWindow() {
	driver.Main(func(s screen.Screen) {
		e.screen = s
		w, err := s.NewWindow(nil)
		e.window = w
		if err != nil {
			panic(err)
			return
		}
		defer w.Release()

		for {
			switch evt := w.NextEvent().(type) {
			case lifecycle.Event:
				if evt.To == lifecycle.StageDead {
					return
				}
			case mouse.Event:
				e.VorpalBus.SendMouseEvent(bus.NewDefaultMouseEvent(int32(evt.X), int32(evt.Y)))
			}
		}
	})
}
