package zombiecide

import (
	"log"

	"github.com/vorpalgame/vorpal/bus"
)

// Duplicate functions for walk, dead, idle, attack can be refactored and paramaterized...
type Zombie interface {
	RunZombie(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent)
}
type zombie struct {
	currentLocation                         Point
	walk, dead, idle, attack, currentSprite SpriteController

	numberIdle int
}

func NewZombie() Zombie {
	//"samples/resources/zombiecide/"+s.fileBaseName+" ("+strconv.Itoa(s.currentFrame)+").png"

	z := zombie{&point{600, 600}, NewSpriteController(10, 3, 200, 300, getFileTemplate("walk")), NewSpriteController(12, 3, 300, 300, getFileTemplate("dead")), NewSpriteController(15, 3, 200, 300, getFileTemplate("idle")), NewSpriteController(9, 3, 200, 300, getFileTemplate("attack")), nil, 0}
	z.dead.SetToLoop(false)
	z.walk.SetAudio("samples/resources/zombiecide/moan.mp3")
	z.attack.SetAudio("samples/resources/zombiecide/roar.mp3")
	return &z
}

func getFileTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

// Note we aren't really "rendering" anything. We are specifying the name of the source file, x,y, width and height coordianates.
// It is metadata for the actual rendering by the engine.
func (z *zombie) RunZombie(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent) {
	var sprite SpriteController
	point := z.calculateMove(mouseEvent)
	flipHorizontal := mouseEvent.GetX() < z.currentLocation.GetX()
	if mouseEvent.LeftButton().IsDown() {
		sprite = z.attack
	} else {
		if point.x == 0 && point.y == 0 {
			z.numberIdle++
			if z.numberIdle < 100 {
				sprite = z.idle
			} else {
				sprite = z.dead
			}
		} else {
			z.numberIdle = 0
			z.currentLocation.Add(point) //In future point dimension/direction/size may determine behavior.
			sprite = z.walk
		}
	}
	log.Default().Println(sprite)
	if sprite != nil {
		if z.currentSprite != nil && z.currentSprite != sprite {
			z.currentSprite.StopSprite()
		}
		z.currentSprite = sprite
		sprite.RunSprite(drawEvent, z.currentLocation, flipHorizontal)
	}
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
