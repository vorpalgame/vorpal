package bus

import "github.com/vorpalgame/vorpal/lib"

type MouseEventListener interface {
	OnMouseEvent(mouseChannel <-chan MouseEvent)
}

//////

func NewMouseEvent(leftButton, rightButton, middleButton lib.MouseButtonState, x int32, y int32) MouseEvent {
	return &mouseEvent{leftButton, rightButton, middleButton, x, y}
}

// //
type MouseEvent interface {
	LeftButton() lib.MouseButtonState
	RightButton() lib.MouseButtonState
	MiddleButton() lib.MouseButtonState
	GetCursorPoint() (x, y int32)
}

type mouseEvent struct {
	left, right, middle lib.MouseButtonState
	x, y                int32
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

func (evt *mouseEvent) GetCursorPoint() (x, y int32) {
	return evt.x, evt.y
}
