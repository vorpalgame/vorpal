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

func StartGame() {
	log.Println("New card game")
	cards.bus = bus.GetVorpalBus()
	cards.bus.AddEngineListener(&cards)
	cards.tarotDeck = NewDeck()
	cards.tarotDeck.Init()

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
			g.bus.SendDrawEvent(bus.NewDrawEvent(displayCard, 10, 10, 1, 400, 680))
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
