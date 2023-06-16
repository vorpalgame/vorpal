package flipbook

import (
	"log"
	"strconv"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

type flipBook struct {
	bus       bus.VorpalBus
	drawEvent bus.DrawEvent
	textEvent bus.TextEvent
}

// TODO The cards should have locations or the game/board should dictate those...
// TODO The fonts and text layers need to be added to the text event.
var animation = flipBook{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"
var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

func Init() {
	log.Println("New card game")
	animation.bus = bus.GetVorpalBus()
	//animation.bus.AddEngineListener(&animation)

	var x = int32(0)
	for {
		x = animation.drawWalk(x)
		animation.drawAttack(x)
		animation.drawAttack(x)
		x = animation.drawWalk(x)
		animation.drawDead(x)

	}

}

func (g *flipBook) drawWalk(x int32) int32 {

	for i := 1; i < 10; i++ {
		g.drawImage(x, 200, 300, i, "walk")
		x = x + 50

		if x > 1900 {
			x = 0
		}
	}
	return x
}

func (g *flipBook) drawAttack(x int32) {

	for i := 1; i < 9; i++ {
		g.drawImage(x, 200, 300, i, "attack")
	}
}

func (g *flipBook) drawDead(x int32) {

	for i := 1; i < 12; i++ {
		g.drawImage(x, 300, 300, i, "dead")
	}
}

func (g *flipBook) drawImage(x, width, height int32, frame int, name string) {
	animation.drawEvent = bus.NewDrawEvent()
	animation.drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/flipbook/background.png", 0, 0, 1920, 1080))
	animation.drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/flipbook/"+name+" ("+strconv.Itoa(frame)+").png", x, 600, width, height))
	animation.bus.SendDrawEvent(animation.drawEvent)
	time.Sleep(150 * time.Millisecond)
}
func (g *flipBook) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			//TODO
		} else if evt.GetKey().EqualsIgnoreCase("N") {
			//TODO
		}

	}

}
