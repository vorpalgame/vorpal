package lib

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
