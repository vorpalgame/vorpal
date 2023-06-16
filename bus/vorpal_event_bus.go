package bus

type vorpalEventBus struct {
	keyEventListenerChannels        []chan KeyEvent
	mouseListenerChannels           []chan MouseEvent
	drawEventListenerChannels       []chan DrawEvent
	audioEventListenerChannels      []chan AudioEvent
	textEventListenerChannels       []chan TextEvent
	imageCacheEventListenerChannels []chan ImageCacheEvent
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
	AddTextEventListener(eventListener TextEventListener)
	AddImageCacheEventListener(eventListener ImageCacheEventListener)

	SendMouseEvent(event MouseEvent)
	SendKeyEvent(event KeyEvent)
	SendDrawEvent(event DrawEvent)
	SendAudioEvent(event AudioEvent)
	SendTextEvent(event TextEvent)
	SendImageCacheEvent(event ImageCacheEvent)
}

var eb = vorpalEventBus{}

func GetVorpalBus() VorpalBus {
	return &eb
}

func (eb *vorpalEventBus) AddControllerListener(eventListener ControllerListener) {
	eb.AddDrawEventListener(eventListener)
	eb.AddAudioEventListener(eventListener)
	eb.AddTextEventListener(eventListener)
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

func (bus *vorpalEventBus) AddImageCacheEventListener(eventListener ImageCacheEventListener) {
	listenerChannel := make(chan ImageCacheEvent, 100)
	bus.imageCacheEventListenerChannels = append(bus.imageCacheEventListenerChannels, listenerChannel)
	go eventListener.OnImageCacheEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendImageCacheEvent(event ImageCacheEvent) {
	for _, channel := range bus.imageCacheEventListenerChannels {
		channel <- event
	}
}

// ///MOUSE BUTTON EVENTS
func (bus *vorpalEventBus) AddMouseListener(eventListener MouseEventListener) {
	listenerChannel := make(chan MouseEvent, 1)
	bus.mouseListenerChannels = append(bus.mouseListenerChannels, listenerChannel)
	go eventListener.OnMouseEvent(listenerChannel)
}

func (bus *vorpalEventBus) SendMouseEvent(event MouseEvent) {
	for _, channel := range bus.mouseListenerChannels {
		channel <- event
	}
}

// The controller events to the engine probably only need a single channel.
func (bus *vorpalEventBus) AddDrawEventListener(eventListener DrawEventListener) {
	listenerChannel := make(chan DrawEvent, 1)
	bus.drawEventListenerChannels = append(bus.drawEventListenerChannels, listenerChannel)
	go eventListener.OnDrawEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendDrawEvent(event DrawEvent) {
	for _, channel := range bus.drawEventListenerChannels {
		channel <- event
	}
}

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
func (bus *vorpalEventBus) AddTextEventListener(eventListener TextEventListener) {
	listenerChannel := make(chan TextEvent, 100)
	bus.textEventListenerChannels = append(bus.textEventListenerChannels, listenerChannel)
	go eventListener.OnTextEvent(listenerChannel)
}
func (bus *vorpalEventBus) SendTextEvent(event TextEvent) {
	for _, channel := range bus.textEventListenerChannels {
		channel <- event
	}
}

//
