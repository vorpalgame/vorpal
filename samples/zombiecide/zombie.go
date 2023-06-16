package zombiecide

import (
	"strconv"

	"github.com/vorpalgame/vorpal/bus"
)

type Zombie interface {
	RenderImage(evt bus.MouseEvent) *bus.ImageLayer
}
type zombie struct {
	currentLocation Point
}

func NewZombie() Zombie {
	return &zombie{&point{600, 600}}
}

var walkFrame = 1
var walkRepeatFrame = 0

// Should we be passing * to evt?
func (z *zombie) RenderImage(evt bus.MouseEvent) *bus.ImageLayer {
	if evt.LeftButton().IsDown() {
		return z.renderAttack()
	} else {
		point := z.calculateMove(evt)
		if point.x == 0 && point.y == 0 {
			return z.renderIdle()
		} else {
			z.currentLocation.Add(point) //In future point dimension/direction/size may determine behavior.
			return z.renderWalk()
		}
	}
}
func (z *zombie) calculateMove(evt bus.MouseEvent) *point {
	x := int32(-4)
	y := int32(-4)
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

func (z *zombie) renderWalk() *bus.ImageLayer {

	walkRepeatFrame++
	if walkRepeatFrame > 3 {
		walkFrame++
		walkRepeatFrame = 0
	}
	if walkFrame > 10 {
		walkFrame = 1
	}

	return z.renderImage(200, 300, walkFrame, "walk")
}

var attackFrame = 1
var attackRepeatFrame = 0

func (z *zombie) renderAttack() *bus.ImageLayer {

	attackRepeatFrame++
	if attackRepeatFrame > 3 {
		attackFrame++
		attackRepeatFrame = 0
	}
	if attackFrame > 8 {
		attackFrame = 1
	}

	return z.renderImage(200, 300, attackFrame, "attack")
}

var idleFrame = 1
var idleRepeatFrame = 0

func (z *zombie) renderIdle() *bus.ImageLayer {

	idleRepeatFrame++
	if idleRepeatFrame > 2 {
		idleFrame++
		idleRepeatFrame = 0
	}
	if idleFrame > 15 {
		idleFrame = 1
	}
	return z.renderImage(200, 300, idleFrame, "Idle")
}

// Cache image layers for zombie???
func (z *zombie) renderImage(width, height int32, frame int, name string) *bus.ImageLayer {
	imgLayer := bus.NewImageLayer("samples/resources/zombiecide/"+name+" ("+strconv.Itoa(frame)+").png", z.currentLocation.GetX(), z.currentLocation.GetY(), width, height)
	return &imgLayer
}
