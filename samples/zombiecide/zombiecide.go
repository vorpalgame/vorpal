package zombiecide

import (
	"log"
	"strconv"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

type zombie struct {
	bus             bus.VorpalBus
	textEvent       bus.TextEvent
	mouseEvent      bus.MouseEvent
	currentLocation point
	background      bus.ImageLayer
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
var animation = zombie{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"
var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

func Init() {
	log.Println("New card game")
	animation.bus = bus.GetVorpalBus()
	animation.bus.AddEngineListener(&animation)

	animation.background = bus.NewImageLayer("samples/resources/zombiecide/background.png", 0, 0, 1920, 1080)

	animation.currentLocation = point{0, 600}
	for {
		if animation.mouseEvent != nil {
			if animation.mouseEvent.LeftButton().IsDown() {
				animation.drawAttack()
			} else {
				point := animation.calculateMove()
				animation.currentLocation.Add(point) //In future point dimension/direction/size may determine behavior.
				animation.drawWalk()

			}
		}
		//
		// animation.drawAttack(currentX)
		// currentX = animation.drawWalk(currentX)
		// animation.drawDead(currentX)

	}

}

func (z *zombie) calculateMove() *point {

	point := point{-25, -25}
	//abs math function is floating point so just -1 multiple
	if z.mouseEvent.GetX() > z.currentLocation.x {
		point.x = point.x * -1
	}
	if z.mouseEvent.GetY() > z.currentLocation.y {
		point.y = point.y * -1
	}
	return &point
}
func (z *zombie) drawWalk() {

	for i := 1; i < 10; i++ {
		z.drawImage(200, 300, i, "walk")

	}

}

func (z *zombie) drawAttack() {

	for i := 1; i < 9; i++ {
		z.drawImage(200, 300, i, "attack")
	}
}

// func (z *zombie) drawDead(x int32) {

// 	for i := 1; i < 12; i++ {
// 		z.drawImage(x, 300, 300, i, "dead")
// 	}
// }

func (z *zombie) drawImage(width, height int32, frame int, name string) {
	drawEvent := bus.NewDrawEvent()
	drawEvent.AddImageLayer(z.background)
	drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/zombiecide/"+name+" ("+strconv.Itoa(frame)+").png", z.currentLocation.x, z.currentLocation.y, width, height))
	z.bus.SendDrawEvent(drawEvent)
	time.Sleep(25 * time.Millisecond)
}
func (z *zombie) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			//TODO
		} else if evt.GetKey().EqualsIgnoreCase("N") {
			//TODO
		}

	}

}
func (z *zombie) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		z.mouseEvent = evt

	}
}
