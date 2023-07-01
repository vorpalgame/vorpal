package bus

// TODO We need to look at race condition handling. While in many cases, lke the DrawEvent, we
// usually won't care because they should be processed much faster than we are sending and if not,
// we don't want to fill a queue up with draw events hoping they get caught up.
// But even limiting the queue to a single event, while likely correct, means that it is possible that
// they are read off the queue, set to the event in here and then overwritten before processed. Highly
// unlikely given the relative speed of the two sides but still...TODO

type StandardMediaPeerController interface {
	ControllerListener

	GetDrawEvent() DrawEvent
	GetAudioEvent() AudioEvent //One event at at time...
	GetTextEvent() TextEvent
	GetImageCacheEvent() ImageCacheEvent
	GetKeysRegistrationEvent() KeysRegistrationEvent
}

type controller struct {
	bus                   VorpalBus
	drawEvent             DrawEvent
	audioEvent            []AudioEvent          //Different audio events for stop, start, etc. so they need to be kept in slice for processing.
	mouseEvent            MouseEvent            //Keep the last state of mouse and keys.
	textEvent             TextEvent             //TODO put multiple keys in one event...
	imageCacheEvent       ImageCacheEvent       //Could have multiples so should be slice...
	keysRegistrationEvent KeysRegistrationEvent //Only one set of keys to listen for at a time.
}

var c = controller{}

func NewGameController() StandardMediaPeerController {
	c.bus = GetVorpalBus()

	c.bus.AddControllerListener(&c)
	c.audioEvent = make([]AudioEvent, 0, 50)
	return &c
}

//We don't want to consume from the channel if we still have an event waiting for processing.
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

func (c *controller) GetTextEvent() TextEvent {
	return c.textEvent
}
func (c *controller) GetKeysRegistrationEvent() KeysRegistrationEvent {
	return c.keysRegistrationEvent
}

// TODO In process
func (c *controller) GetImageCacheEvent() ImageCacheEvent {
	return c.imageCacheEvent
}
