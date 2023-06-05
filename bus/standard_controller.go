package bus

import (
	"log"
)

type StandardGameController interface {
	EngineListener
	GameListener
	GetImageDrawEvents() DrawEvent //TODO This needs to be an array of arrays of events for different layers.
	ClearImagesToLoad()
}

type controller struct {
	bus          VorpalBus
	engine       LifeCycle
	loadImages   []string //We need structs to store for these and not strings...
	keyEvents    []string
	imagesToDraw DrawEvent  //Should be multiple draw events sorted by layer and then receive order within layer...
	mouseEvent   MouseEvent //Keep the last state of mouse and keys.
}

var c = controller{}

func NewGameController() StandardGameController {
	c.bus = GetVorpalBus()
	InitKeys()
	c.bus.AddGameListener(&c)
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

func (c *controller) OnMouseEvent(mouseChannel <-chan MouseEvent) {
	for evt := range mouseChannel {
		//log.Default().Println("Load Images in Controller: ")
		c.mouseEvent = evt
	}
}

// /Called by engine...Need better struct values to return...
func (c *controller) ClearImagesToLoad() {
	c.loadImages = nil
}

func (c *controller) GetImagesToLoad() []string {
	//log.Default().Println("Load images from file name")
	return c.loadImages

}

// TODO return multiples...
func (c *controller) GetImageDrawEvents() DrawEvent {
	//log.Default().Println("Get Draw Image")
	return c.imagesToDraw
}
