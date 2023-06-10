package bus

import (
	"log"
)

type StandardMediaPeerController interface {
	EngineListener
	ControllerListener

	GetImageDrawEvents() map[int32]DrawEvent
	GetAudioEvent() AudioEvent //One event at at time...
	GetTextEvent() TextEvent
}

type controller struct {
	bus          VorpalBus
	imagesToDraw map[int32]DrawEvent
	audioEvent   AudioEvent //Only one audio event at a time but need controller flags for load, start, pause, unload...
	mouseEvent   MouseEvent //Keep the last state of mouse and keys.
	textEvent    TextEvent
	keyEvents    []string //TODO use the actuaal key events and not the strings...
	newDrawEvent bool
}

var c = controller{}

func NewGameController() StandardMediaPeerController {
	c.bus = GetVorpalBus()
	InitKeys()
	c.imagesToDraw = make(map[int32]DrawEvent) //100 possible layers. Could have layers within layers...
	c.newDrawEvent = false
	c.bus.AddControllerListener(&c)
	c.bus.AddEngineListener(&c)
	return &c
}

// Usig the tag/name instead of actual imge assures race condtions aren't a problem.
func (c *controller) OnDrawEvent(drawChannel <-chan DrawEvent) {
	for evt := range drawChannel {

		log.Default().Println("Received draw event: " + evt.GetImage())
		c.imagesToDraw[evt.GetZ()] = evt
		c.newDrawEvent = true
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
func (c *controller) GetImageDrawEvents() map[int32]DrawEvent {
	//log.Default().Println("Get Draw Image")
	return c.imagesToDraw
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
