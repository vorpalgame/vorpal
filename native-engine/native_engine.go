package native_engine

import (
	"fmt"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type engine struct {
	bus.VorpalBus
	//bus.StandardMediaPeerController
	MediaCache
	currentRenderImage, currentDisplayImage image.Image
	window                                  screen.Window
	screen                                  screen.Screen
	drawEventChannel                        chan bus.DrawEvent
	textEventChannel                        chan bus.TextEvent
}

func (e *engine) OnTextEvent(inputChannel <-chan bus.TextEvent) {
	for evt := range inputChannel {
		e.textEventChannel <- evt
	}
}

func (e *engine) OnDrawEvent(inputChannel <-chan bus.DrawEvent) {
	for evt := range inputChannel {
		e.drawEventChannel <- evt
	}
}

func NewEngine() lib.Engine {

	log.Println("New native engine...")
	e := engine{}
	e.MediaCache = NewMediaCache()
	e.drawEventChannel = make(chan bus.DrawEvent, 10)
	e.textEventChannel = make(chan bus.TextEvent, 10)
	e.VorpalBus = bus.GetVorpalBus()
	e.AddDrawEventListener(&e)
	e.AddTextEventListener(&e)
	go initWindow(&e)
	return &e
}

func (e *engine) Start() {
	log.Println("Start Native Engine...")
	go textPipeline(e, e.textEventChannel)
	go renderPipeline(e, e.drawEventChannel)
	for {
		//Loop until exit is called...
	}
}

var initWindow = func(e *engine) {
	driver.Main(func(s screen.Screen) {
		e.screen = s
		w, err := s.NewWindow(nil)
		e.window = w
		if err != nil {
			panic(err)
			return
		}
		defer w.Release()
		mouseChannel := make(chan mouse.Event, 1)
		go processMouse(e, mouseChannel)
		lifecycleChannel := make(chan lifecycle.Event, 1)
		go processLifeCycle(e, lifecycleChannel)
		keyChannel := make(chan key.Event, 1)
		go processKey(e, keyChannel)
		for {
			switch evt := w.NextEvent().(type) {
			case lifecycle.Event:
				lifecycleChannel <- evt
			case mouse.Event:
				mouseChannel <- evt
			case key.Event:
				keyChannel <- evt
			}
		}
	})
}

var textPipeline = func(e *engine, inputChannel <-chan bus.TextEvent) {
	for evt := range inputChannel {
		_ = evt
		//log.Println("Text Event...")
		//log.Println(evt)
	}
}

var renderPipeline = func(e *engine, inputChannel <-chan bus.DrawEvent) {
	cacheChan := make(chan bus.DrawLayersEvent, 1)
	go cacheImages(e, cacheChan)
	for evt := range inputChannel {
		switch evt := evt.(type) {
		case bus.DrawLayersEvent:
			cacheChan <- evt
		}

	}
}

var cacheImages = func(e *engine, inputChannel <-chan bus.DrawLayersEvent) {
	renderChan := make(chan bus.DrawLayersEvent, 1)
	go renderImageLayers(e, renderChan)
	for evt := range inputChannel {
		e.CacheImages(evt)
		renderChan <- evt
	}

}

var renderImageLayers = func(e *engine, inputChannel <-chan bus.DrawLayersEvent) {
	blitChannel := make(chan *image.RGBA, 1)
	go blitImage(e, blitChannel)
	for evt := range inputChannel {
		var buffer *image.RGBA
		for _, layer := range evt.GetImageLayers() {
			for _, imgData := range layer.LayerMetadata {
				img := *e.GetImage(imgData.ImageFileName)
				if buffer == nil {
					buffer = image.NewRGBA(img.Bounds())
				}
				if imgData.HorizontalFlip {
					img = flip(img)
				}
				draw.Draw(buffer, getRect(imgData), img, *getPoint(0, 0), draw.Over)
			}

		}
		blitChannel <- buffer
	}
}

func flip(img image.Image) image.Image {
	// create and return a grid of pixels
	size := img.Bounds().Size()
	var grid [][]color.Color
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
		}
		grid = append(grid, y)
	}
	flipHorizontal(grid)
	xlen, ylen := len(grid), len(grid[0])
	rect := image.Rect(0, 0, xlen, ylen)
	flipImg := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			flipImg.Set(x, y, grid[x][y])
		}
	}
	return image.Image(flipImg)
}
func flipHorizontal(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ {
			z := len(col) - y - 1
			col[y], col[z] = col[z], col[y]
		}
	}
}
func flipVertical(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ {
			z := len(col) - y - 1
			col[y], col[z] = col[z], col[y]
		}
	}
}

var blitImage = func(e *engine, channel <-chan *image.RGBA) {

	for buffer := range channel {
		b, _ := e.screen.NewBuffer(buffer.Bounds().Max)
		draw.Draw(b.RGBA(), b.Bounds(), buffer, *getPoint(0, 0), draw.Over)
		e.window.Upload(image.Point{0, 0}, b, buffer.Bounds())
		e.window.Publish()
		b.Release()
		buffer = nil
	}
}

// TODO Rewrite the key even to use the key.Event internally.
var processKey = func(e *engine, channel <-chan key.Event) {
	for event := range channel {
		log.Println(event)
		e.SendKeyEvent(bus.NewKeyEvent(event))
	}

}

var processMouse = func(e *engine, channel <-chan mouse.Event) {
	for event := range channel {
		e.SendMouseEvent(bus.NewMouseEvent(event))
	}

}

var processLifeCycle = func(e *engine, channel <-chan lifecycle.Event) {
	for event := range channel {
		log.Println("processLifecycle")
		log.Println(event)

		if event.To == lifecycle.StageDead {
		} else if event.To == lifecycle.StageVisible {
		}
	}

}

// ////////////////////////////////////////////////////////////////////////////
// TODO We need these on ImageMetadta interfaces...
func getRect(imgData *lib.ImageMetadata) image.Rectangle {
	return image.Rectangle{*getPoint(imgData.X, imgData.Y), *getPoint(imgData.X+imgData.Width, imgData.Y+imgData.Height)}
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
