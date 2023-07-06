package lib

//TODO Unit tests...

//"log"

func NewLocation() Navigator {
	return &NavigatorData{}
}

func NewNavigator(point PointData, xMove, yMove, maxXOffset, maxYOffset int32) Navigator {
	return &NavigatorData{&point, xMove, yMove, maxXOffset, maxYOffset}
}

type Navigator interface {
	Point
	MoveTowardMouse(cursorPoint Point)
	Move(toPoint Point)
	CalculateMove(pointerLocation Point) Point
}

type NavigatorData struct {
	Location   *PointData `yaml:"CurrentLocation"`
	XMove      int32      `yaml:"XMove"`
	YMove      int32      `yaml:"YMove"`
	MaxXOffset int32      `yaml:"MaxXOffset"`
	MaxYOffset int32      `yaml:"MaxYOffset"`
}

func (cl *NavigatorData) Add(point Point) {
	cl.Location.Add(point)
}

func (cl *NavigatorData) GetX() int32 {
	return cl.Location.GetX()
}
func (cl *NavigatorData) GetY() int32 {
	return cl.Location.GetY()
}

// TODO Should probably clone...
func (cl *NavigatorData) GetCurrentPoint() Point {
	return cl.Location
}

func (cl *NavigatorData) Move(toPoint Point) {
	cl.Location.Add(toPoint)
}

func (cl *NavigatorData) MoveTowardMouse(cursorPoint Point) {
	cl.Move(cl.CalculateMove(cursorPoint))

}
func (cl *NavigatorData) CalculateMove(p Point) Point {

	var point = PointData{cl.XMove, cl.YMove}

	if p.GetX() < cl.GetX() {
		point.X = point.X * -1
	}

	if p.GetY() < cl.GetY() {
		point.Y = point.Y * -1
	}

	var xOffset = p.GetX() - cl.GetX()
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < cl.MaxXOffset {
		point.X = 0
	}
	yOffset := p.GetY() - cl.GetY()
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < cl.MaxYOffset {
		point.Y = 0
	}

	return &point

}
