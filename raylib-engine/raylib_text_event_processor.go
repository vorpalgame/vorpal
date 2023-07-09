package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// /////////////////////////////////////////////////////////////////
// /// Text Event Processor
// /////////////////////////////////////////////////////////////////

func NewTextEventProcessor(mediaCache MediaCache) bus.TextEventProcessor {
	return &textData{mediaCache}
}

type textData struct {
	MediaCache
}

// TODO Text event processing is stil undifferentiated and brittle. For example,
// a missing font will crash. As we get new text event types they may not use fonts. Etc.
// We are rewriting the TextEvent interface and types to be able to switch and differentiate
// behavior.
func (tep *textData) ProcessTextEvent(evt bus.TextEvent) {

	//Currently the switch isn't differentiating on different
	//subtypes of TextEvent but it will.
	if evt != nil {
		switch evt := evt.(type) {
		case bus.MultilineTextEvent:
			tep.renderText(evt)
		}
	}
}
func (tep *textData) renderText(evt bus.MultilineTextEvent) {
	//Make sure we don't get a race on nil check.
	renderImg := tep.GetCurrentRenderImage()

	if renderImg != nil {
		tep.CacheFonts(evt)
		//TODO The lines will not be wrapped here so this is temporary
		//The next step is to send each presplit line from the other side of the bus
		//and then iterate over it here.
		//
		intx, inty := evt.GetLocation()
		var x = float32(intx)
		var y = float32(inty)
		for _, txt := range evt.GetText() {
			tep.CacheFonts(txt)
			rl.ImageDrawTextEx(renderImg, rl.Vector2{x, y}, *tep.GetFont(txt.GetFont()), txt.GetText(), float32(txt.GetFontSize()), 0, rl.Black)
			//How to do line spacing????
			y += float32(txt.GetFontSize()) * float32(1.1) //Extra space..
		}

	}

}
