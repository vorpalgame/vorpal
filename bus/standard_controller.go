package bus

import (
	"log"
)

type StandardMediaPeerController interface {
	EngineListener
	ControllerListener
	GetImageDrawEvents() DrawEvent //TODO This needs to be an array of arrays of events for different layers.
	GetAudioEvent() AudioEvent     //One event at at time...
}

type controller struct {
	bus          VorpalBus
	engine       LifeCycle  //Rethink the lifecycle requirements (if any)
	imagesToDraw DrawEvent  //Should be multiple draw events sorted by layer and ordered within layer...
	audioEvent   AudioEvent //Only one audio event at a time but need controller flags for load, start, pause, unload...
	mouseEvent   MouseEvent //Keep the last state of mouse and keys.
	keyEvents    []string   //TODO use the actuaal key events and not the strings...

}

var c = controller{}

func NewGameController() StandardMediaPeerController {
	c.bus = GetVorpalBus()
	InitKeys()
	c.bus.AddControllerListener(&c)
	c.bus.AddEngineListener(&c)
	return &c
}

// Need standards for new/init/close...
func (c *controller) Start() {

	c.engine.Start()
}

// Usig the tag/name instead of actual imge assures race condtions aren't a problem.
func (c *controller) OnDrawEvent(drawChannel <-chan DrawEvent) {
	for evt := range drawChannel {

		log.Default().Println("Received draw event: " + evt.GetImage())
		c.imagesToDraw = evt
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

func (c *controller) OnMouseEvent(mouseChannel <-chan MouseEvent) {
	for evt := range mouseChannel {
		//log.Default().Println("Load Images in Controller: ")
		c.mouseEvent = evt
	}
}

// TODO return multiples...
func (c *controller) GetImageDrawEvents() DrawEvent {
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
