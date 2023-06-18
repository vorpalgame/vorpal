package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

// Duplicate functions for walk, dead, idle, attack can be refactored and paramaterized...
type Zombie interface {
	RenderImage(evt bus.MouseEvent) *bus.ImageLayer
}
type zombie struct {
	currentLocation          Point
	walk, dead, idle, attack Sprite
	numberIdle               int
}

func NewZombie() Zombie {

	z := zombie{&point{600, 600}, NewSprite(10, 3, 200, 300, "walk"), NewSprite(12, 3, 300, 300, "dead"), NewSprite(15, 3, 200, 300, "idle"), NewSprite(9, 3, 200, 300, "attack"), 0}
	z.dead.SetToLoop(false)
	return &z
}

// Note we aren't really "rendering" anything. We are specifying the name of the source file, x,y, width and height coordianates.
// It is metadata for the actual rendering by the engine.
func (z *zombie) RenderImage(evt bus.MouseEvent) *bus.ImageLayer {
	var imageLayer *bus.ImageLayer
	point := z.calculateMove(evt)
	flipHorizontal := evt.GetX() < z.currentLocation.GetX()
	if evt.LeftButton().IsDown() {
		imageLayer = z.attack.RenderNext(z.currentLocation, flipHorizontal)
	} else {
		if point.x == 0 && point.y == 0 {
			z.numberIdle++
			if z.numberIdle < 100 {
				imageLayer = z.idle.RenderNext(z.currentLocation, flipHorizontal)
			} else {
				imageLayer = z.dead.RenderNext(z.currentLocation, flipHorizontal)
			}
		} else {
			z.numberIdle = 0

			z.currentLocation.Add(point) //In future point dimension/direction/size may determine behavior.
			imageLayer = z.walk.RenderNext(z.currentLocation, flipHorizontal)
		}
	}
	return imageLayer
}

// TODO The calcs are using the upper left for location relative to image and that probably isn't desired.
func (z *zombie) calculateMove(evt bus.MouseEvent) *point {
	x := int32(-4)
	y := int32(-2)
	point := point{x, y}
	//abs math function is floating point so just -1 multiple
	if evt.GetX() > z.currentLocation.GetX() {
		point.x = point.x * -1
	}
	if evt.GetY() > z.currentLocation.GetY() {
		point.y = point.y * -1
	}

	var xOffset = evt.GetX() - z.currentLocation.GetX()
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < 5 {
		point.x = 0
	}
	yOffset := evt.GetY() - z.currentLocation.GetY()
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < 5 {
		point.y = 0
	}
	return &point
}
