package bus

type MouseEventListener interface {
	OnMouseEvent(mouseChannel <-chan MouseEvent)
}

//////

func NewMouseEvent(leftButton, rightButton, middleButton MouseButtonState, x int32, y int32) MouseEvent {
	return &mouseEvent{leftButton, rightButton, middleButton, x, y}
}

// //
type MouseEvent interface {
	LeftButton() MouseButtonState
	RightButton() MouseButtonState
	MiddleButton() MouseButtonState
	GetX() int32
	GetY() int32
}

type mouseEvent struct {
	left, right, middle MouseButtonState
	x, y                int32
}

func (evt *mouseEvent) LeftButton() MouseButtonState {
	return evt.left
}
func (evt *mouseEvent) RightButton() MouseButtonState {
	return evt.right
}

func (evt *mouseEvent) MiddleButton() MouseButtonState {
	return evt.middle
}
func (evt *mouseEvent) GetX() int32 {
	return evt.x
}

func (evt *mouseEvent) GetY() int32 {
	return evt.y
}

//////

type MouseButtonState interface {
	Name() string
	IsDown() bool
	IsUp() bool
}

type mouseButtonState struct {
	name string
	down bool
}

func NewMouseButtonState(name string, down bool) MouseButtonState {
	return &mouseButtonState{name, down}
}
func (btn *mouseButtonState) IsDown() bool {
	return btn.down
}

func (btn *mouseButtonState) IsUp() bool {
	return !btn.down
}

func (btn *mouseButtonState) Name() string {
	return btn.name
}
