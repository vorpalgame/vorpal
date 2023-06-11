package tarot

import (
	"fmt"
	"log"

	"github.com/vorpalgame/vorpal/bus"
)

type tarot struct {
	bus       bus.VorpalBus
	tarotDeck TarotDeck
}

// TODO The cards should have locations or the game/board should dictate those...
var cards = tarot{}

func InitGame() {
	log.Println("New card game")
	cards.bus = bus.GetVorpalBus()
	cards.bus.AddEngineListener(&cards)
	cards.tarotDeck = NewDeck()

}

// TODO Need a better mechanism for start up mechanics so that listeners get registered before we send
func StartGame() {
	drawEvent := bus.NewDrawEvent()
	drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/tarot/table.png", 0, 0, 1900, 1200))
	cards.bus.SendDrawEvent(drawEvent)
	cards.bus.SendTextEvent(bus.NewTextEvent("Press S to shuffle the deck and show a card.\nPress N to show the next card.", 800, 200))

}

func (g *tarot) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		//fmt.Println(evt.GetKey().ToString())
		//TODO finish the upper/lower case
		var sendCard = false
		if evt.GetKey().EqualsIgnoreCase("S") {
			g.tarotDeck.Shuffle()

			g.bus.SendAudioEvent(bus.NewAudioEvent("samples/resources/audio/shuffle.mp3"))
			sendCard = true
		} else if evt.GetKey().EqualsIgnoreCase("N") {
			sendCard = true
		}
		if sendCard {
			displayCard := g.tarotDeck.GetTopCard().GetCardImg()
			fmt.Println("Send card image name: " + displayCard)
			drawEvent := bus.NewDrawEvent()
			drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/tarot/table.png", 0, 0, 1900, 1200))
			drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, 300, 300, 400, 680))
			g.bus.SendDrawEvent(drawEvent)
		}
	}

}

func (g *tarot) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		if evt.IsButtonToggled() {
			fmt.Println(evt)
		}
	}
}
