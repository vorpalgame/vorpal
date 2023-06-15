package tarot

import (
	"fmt"
	"log"
	"strings"

	"github.com/vorpalgame/vorpal/bus"
)

type tarot struct {
	bus       bus.VorpalBus
	tarotDeck TarotDeck
	drawEvent bus.DrawEvent
	textEvent bus.TextEvent
}

// TODO The cards should have locations or the game/board should dictate those...
// TODO The fonts and text layers need to be added to the text event.
var cards = tarot{}
var fontName = "samples/resources/fonts/Roboto-Regular.ttf"
var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

func InitGame() {
	log.Println("New card game")
	cards.bus = bus.GetVorpalBus()
	cards.bus.AddEngineListener(&cards)
	cards.tarotDeck = NewDeck()
	cards.drawEvent = bus.NewDrawEvent()

	cards.drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/tarot/table.png", 0, 0, 1920, 1080))
	//We are making event id do double duty here to track state as well as identify uniqueness. But its primary
	//purpose is for uniqueness and state should be tracked separately.
	//Need better coordination on the event numbers...
	cards.drawEvent.SetId(-2)
	cards.bus.SendDrawEvent(cards.drawEvent)
	cards.textEvent = bus.NewTextEvent(fontName, 18, 0, 0).AddText("Press S to shuffle and N to deal next card.").SetX(120).SetY(100)
	cards.bus.SendTextEvent(cards.textEvent)

}

func (g *tarot) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			g.tarotDeck.Shuffle()

			g.bus.SendAudioEvent(bus.NewAudioEvent("samples/resources/audio/shuffle.mp3"))
			cards.drawEvent = bus.NewDrawEvent()
			cards.drawEvent.AddImageLayer(bus.NewImageLayer("samples/resources/tarot/table.png", 0, 0, 1920, 1080))
			cards.drawEvent.SetId(0)
			cards.bus.SendTextEvent(cards.textEvent.Reinitialize().AddText(""))
			g.doSendCard()
		} else if evt.GetKey().EqualsIgnoreCase("N") && cards.drawEvent.GetId() >= 0 {
			g.drawEvent.SetId(g.drawEvent.GetId() + 1)
			g.doSendCard()
			//TODO Check for max cards and end game...
		}

	}

}

func (g *tarot) doSendCard() {
	card := g.tarotDeck.GetTopCard()
	//TODO iterate over it and split into lines.
	cards.textEvent.Reinitialize().SetX(120).SetY(100)
	g.formatCardText(card)
	cards.bus.SendTextEvent(cards.textEvent)
	displayCard := card.GetCardImg()
	cardWidth := int32(200)
	cardHeight := int32(340)
	mainY := int32(360)
	mainX := int32(1000)

	rightY := int32(mainY + 350)
	rightX := int32(mainX + 500)

	switch g.drawEvent.GetId() {

	case 0:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, mainX, mainY, cardWidth, cardHeight))
	case 1:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, mainX+50, mainY, cardWidth, cardHeight)) //TODO rotate..
	case 2:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, mainX-cardWidth-50, mainY, cardWidth, cardHeight))
	case 3:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, mainX+cardWidth+75, mainY, cardWidth, cardHeight))
	case 4:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, mainX, mainY+cardHeight+10, cardWidth, cardHeight))
	case 5:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, mainX, mainY-cardHeight-10, cardWidth, cardHeight))
	case 6:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, rightX, rightY, cardWidth, cardHeight))
	case 7:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, rightX, rightY-cardHeight-10, cardWidth, cardHeight))
	case 8:
		g.drawEvent.AddImageLayer(bus.NewImageLayer(displayCard, rightX, rightY-(2*cardHeight)-20, cardWidth, cardHeight))
	}
	g.bus.SendDrawEvent(g.drawEvent)
}

func (e *tarot) formatCardText(card TarotCard) {
	lineLength := 70

	sentences := strings.Split(card.GetCardText(), "\n")

	for lineNo, sentence := range sentences {
		sentence = strings.Trim(sentence, " ")
		if lineNo == 0 {
			cards.textEvent.AddTextLine(bus.NewTextLine(sentence, headerFontName, 24))
		} else {
			words := strings.Split(sentence, " ")
			var line string
			for _, word := range words {
				line += word + " "
				if len(line) > int(lineLength) {
					cards.textEvent.AddText(strings.Trim(line, " "))
					line = ""
				}

			}
			cards.textEvent.AddText(strings.Trim(line, " "))
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
