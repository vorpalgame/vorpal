package bus

// Initial cut at this...may need different data types
// and that may require generics...
// TODO Test using a tagging inteface with different extensions...
type ControlEventListener interface {
	OnControlEvent(controlChannel <-chan ControlEvent)
}
type ControlEventProcessor interface {
	ProcessControlEvent(evt ControlEvent)
}

func NewWindowTitleEvent(data string) ControlEvent {
	return &windowTitleData{data}
}

type ControlEvent interface{}
type WindowTitleEvent interface {
	ControlEvent
	GetTitle() string
}

type windowTitleData struct {
	title string
}

func (w *windowTitleData) GetTitle() string {
	return w.title
}
