package bus

import "github.com/vorpalgame/vorpal/lib"

type MouseEventListener interface {
	OnMouseEvent(mouseChannel <-chan MouseEvent)
}

//////

func NewMouseEvent(leftButton, rightButton, middleButton lib.MouseButtonState, x int32, y int32) MouseEvent {
	return &mouseEvent{leftButton, rightButton, middleButton, lib.NewPoint(x, y)}
}

// //
type MouseEvent interface {
	LeftButton() lib.MouseButtonState
	RightButton() lib.MouseButtonState
	MiddleButton() lib.MouseButtonState
	GetX() int32
	GetY() int32
	GetCursorPoint() lib.Point
}

type mouseEvent struct {
	left, right, middle lib.MouseButtonState
	cursortLocation     lib.Point
}

func (evt *mouseEvent) LeftButton() lib.MouseButtonState {
	return evt.left
}
func (evt *mouseEvent) RightButton() lib.MouseButtonState {
	return evt.right
}

func (evt *mouseEvent) MiddleButton() lib.MouseButtonState {
	return evt.middle
}
func (evt *mouseEvent) GetX() int32 {
	return evt.cursortLocation.GetX()
}

func (evt *mouseEvent) GetY() int32 {
	return evt.cursortLocation.GetY()
}
func (evt *mouseEvent) GetCursorPoint() lib.Point {
	return evt.cursortLocation
}
