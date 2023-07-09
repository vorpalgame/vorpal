package lib

import "log"

//TODO Unit tests...

//"log"

func NewLocation() Navigator {
	return &NavigatorData{}
}

func NewNavigator(x, y, xMove, yMove, maxXOffset, maxYOffset int32, controller ActionStageController) Navigator {
	return &NavigatorData{x, y, xMove, yMove, maxXOffset, maxYOffset, controller}
}

type Navigator interface {
	MoveTowardMouse(x, y int32)
	MoveByIncrement(x, y int32)
	CalculateMoveIncrement(moveToX, moveToY int32) (x, y int32)
	GetCurrentPoint() (x, y int32)
	ActionStageController
}

type NavigatorData struct {
	X                     int32 `yaml:"CurrentX"`
	Y                     int32 `yaml:"CurrentY"`
	XMove                 int32 `yaml:"XMove"`
	YMove                 int32 `yaml:"YMove"`
	MaxXOffset            int32 `yaml:"MaxXOffset"`
	MaxYOffset            int32 `yaml:"MaxYOffset"`
	ActionStageController `yaml:"-"`
}

func (cl *NavigatorData) Add(x, y int32) {
	cl.X += x
	cl.Y += y
}

// Clone
func (cl *NavigatorData) GetCurrentPoint() (x, y int32) {
	return cl.X, cl.Y
}

// OK hack the color check for now...
func (cl *NavigatorData) MoveByIncrement(x, y int32) {

	if cl.isLegal(x, y) {
		cl.Add(x, y)
	}

}

// This inelegant hack is here for prototyping...
// Golang doesn't use 0...255 for RGB

// The move is legal if it doesn't wander into illegal areas or if no behavior map is set.
func (cl *NavigatorData) isLegal(x, y int32) bool {
	absoluteX, absoluteY := cl.GetCurrentPoint()
	absoluteX += x
	absoluteY += y
	var isLegal bool = true
	if cl.ActionStageController != nil {
		color := cl.ActionStageController.CheckBehaviorColorAt(absoluteX, absoluteY)

		r, g, b, _ := color.RGBA()
		//TODO Should probably use black as legal color so we just check 0
		isLegal = 65535 == r && 65535 == g && 65535 == b
	}
	log.Default().Println(isLegal)
	return isLegal
}

func (cl *NavigatorData) MoveTowardMouse(x, y int32) {
	cl.MoveByIncrement(cl.CalculateMoveIncrement(x, y))

}

// This is in the middle of refactoring.
func (cl *NavigatorData) CalculateMoveIncrement(moveToX, moveToY int32) (x, y int32) {

	incrementX, incrementY := cl.increment(moveToX, moveToY)
	//var incrementX, incrementY = getIncrement(cl.XMove, cl.YMove, cl.X, cl.Y, moveToX, moveToY)
	//TODO Add the offset checks back in.
	var xOffset = moveToX - cl.X
	if xOffset < 0 {
		xOffset *= -1
	}
	if xOffset < cl.MaxXOffset {
		incrementX = 0
	}
	yOffset := moveToY - cl.Y
	if yOffset < 0 {
		yOffset *= -1
	}
	if yOffset < cl.MaxYOffset {
		incrementY = 0
	}

	return incrementX, incrementY
}

func (cl *NavigatorData) increment(moveToX, moveToY int32) (incrementX, incrementY int32) {
	x, y := cl.GetCurrentPoint()
	incrementX = cl.XMove
	incrementY = cl.YMove
	if moveToX < x {
		incrementX *= -1
	}

	if moveToY < y {
		incrementY *= -1
	}
	return incrementX, incrementY
}
