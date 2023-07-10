package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// /////////////////////////////////////////////////////////////////
// /// Text Event Processor
// /////////////////////////////////////////////////////////////////
var raylibProcessTextEvent = func(evt bus.TextEvent, cache MediaCache) {

	//Currently the switch isn't differentiating on different
	//subtypes of TextEvent but it will.
	if evt != nil {
		switch evt := evt.(type) {
		case bus.MultilineTextEvent:
			renderText(evt, cache)
		}
	}
}

func renderText(evt bus.MultilineTextEvent, cache MediaCache) {
	//Make sure we don't get a race on nil check.
	renderImg := cache.GetCurrentRenderImage()

	if renderImg != nil {
		cache.CacheFonts(evt)
		//TODO The lines will not be wrapped here so this is temporary
		//The next step is to send each presplit line from the other side of the bus
		//and then iterate over it here.
		//
		intx, inty := evt.GetLocation()
		var x = float32(intx)
		var y = float32(inty)
		for _, txt := range evt.GetText() {
			cache.CacheFonts(txt)
			rl.ImageDrawTextEx(renderImg, rl.Vector2{x, y}, *cache.GetFont(txt.GetFont()), txt.GetText(), float32(txt.GetFontSize()), 0, rl.Black)
			//How to do line spacing????
			y += float32(txt.GetFontSize()) * float32(1.1) //Extra space..
		}

	}

}
