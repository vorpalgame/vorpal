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
	IsButtonToggled() bool
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

//TODO...
func (evt *mouseEvent) IsButtonToggled() bool {
	return evt.left.isReleased() || evt.left.IsPressed() || evt.middle.isReleased() || evt.middle.IsPressed() || evt.right.isReleased() || evt.right.IsPressed()
}

//////

type MouseButtonState interface {
	Name() string
	IsPressed() bool
	isReleased() bool
}

type mouseButtonState struct {
	name              string
	pressed, released bool
}

func NewMouseButtonState(name string, pressed bool, released bool) MouseButtonState {
	return &mouseButtonState{name, pressed, released}
}
func (btn *mouseButtonState) IsPressed() bool {
	return btn.pressed
}

func (btn *mouseButtonState) isReleased() bool {
	return btn.released
}

func (btn *mouseButtonState) Name() string {
	return btn.name
}
