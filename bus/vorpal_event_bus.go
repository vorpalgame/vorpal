package bus

// TODO These should be Map keyed by listener so that we don't end up with mulitple registrations.
type vorpalEventBus struct {
	keyEventListenerChannels  []chan KeyEvent
	mouseListenerChannels     []chan MouseEvent
	drawEventListenerChannels []chan DrawEvent
}

type VorpalBus interface {
	//Conveinience listeners...
	AddGameListener(eventListener GameListener)
	AddEngineListener(eventListener EngineListener)
	///
	AddKeyEventListener(eventListener KeyEventListener)
	AddMouseListener(eventListener MouseEventListener)
	AddDrawEventListener(eventListener DrawEventListener)

	SendMouseEvent(event MouseEvent)
	SendKeyEvent(event KeyEvent)
	SendDrawEvent(event DrawEvent)
}

var eb = vorpalEventBus{}

func GetVorpalBus() VorpalBus {
	return &eb
}

func (eb *vorpalEventBus) AddGameListener(eventListener GameListener) {
	eb.AddDrawEventListener(eventListener)
}

func (eb *vorpalEventBus) AddEngineListener(eventListener EngineListener) {
	eb.AddMouseListener(eventListener)
	eb.AddKeyEventListener(eventListener)
}

// /KEY EVENTS
func (bus *vorpalEventBus) AddKeyEventListener(eventListener KeyEventListener) {
	listenerChannel := make(chan KeyEvent, 100)
	bus.keyEventListenerChannels = append(bus.keyEventListenerChannels, listenerChannel)
	go eventListener.OnKeyEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendKeyEvent(event KeyEvent) {
	for _, channel := range bus.keyEventListenerChannels {
		channel <- event
	}
}

// ///MOUSE BUTTON EVENTS
func (bus *vorpalEventBus) AddMouseListener(eventListener MouseEventListener) {
	listenerChannel := make(chan MouseEvent, 100)
	bus.mouseListenerChannels = append(bus.mouseListenerChannels, listenerChannel)
	go eventListener.OnMouseEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendMouseEvent(event MouseEvent) {
	for _, channel := range bus.mouseListenerChannels {
		channel <- event
	}
}

// DrawEvent //TODO change to take the image
func (bus *vorpalEventBus) AddDrawEventListener(eventListener DrawEventListener) {
	listenerChannel := make(chan DrawEvent, 100)
	bus.drawEventListenerChannels = append(bus.drawEventListenerChannels, listenerChannel)
	go eventListener.OnDrawEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendDrawEvent(event DrawEvent) {
	for _, channel := range bus.drawEventListenerChannels {
		channel <- event
	}
}

//
