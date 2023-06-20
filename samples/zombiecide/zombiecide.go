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
	zombie     ZombieSprite
}

// TODO The cards should have locations or the game/board should dictate those...
// TODO The fonts and text layers need to be added to the text event.
var zombies = zombiecide{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"
var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

// TODO Refactor this start up to make it more idiomatic.
func Init() {
	log.Println("New z game")
	zombies.zombie = NewZombie()
	vbus := bus.GetVorpalBus()
	vbus.AddMouseListener(&zombies)
	vbus.AddKeyEventListener(&zombies)
	//e for exit
	//r for reset zombie to beginning
	vbus.SendKeysRegistrationEvent(bus.NewKeysRegistrationEvent("e", "r"))
	zombies.bus = vbus

	//TODO We need config probably through JSON file when prototyping is complete.
	zombies.background = bus.NewImageLayer("samples/resources/zombiecide/background.png", 0, 0, 1920, 1080)
	zombies.mouseEvent = nil
	//TODO text event is flickering...
	textEvent := bus.NewTextEvent(fontName, 18, 0, 0).AddText("Henry follows mouse pointer. \nLeft Mouse Button to Attack. \nStand still too long and he dies!\n Press 'e' to exit or 'r' to restart.").SetX(1200).SetY(100)
	vbus.SendTextEvent(textEvent)
	for {
		if zombies.mouseEvent != nil {
			evt := bus.NewDrawEvent()
			evt.AddImageLayer(zombies.background)

			zombies.zombie = zombies.zombie.RunSprite(evt, zombies.mouseEvent)
			time.Sleep(20 * time.Millisecond)
		}

	}

}

func (z *zombiecide) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		log.Default().Println(evt.GetKey().ToString())
		if evt.GetKey().EqualsIgnoreCase("e") {
			os.Exit(0)
		}
		if evt.GetKey().EqualsIgnoreCase("r") {
			zombies.zombie = NewZombie()
		}

	}

}
func (z *zombiecide) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		z.mouseEvent = evt

	}
}
