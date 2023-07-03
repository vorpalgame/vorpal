package bus

type StandardMediaPeerController interface {
	ControllerListener
	GetControlEvent() ControlEvent
	GetDrawEvent() DrawEvent
	GetAudioEvent() AudioEvent //One event at at time...
	GetTextEvent() TextEvent
	GetImageCacheEvent() ImageCacheEvent
	GetKeysRegistrationEvent() KeysRegistrationEvent
}

// TODO It will make sense for to keep events in a slice of events.
// In many cases we only want or need the last event. However, keeping a slice of events
// that come in between calls would permit the engine to determine if it only wants the
// last. Some, like KeysRegistrationEvent ever only need 1 because all listened for keys are sent
// ControlEvents will need multiple and be cleared so they don't repeat processing.
type controller struct {
	bus                   VorpalBus
	drawEvent             DrawEvent
	audioEvent            []AudioEvent          //Different audio events for stop, start, etc. so they need to be kept in slice for processing.
	textEvent             TextEvent             //TODO put multiple keys in one event...
	imageCacheEvent       ImageCacheEvent       //Could have multiples so should be slice...
	keysRegistrationEvent KeysRegistrationEvent //Only one set of keys to listen for at a time.
	controlEvent          ControlEvent          //May need slice..
}

var c = controller{}

func NewGameController() StandardMediaPeerController {
	c.bus = GetVorpalBus()

	c.bus.AddControllerListener(&c)
	c.audioEvent = make([]AudioEvent, 0, 50)
	return &c
}
func (c *controller) OnControlEvent(controlChannel <-chan ControlEvent) {
	for evt := range controlChannel {
		c.controlEvent = evt
	}
}

func (c *controller) OnDrawEvent(drawChannel <-chan DrawEvent) {
	for evt := range drawChannel {
		c.drawEvent = evt
	}

}

func (c *controller) OnImageCacheEvent(cacheChannel <-chan ImageCacheEvent) {
	for evt := range cacheChannel {
		c.imageCacheEvent = evt
	}
}
func (c *controller) OnKeyRegistrationEvent(keyRegistrationChannel <-chan KeysRegistrationEvent) {
	for evt := range keyRegistrationChannel {
		c.keysRegistrationEvent = evt
	}
}

func (c *controller) OnAudioEvent(audioChannel <-chan AudioEvent) {
	for evt := range audioChannel {
		c.audioEvent = append(c.audioEvent, evt)
	}
}

func (c *controller) OnTextEvent(textChannel <-chan TextEvent) {
	for evt := range textChannel {
		c.textEvent = evt

	}
}

func (c *controller) GetDrawEvent() DrawEvent {
	evt := c.drawEvent
	c.drawEvent = nil
	return evt
}

func (c *controller) GetAudioEvent() AudioEvent {
	var evt AudioEvent = nil
	if len(c.audioEvent) > 0 {
		evt, c.audioEvent = c.audioEvent[0], c.audioEvent[1:]
	}
	return evt
}

// Don't repeat process.
func (c *controller) GetControlEvent() ControlEvent {
	temp := c.controlEvent
	c.controlEvent = nil
	return temp
}

// Need better coordination mechanism for rendering
// of draw and text events. Probably should compose
// them together and then eliminate the need for the
// use of temp

func (c *controller) GetTextEvent() TextEvent {
	temp := c.textEvent
	c.textEvent = nil
	return temp
}

func (c *controller) GetKeysRegistrationEvent() KeysRegistrationEvent {
	return c.keysRegistrationEvent
}

// TODO In process
func (c *controller) GetImageCacheEvent() ImageCacheEvent {
	return c.imageCacheEvent
}
