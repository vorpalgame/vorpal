package zombiecide

import (
	"github.com/vorpalgame/vorpal/samples/zombiecide/state_machines"
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
	keyEvent      bus.KeyEvent
	currentZombie string
}

//Sample file for different possible use cases...

var zombies = zombiecide{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"

// var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"
const (
	henry = "/samples/etc/henry.yaml"
	karen = "/samples/etc/karen.yaml"
)

// TODO Refactor this start up to make it more idiomatic and bootstrap from the ymal
func Init() {
	log.Println("New zombie game")

	vbus := bus.GetVorpalBus()

	vbus.SendControlEvent(bus.NewWindowSizeEvent(1920, 1080))
	vbus.SendControlEvent(bus.NewWindowTitleEvent("Zombicicide!"))
	//log.Default().Println(configKeys)
	configKeys := lib.NewKeys(viper.GetStringSlice("RegisterKeys"))
	evt := bus.NewKeysRegistrationEvent(configKeys)
	//log.Default().Println(evt.GetKeys())
	vbus.SendKeysRegistrationEvent(evt)

	vbus.AddMouseListener(&zombies)
	vbus.AddKeyEventListener(&zombies)
	zombies.currentZombie = "h"
	zombies.bus = vbus

	zombies.mouseEvent = nil
	textEvent := bus.NewMultilineTextEvent(fontName, 18, 0, 0).AddText("Press 'g' for George or 'h' for Henry. \n Zombies follow the mouse pointer. \nLeft Mouse Button causes Henry to Attack. \nStand still too long and he dies!\n Press 'e' to exit or 'r' to restart.\n NOTE: George the parts zombie is still being worked on.").SetLocation(1200, 100)
	vbus.SendTextEvent(textEvent)

	subsumptionZombie := newSubsumptionZombie()

	//MoveByIncrement to zombicide yaml
	dir, _ := os.Getwd()

	statesFile := dir + henry

	log.Default().Println(dir)
	f, e := os.ReadFile(statesFile)
	if e != nil {
		log.Default().Println(e)
		os.Exit(1)
	}
	stateMachineZombie := state_machines.UnmarshalZombie(f)
	//Attachable functions for testing conditions should be added so
	//they can be queried.
	//TODO we need to switch both background types to use absolute size while sprites can use percent
	//scale or perhaps both scale and width/height.
	ac := lib.ActionStageControllerData{}

	//TODO We need to revamp the configurator to eliminate Viper and to handle paths to
	//resources.
	//Need new behavior map for different environment
	ac.LoadControlMapFromFile("samples/resources/zombiecide/behaviorland.png", 1920, 1080)

	stateMachineZombie.Navigator.ActionStageController = &ac

	//
	for {
		if zombies.mouseEvent != nil {
			drawEvt := bus.NewDrawLayersEvent()
			createBackground(drawEvt)
			if zombies.currentZombie == "h" {
				stateMachineZombie.Execute(drawEvt, zombies.mouseEvent, zombies.keyEvent)
				createForeground(drawEvt)
			} else {
				drawEvt.AddImageLayer(subsumptionZombie.CreateImageLayer(zombies.mouseEvent))
			}
			vbus.SendDrawEvent(drawEvt)
			zombies.keyEvent = nil
			time.Sleep(20 * time.Millisecond)
			//Execute to send image and sound

		}

	}

}

func createBackground(evt bus.DrawLayersEvent) {
	layer := lib.NewImageLayer()
	layer.AddLayerData(lib.NewImageMetadata("samples/resources/zombiecide/background.png", 0, 0, 1920, 1080))
	evt.AddImageLayer(layer)

}
func createForeground(evt bus.DrawLayersEvent) {
	layer := lib.NewImageLayer()
	layer.AddLayerData(lib.NewImageMetadata("samples/resources/zombiecide/foreground.png", 0, 0, 1920, 1080))
	evt.AddImageLayer(layer)
}
func (z *zombiecide) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {
		//Using explicit letters due to misreported case from raylib...
		//	log.Default().Println(evt.GetKey().ToString())
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
