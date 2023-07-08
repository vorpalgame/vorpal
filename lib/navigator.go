package lib

import "log"

//TODO Unit tests...

//"log"

func NewLocation() Navigator {
	return &NavigatorData{}
}

func NewNavigator(point PointData, xMove, yMove, maxXOffset, maxYOffset int32, controller ActionStageController) Navigator {
	return &NavigatorData{&point, xMove, yMove, maxXOffset, maxYOffset, controller}
}

type Navigator interface {
	Point
	MoveTowardMouse(cursorPoint Point)
	MoveByIncrement(toPoint Point)
	CalculateMoveIncrement(pointerLocation Point) Point
	ActionStageController
}

type NavigatorData struct {
	Location              *PointData `yaml:"CurrentLocation"`
	XMove                 int32      `yaml:"XMove"`
	YMove                 int32      `yaml:"YMove"`
	MaxXOffset            int32      `yaml:"MaxXOffset"`
	MaxYOffset            int32      `yaml:"MaxYOffset"`
	ActionStageController `yaml:"-"`
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

// Clone
func (cl *NavigatorData) GetCurrentPoint() Point {
	return &PointData{cl.Location.GetX(), cl.Location.GetY()}
}

// OK hack the color check for now...
func (cl *NavigatorData) MoveByIncrement(toPoint Point) {

	if cl.isLegal(toPoint) {
		cl.Location.Add(toPoint)
	}

}

// This inelegant hack is here for prototyping...
// Golang doesn't use 0...255 for RGB

// The move is legal if it doesn't wander into illegal areas or if no behavior map is set.
func (cl *NavigatorData) isLegal(toPoint Point) bool {
	absolute := cl.GetCurrentPoint()
	absolute.Add(toPoint)
	var isLegal bool = true
	if cl.ActionStageController != nil {
		color := cl.ActionStageController.CheckBehaviorColorAt(absolute)

		r, g, b, _ := color.RGBA()
		//We are in transparent...legal...
		isLegal = 0 == r && 0 == g && 0 == b
	}
	log.Default().Println(isLegal)
	return isLegal
}

func (cl *NavigatorData) MoveTowardMouse(cursorPoint Point) {
	cl.MoveByIncrement(cl.CalculateMoveIncrement(cursorPoint))

}
func (cl *NavigatorData) CalculateMoveIncrement(p Point) Point {

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
