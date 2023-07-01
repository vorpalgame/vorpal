package lib

//TODO Unit tests...

//"log"

func NewLocation() Navigator {
	return &navigatorData{}
}

func NewNavigatorOffset(point Point, xMove, yMove, maxXOffset, maxYOffset int32) Navigator {
	return &navigatorData{point, xMove, yMove, maxXOffset, maxYOffset}
}

type Navigator interface {
	GetCurrentPoint() Point
	Move(toPoint Point)
	CalculateMove(pointerLocation Point) Point
	GetX() int32
	GetY() int32
}
type navigatorData struct {
	currentLocation                      Point
	xMove, yMove, maxXOffset, maxYOffset int32
}

func (cl *navigatorData) GetX() int32 {
	return cl.currentLocation.GetX()
}
func (cl *navigatorData) GetY() int32 {
	return cl.currentLocation.GetY()
}
func (cl *navigatorData) GetCurrentPoint() Point {
	return cl.currentLocation
}

func (cl *navigatorData) Move(toPoint Point) {
	cl.currentLocation.Add(toPoint)
	//log.Default().Println(cl.GetCurrentPoint())
}

func (cl *navigatorData) CalculateMove(p Point) Point {

	point := point{cl.xMove, cl.yMove}

	//abs math function is floating point so just -1 multiple
	if p.GetX() > cl.GetX() {
		point.x = point.x * -1
	}

	if p.GetY() > cl.GetY() {
		point.y = point.y * -1
	}

	var xOffset = p.GetX() - cl.GetX()
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < cl.maxXOffset {
		point.x = 0
	}
	yOffset := p.GetY() - cl.GetY()
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < cl.maxYOffset {
		point.y = 0
	}
	return &point

}
