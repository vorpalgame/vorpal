package raylibengine

import "github.com/vorpalgame/vorpal/bus"

// The peer raylibPeerController is to mediate any impedance mismatch between the tight single
// threaded loop of the Raylib engine and the concurrent mechanisms of Golang.
// Note: Raylib is very fast but very intolerant of mutations on any data it is using.
// So the peer and MediaCache can keep off thread data for rendering, listening and
// sending. Raylib loop will query this raylibPeerController for data it needs.
func NewRaylibPeerController() RaylibPeerController {
	controller := raylibPeerController{}
	controller.MediaCache = NewMediaCache()
	return &controller
}

type RaylibPeerController interface {
	bus.DrawEventListener
	bus.AudioEventListener
	bus.TextEventListener
	bus.KeysRegistrationEventListener
	bus.ControlEventListener
	MediaCache
	GetControlEvents() []bus.ControlEvent
	GetDrawEvent() bus.DrawEvent
	GetAudioEvent() bus.AudioEvent //One event at at time...
	GetTextEvent() bus.TextEvent
	GetKeysRegistrationEvent() bus.KeysRegistrationEvent
}

type raylibPeerController struct {
	MediaCache
	bus                   bus.VorpalBus
	drawEvent             bus.DrawEvent
	audioEvent            []bus.AudioEvent          //Different audio events for stop, start, etc. so they need to be kept in slice for processing.
	textEvent             bus.TextEvent             //TODO put multiple keys in one event...
	keysRegistrationEvent bus.KeysRegistrationEvent //Only one set of keys to listen for at a time.
	controlEvents         []bus.ControlEvent        //May need slice..
}

func (c *raylibPeerController) OnControlEvent(controlChannel <-chan bus.ControlEvent) {
	for evt := range controlChannel {
		c.controlEvents = append(c.controlEvents, evt)
	}
}

func (c *raylibPeerController) OnDrawEvent(drawChannel <-chan bus.DrawEvent) {
	for evt := range drawChannel {
		c.drawEvent = evt
	}

}

func (c *raylibPeerController) OnKeyRegistrationEvent(keyRegistrationChannel <-chan bus.KeysRegistrationEvent) {
	for evt := range keyRegistrationChannel {
		c.keysRegistrationEvent = evt
	}
}

func (c *raylibPeerController) OnAudioEvent(audioChannel <-chan bus.AudioEvent) {
	for evt := range audioChannel {
		c.audioEvent = append(c.audioEvent, evt)
	}
}

func (c *raylibPeerController) OnTextEvent(textChannel <-chan bus.TextEvent) {
	for evt := range textChannel {
		c.textEvent = evt

	}
}

func (c *raylibPeerController) GetDrawEvent() bus.DrawEvent {
	//evt := c.drawEvent
	//c.drawEvent = nil
	return c.drawEvent
}

func (c *raylibPeerController) GetAudioEvent() bus.AudioEvent {
	var evt bus.AudioEvent = nil
	if len(c.audioEvent) > 0 {
		evt, c.audioEvent = c.audioEvent[0], c.audioEvent[1:]
	}
	return evt
}

// Don't repeat process.
func (c *raylibPeerController) GetControlEvents() []bus.ControlEvent {
	temp := c.controlEvents
	c.controlEvents = nil
	return temp
}

// Need better coordination mechanism for rendering
// of draw and text events. Probably should compose
// them together and then eliminate the need for the
// use of temp

func (c *raylibPeerController) GetTextEvent() bus.TextEvent {
	temp := c.textEvent
	c.textEvent = nil
	return temp
}

func (c *raylibPeerController) GetKeysRegistrationEvent() bus.KeysRegistrationEvent {
	return c.keysRegistrationEvent
}
