package zombiecide

import (
	"github.com/vorpalgame/vorpal/bus"
)

// Duplicate functions for walk, dead, idle, attack can be refactored and paramaterized...
type Zombie interface {
	RunZombie(drawEvent bus.DrawEvent, mouseEvent bus.MouseEvent)
}
type zombie struct {
	currentLocation          Point
	walk, dead, idle, attack SpriteController
	numberIdle               int
}

func NewZombie() Zombie {
	//"samples/resources/zombiecide/"+s.fileBaseName+" ("+strconv.Itoa(s.currentFrame)+").png"

	z := zombie{&point{600, 600}, newSprite(10, 3, 200, 300, "walk"), newSprite(12, 3, 300, 300, "dead"), newSprite(15, 3, 200, 300, "idle"), newSprite(8, 3, 200, 300, "attack"), 0}
	z.dead.SetToLoop(false)

	return &z
}

func newSprite(x, y, width, height int, name string) SpriteController {
	sprite := NewSpriteController(10, 3, 200, 300)
	sprite.SetImageTemplate(getImageTemplate(name))
	sprite.SetAudio(getAudioTemplate(name))
	return sprite
}

func getImageTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + " (%d).png"
}

func getAudioTemplate(name string) string {
	return "samples/resources/zombiecide/" + name + ".mp3"
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
			if z.numberIdle < 250 {
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

	sprite.RunSprite(drawEvent, z.currentLocation, flipHorizontal)
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
