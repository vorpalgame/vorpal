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
	"image/png"
	"log"
	"os"
)

type engine struct {
	bus.VorpalBus
	bus.StandardMediaPeerController
	MediaCache
	currentRenderImage, currentDisplayImage image.Image
	window                                  screen.Window
	screen                                  screen.Screen
}

func NewEngine() lib.Engine {

	log.Println("New native engine...")
	e := engine{}
	e.MediaCache = NewMediaCache()
	e.StandardMediaPeerController = bus.NewGameController()
	e.VorpalBus = bus.GetVorpalBus()
	go initWindow(&e)
	return &e
}

func (e *engine) Start() {
	log.Println("Start Native Engine...")
	renderPipelineChan := make(chan bus.DrawEvent, 1)
	go renderPipeline(e, renderPipelineChan)
	for {
		renderPipelineChan <- e.GetDrawEvent()
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

var renderPipeline = func(e *engine, drawChannel <-chan bus.DrawEvent) {
	cacheChan := make(chan bus.DrawLayersEvent, 1)
	go cacheImages(e, cacheChan)
	for evt := range drawChannel {
		switch evt := evt.(type) {
		case bus.DrawLayersEvent:
			cacheChan <- evt
		}

	}
}

var cacheImages = func(e *engine, drawChannel <-chan bus.DrawLayersEvent) {
	renderChan := make(chan bus.DrawLayersEvent, 1)
	go renderImageLayers(e, renderChan)
	for evt := range drawChannel {
		e.CacheImages(evt)
		renderChan <- evt
	}

}

var renderImageLayers = func(e *engine, drawChannel <-chan bus.DrawLayersEvent) {
	blitChannel := make(chan *image.RGBA, 1)
	go blitImage(e, blitChannel)
	for evt := range drawChannel {
		var buffer *image.RGBA
		for _, layer := range evt.GetImageLayers() {
			for _, imgData := range layer.LayerMetadata {
				img := *e.GetImage(imgData.ImageFileName)
				if buffer == nil {
					buffer = image.NewRGBA(img.Bounds())
				}
				draw.Draw(buffer, getRect(imgData), img, *getPoint(0, 0), draw.Over)
			}

		}
		blitChannel <- buffer
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
