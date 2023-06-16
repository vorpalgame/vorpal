package zombiecide

import (
	"log"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

type zombiecide struct {
	bus        bus.VorpalBus
	textEvent  bus.TextEvent
	mouseEvent bus.MouseEvent
	background bus.ImageLayer
	zombie     Zombie
}

type Point interface {
	GetX() int32
	GetY() int32
	Add(Point)
}

type point struct {
	x, y int32
}

func (p *point) GetX() int32 {
	return p.x
}

func (p *point) GetY() int32 {
	return p.y
}

func (p *point) Add(addPoint Point) {
	p.x += addPoint.GetX()
	p.y += addPoint.GetY()
}

// TODO The cards should have locations or the game/board should dictate those...
// TODO The fonts and text layers need to be added to the text event.
var zombies = zombiecide{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"
var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

func Init() {
	log.Println("New card game")
	zombies.zombie = NewZombie()
	zombies.bus = bus.GetVorpalBus()
	zombies.bus.AddEngineListener(&zombies)

	zombies.background = bus.NewImageLayer("samples/resources/zombiecide/background.png", 0, 0, 1920, 1080)

	for {
		if zombies.mouseEvent != nil {
			zombies.drawImage(zombies.zombie.RenderImage(zombies.mouseEvent))
			time.Sleep(20 * time.Millisecond)
		}

	}

}

// TODO Current location should be center or front of center.

func (z *zombiecide) drawImage(img *bus.ImageLayer) {
	drawEvent := bus.NewDrawEvent()
	drawEvent.AddImageLayer(z.background)
	drawEvent.AddImageLayer(*img)
	z.bus.SendDrawEvent(drawEvent)

}

func (z *zombiecide) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			//TODO
		} else if evt.GetKey().EqualsIgnoreCase("N") {
			//TODO
		}

	}

}
func (z *zombiecide) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		z.mouseEvent = evt

	}
}
