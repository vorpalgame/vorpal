package lib

//TODO Unit tests...
import (
	//"log"

	"github.com/vorpalgame/vorpal/bus"
)

func NewCurrentLocation() CurrentLocation {
	return &currentLocationData{&point{600, 600}, -4, -2, 5, 5} //TODO Get rid of bad defaults...Convenient for now...
}

type CurrentLocation interface {
	GetCurrentPoint() Point
	Move(toPoint Point)
	CalculateMove(evt bus.MouseEvent) Point
	GetX() int32
	GetY() int32
}
type currentLocationData struct {
	currentLocation                      Point
	xMove, yMove, maxXOffset, maxYOffset int32
}

func (cl *currentLocationData) GetX() int32 {
	return cl.currentLocation.GetX()
}
func (cl *currentLocationData) GetY() int32 {
	return cl.currentLocation.GetY()
}
func (cl *currentLocationData) GetCurrentPoint() Point {
	return cl.currentLocation
}

func (cl *currentLocationData) Move(toPoint Point) {
	cl.currentLocation.Add(toPoint)
	//log.Default().Println(cl.GetCurrentPoint())
}

func (cl *currentLocationData) CalculateMove(evt bus.MouseEvent) Point {

	point := point{cl.xMove, cl.yMove}

	//abs math function is floating point so just -1 multiple
	if evt.GetX() > cl.GetX() {
		point.x = point.x * -1
	}

	if evt.GetY() > cl.GetY() {
		point.y = point.y * -1
	}

	var xOffset = evt.GetX() - cl.GetX()
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < cl.maxXOffset {
		point.x = 0
	}
	yOffset := evt.GetY() - cl.GetY()
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < cl.maxYOffset {
		point.y = 0
	}
	return &point

}
