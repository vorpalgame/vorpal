package zombiecide

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

type zombiecide struct {
	bus bus.VorpalBus
	//textEvent  bus.TextEvent
	mouseEvent    bus.MouseEvent
	background    bus.ImageLayer
	keyEvent      bus.KeyEvent
	currentZombie string
}

// TODO Need a configuration mechanism with YAML or JSON to eliminate the need for hard code.
var zombies = zombiecide{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"

// var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

// TODO Refactor this start up to make it more idiomatic.
func Init() {
	log.Println("New zombie game")

	vbus := bus.GetVorpalBus()
	configKeys := lib.NewKeys(viper.GetStringSlice("RegisterKeys"))

	log.Default().Println(configKeys)
	evt := bus.NewKeysRegistrationEvent(configKeys)
	log.Default().Println(evt.GetKeys())
	vbus.SendKeysRegistrationEvent(evt)

	vbus.AddMouseListener(&zombies)
	vbus.AddKeyEventListener(&zombies)
	zombies.currentZombie = "h"
	zombies.bus = vbus

	//TODO The data elements here should come from the yaml.
	zombies.background = bus.NewImageLayer().AddLayerData(bus.NewImageMetadata("samples/resources/zombiecide/background.png", 0, 0, 33))
	zombies.mouseEvent = nil
	textEvent := bus.NewMultilineTextEvent(fontName, 18, 0, 0).AddText("Press 'g' for George or 'h' for Henry. \n Zombies follow the mouse pointer. \nLeft Mouse Button causes Henry to Attack. \nStand still too long and he dies!\n Press 'e' to exit or 'r' to restart.\n NOTE: George the parts zombie is still being worked on.").SetLocation(1200, 100)
	vbus.SendTextEvent(textEvent)

	subsumptionZombie := newSubsumptionZombie()
	stateMachineZombie := NewZombieStateMachine()
	//
	for {
		if zombies.mouseEvent != nil {
			drawEvt := bus.NewDrawLayersEvent()
			drawEvt.AddImageLayer(zombies.background)
			if zombies.currentZombie == "h" {
				stateMachineZombie.Execute(drawEvt, zombies.mouseEvent, zombies.keyEvent)
			} else {
				drawEvt.AddImageLayer(subsumptionZombie.CreateImageLayer(zombies.mouseEvent)) //This shoulc change to look more like state zombie.
			}
			vbus.SendDrawEvent(drawEvt)
			zombies.keyEvent = nil
			time.Sleep(20 * time.Millisecond)
			//Execute to send image and sound

		}

	}

}

func (z *zombiecide) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {
		//Using explicit letters due to misreported case from raylib...
		log.Default().Println(evt.GetKey().ToString())
		if evt.GetKey().EqualsIgnoreCase("e") {
			os.Exit(0)
		} else if evt.GetKey().EqualsIgnoreCase("r") {
			//TODO Stop and close old resources if necessary...
			//zombies.zombie = NewZombie()
		} else if evt.GetKey().EqualsIgnoreCase("h") {
			z.currentZombie = "h"
		} else if evt.GetKey().EqualsIgnoreCase("g") {
			z.currentZombie = "g"
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
