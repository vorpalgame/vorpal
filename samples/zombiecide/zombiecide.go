package zombiecide

import (
	"log"
	"os"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

type zombiecide struct {
	bus bus.VorpalBus
	//textEvent  bus.TextEvent
	mouseEvent bus.MouseEvent
	background bus.ImageLayer
	keyEvent   bus.KeyEvent
}

// TODO Need a configuration mechanism with YAML or JSON to eliminate the need for hard code.
var zombies = zombiecide{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"

// var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

// TODO Refactor this start up to make it more idiomatic.
func Init() {
	log.Println("New zombie game")
	//e for exit
	//r for reset zombie to beginning
	vbus := bus.GetVorpalBus()

	//We have to register both upper/lower case possiblities as it isn't clear why we are getting some results otherwise.
	vbus.SendKeysRegistrationEvent(bus.NewKeysRegistrationEvent("e", "E", "R", "r", "+", "=", "-", "_"))

	vbus.AddMouseListener(&zombies)
	vbus.AddKeyEventListener(&zombies)

	zombies.bus = vbus

	//TODO We need config probably through JSON file when prototyping is complete.
	zombies.background = bus.NewImageLayer().AddLayerData(bus.NewImageMetadata("samples/resources/zombiecide/background.png", 0, 0, 33))
	zombies.mouseEvent = nil

	textEvent := bus.NewTextEvent(fontName, 18, 0, 0).AddText("Henry follows mouse pointer. \nLeft Mouse Button to Attack. \nStand still too long and he dies!\n Press 'e' to exit or 'r' to restart.\n NOTE: George the parts zombie is still being worked on.").SetX(1200).SetY(100)
	vbus.SendTextEvent(textEvent)
	var currentState = NewZombie(50) //Convenience var until we refactor.
	zombieParts := newPartsZommbie()
	//
	for {
		if zombies.mouseEvent != nil {
			drawEvt := bus.NewDrawEvent()
			drawEvt.AddImageLayer(zombies.background)
			currentState.Execute(drawEvt, zombies.keyEvent, zombies.mouseEvent)
			zombies.keyEvent = nil
			drawEvt.AddImageLayer(zombieParts.CreateImageLayer(zombies.mouseEvent))
			vbus.SendDrawEvent(drawEvt)

			time.Sleep(20 * time.Millisecond)
			//Execute to send image and sound

		}

	}

}

func (z *zombiecide) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		log.Default().Println(evt.GetKey().ToString())
		if evt.GetKey().EqualsIgnoreCase("e") {
			os.Exit(0)
		} else if evt.GetKey().EqualsIgnoreCase("r") {
			//TODO Stop and close old resources if necessary...
			//zombies.zombie = NewZombie()
		} else {
			z.keyEvent = evt
		}

	}

}
func (z *zombiecide) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		z.mouseEvent = evt

	}
}
