package bus

import (
	"log"
)

type StandardMediaPeerController interface {
	EngineListener
	ControllerListener

	GetDrawEvent() DrawEvent
	GetAudioEvent() AudioEvent //One event at at time...
	GetTextEvent() TextEvent
	GetImageCacheEvent() ImageCacheEvent
}

type controller struct {
	bus             VorpalBus
	drawEvent       DrawEvent
	audioEvent      AudioEvent      //Only one audio event at a time but need controller flags for load, start, pause, unload...
	mouseEvent      MouseEvent      //Keep the last state of mouse and keys.
	textEvent       TextEvent       //TODO put multiple keys in one event...
	imageCacheEvent ImageCacheEvent //Could have multiples so should be slice...
	keyEvents       []string        //TODO use the actual key events and not the strings...
}

var c = controller{}

func NewGameController() StandardMediaPeerController {
	c.bus = GetVorpalBus()
	InitKeys()
	c.bus.AddControllerListener(&c)
	c.bus.AddEngineListener(&c)
	return &c
}

// Usig the tag/name instead of actual imge assures race condtions aren't a problem.
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

func (c *controller) OnKeyEvent(keyChannel <-chan KeyEvent) {
	for evt := range keyChannel {
		log.Default().Println("keyEvent ")

		c.keyEvents = append(c.keyEvents, evt.GetKey().ToString())
	}
}

func (c *controller) OnAudioEvent(audioChannel <-chan AudioEvent) {
	for evt := range audioChannel {
		log.Default().Println("audioEvent ")

		c.audioEvent = evt
	}
}

func (c *controller) OnTextEvent(textChannel <-chan TextEvent) {
	for evt := range textChannel {
		log.Default().Println(evt)

		c.textEvent = evt

	}
}

func (c *controller) OnMouseEvent(mouseChannel <-chan MouseEvent) {
	for evt := range mouseChannel {
		//log.Default().Println("Load Images in Controller: ")
		c.mouseEvent = evt
	}
}

// TODO return multiples...
func (c *controller) GetDrawEvent() DrawEvent {
	//log.Default().Println("Get Draw Image")
	return c.drawEvent
}

func (c *controller) GetImageCacheEvent() ImageCacheEvent {
	//log.Default().Println("Get Draw Image")
	return c.imageCacheEvent
}

func (c *controller) GetAudioEvent() AudioEvent {
	//log.Default().Println("Get Draw Image")
	//TODO audio event should be cleared by engine ack on callback.
	//Just wire it in for now.
	temp := c.audioEvent
	c.audioEvent = nil
	return temp
}

func (c *controller) GetTextEvent() TextEvent {
	//log.Default().Println("Get Text Event")

	return c.textEvent

}
