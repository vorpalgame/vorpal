package bus

// TODO These should  so that we don't end up with mulitple registrations and to permit
//easy derigistration. Either that or hand rolled linked list where registration iterates
//the list ensuring that none of the current elements match before adding it to the end.

type vorpalEventBus struct {
	keyEventListenerChannels   []chan KeyEvent
	mouseListenerChannels      []chan MouseEvent
	drawEventListenerChannels  []chan DrawEvent
	audioEventListenerChannels []chan AudioEvent
}

type VorpalBus interface {
	//Convenience listener collectionss...
	AddControllerListener(eventListener ControllerListener)
	AddEngineListener(eventListener EngineListener)
	///Individual listeners....
	AddKeyEventListener(eventListener KeyEventListener)
	AddMouseListener(eventListener MouseEventListener)
	AddDrawEventListener(eventListener DrawEventListener)
	AddAudioEventListener(eventListener AudioEventListener)

	SendMouseEvent(event MouseEvent)
	SendKeyEvent(event KeyEvent)
	SendDrawEvent(event DrawEvent)
	SendAudioEvent(event AudioEvent)
}

var eb = vorpalEventBus{}

func GetVorpalBus() VorpalBus {
	return &eb
}

func (eb *vorpalEventBus) AddControllerListener(eventListener ControllerListener) {
	eb.AddDrawEventListener(eventListener)
	eb.AddAudioEventListener(eventListener)
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

//The controller events to the engine probably only need a single channel.
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

//Likely only need single channel...TBD
func (bus *vorpalEventBus) AddAudioEventListener(eventListener AudioEventListener) {
	listenerChannel := make(chan AudioEvent, 100)
	bus.audioEventListenerChannels = append(bus.audioEventListenerChannels, listenerChannel)
	go eventListener.OnAudioEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendAudioEvent(event AudioEvent) {
	for _, channel := range bus.audioEventListenerChannels {
		channel <- event
	}
}

//
