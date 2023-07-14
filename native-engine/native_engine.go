package native_engine

import (
	"fmt"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

//TODO We should revisit the caching mechanisms and flushing and perhaps caching reversed/flipped images.

type engine struct {
	bus.VorpalBus
	//bus.StandardMediaPeerController
	MediaCache
	currentRenderImage, currentDisplayImage image.Image
	window                                  screen.Window
	screen                                  screen.Screen
	drawEventChannel                        chan bus.DrawEvent
	textEventChannel                        chan bus.TextEvent
	controlEventChannel                     chan bus.ControlEvent
	audioEventChannel                       chan bus.AudioEvent
}

///////////////////////////////////////////////////////////////////////////////////////
//// Listener methods to take messages from bus channel and them to local channel
//// This permits local implementation and processing on channel to be blocking, non-blocking
//// and split on multiple pipeline channels whichever makes sense for the current engine
//// implementation. The bus conveys the message but the internal channels are set up
//// to reflect the processing and concurrency requirements.
/////////////////////////////////////////////////////////////////////////////////////////

func (e *engine) OnAudioEvent(inputChannel <-chan bus.AudioEvent) {
	for evt := range inputChannel {
		e.audioEventChannel <- evt
	}
}

func (e *engine) OnControlEvent(inputChannel <-chan bus.ControlEvent) {
	for evt := range inputChannel {
		e.controlEventChannel <- evt
	}
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
	//Adjust channel sizes if needed....
	e.drawEventChannel = make(chan bus.DrawEvent, 1)
	e.textEventChannel = make(chan bus.TextEvent, 1)
	e.controlEventChannel = make(chan bus.ControlEvent, 1)
	e.audioEventChannel = make(chan bus.AudioEvent, 1)

	e.VorpalBus = bus.GetVorpalBus()
	e.AddDrawEventListener(&e)
	e.AddTextEventListener(&e)
	e.AddControlEventListener(&e)
	e.AddAudioEventListener(&e)
	go initWindow(&e)
	return &e
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

func (e *engine) Start() {
	log.Println("Start Native Engine...")
	go textPipeline(e, e.textEventChannel)
	go audioPipeline(e, e.audioEventChannel)
	go controlPipeline(e, e.controlEventChannel)
	go renderPipeline(e, e.drawEventChannel)
	for { //Loop until exit is called...
	}
}

var audioPipeline = func(e *engine, inputChannel chan bus.AudioEvent) {
	for event := range inputChannel {
		_ = event
	}
}

var controlPipeline = func(e *engine, inputChannel chan bus.ControlEvent) {
	for event := range inputChannel {
		_ = event
	}
}

var textPipeline = func(e *engine, inputChannel <-chan bus.TextEvent) {
	for event := range inputChannel {
		switch evt := event.(type) {
		case bus.MultilineTextEvent:
			_ = evt
		}
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
		//drawText(buffer)
		blitChannel <- buffer
	}
}

func pt(p fixed.Point26_6) image.Point {
	return image.Point{
		X: int(p.X+32) >> 6,
		Y: int(p.Y+32) >> 6,
	}
}

func drawText(dst *image.RGBA) {
	//draw.Draw(dst, dst.Bounds(), image.Black, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  dst,
		Src:  image.White,
		Face: basicfont.Face7x13,
	}
	ss := []string{
		"The quick brown fox jumps over the lazy dog.",
		"Hello, 世界.",
		"U+FFFD is \ufffd.",
	}
	for i, s := range ss {
		d.Dot = fixed.P(20, 100*i+80)
		dot0 := pt(d.Dot)
		d.DrawString(s)
		dot1 := pt(d.Dot)
		dst.SetRGBA(dot0.X, dot0.Y, color.RGBA{0xff, 0x00, 0x00, 0xff})
		dst.SetRGBA(dot1.X, dot1.Y, color.RGBA{0x00, 0x00, 0xff, 0xff})
	}
}

func flip(img image.Image) image.Image {
	size := img.Bounds().Size()
	flipImg := image.NewNRGBA(img.Bounds())
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			xp := size.X - x - 1
			flipImg.Set(x, y, img.At(xp, y))
		}
	}
	return flipImg
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
