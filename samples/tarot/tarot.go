package tarot

import (
	"fmt"
	"log"

	"github.com/vorpalgame/vorpal/bus"
)

type tarot struct {
	bus       bus.VorpalBus
	tarotDeck TarotDeck
	evt       bus.DrawEvent
}

// TODO The cards should have locations or the game/board should dictate those...
var cards = tarot{}

func InitGame() {
	log.Println("New card game")
	cards.bus = bus.GetVorpalBus()
	cards.bus.AddEngineListener(&cards)
	cards.tarotDeck = NewDeck()
	cards.evt = bus.NewDrawEvent()
	cards.evt.AddImageLayer(bus.NewImageLayer("samples/resources/tarot/table.png", 0, 0, 1900, 1400))
	//We are making event id do double duty here to track state as well as identify uniqueness. But its primary
	//purpose is for uniqueness and state should be tracked separately.
	cards.evt.SetId(0)
	cards.bus.SendDrawEvent(cards.evt)
	cards.bus.SendTextEvent(bus.NewTextEvent("Press S to shuffle the deck and show a card.\nPress N to show the next card.", 800, 200))

}

func (g *tarot) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			g.tarotDeck.Shuffle()

			g.bus.SendAudioEvent(bus.NewAudioEvent("samples/resources/audio/shuffle.mp3"))
			cards.evt = bus.NewDrawEvent()
			cards.evt.AddImageLayer(bus.NewImageLayer("samples/resources/tarot/table.png", 0, 0, 1900, 1400))
			cards.evt.SetId(0)
			cards.bus.SendTextEvent(bus.NewTextEvent("", 800, 200)) //Clear this way for now....need clear event..
			g.doSendCard()
		} else if evt.GetKey().EqualsIgnoreCase("N") {
			g.evt.SetId(g.evt.GetId() + 1)
			g.doSendCard()
			//TODO Check for max cards and end game...
		}

	}

}

func (g *tarot) doSendCard() {
	displayCard := g.tarotDeck.GetTopCard().GetCardImg()
	cardWidth := int32(200)
	cardHeight := int32(340)
	mainY := int32(360)
	mainX := int32(600)
	rightY := int32(710)
	rightX := int32(1100)

	switch g.evt.GetId() {

	case 0:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, mainX, mainY, cardWidth, cardHeight))
	case 1:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, mainX+50, mainY, cardWidth, cardHeight)) //TODO rotate..
	case 2:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, mainX-cardWidth-50, mainY, cardWidth, cardHeight))
	case 3:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, mainX+cardWidth+75, mainY, cardWidth, cardHeight))
	case 4:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, mainX, mainY+cardHeight+10, cardWidth, cardHeight))
	case 5:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, mainX, mainY-cardHeight-10, cardWidth, cardHeight))
	case 6:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, rightX, rightY, cardWidth, cardHeight))
	case 7:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, rightX, rightY-cardHeight-10, cardWidth, cardHeight))
	case 8:
		g.evt.AddImageLayer(bus.NewImageLayer(displayCard, rightX, rightY-(2*cardHeight)-20, cardWidth, cardHeight))
	}
	g.bus.SendDrawEvent(g.evt)
}
func (g *tarot) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		if evt.IsButtonToggled() {
			fmt.Println(evt)
		}
	}
}
