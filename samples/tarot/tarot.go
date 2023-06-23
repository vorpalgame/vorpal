package tarot

import (
	"fmt"
	"strings"

	"github.com/vorpalgame/vorpal/bus"
)

type TarotGame interface {
}

type tarot struct {
	bus         bus.VorpalBus
	tarotDeck   TarotDeck
	drawEvent   bus.DrawEvent
	textEvent   bus.TextEvent
	currentCard int32
}

// TODO The cards should have locations or the game/board should dictate those...
// TODO The fonts and text layers need to be added to the text event.

var fontName = "samples/resources/fonts/Roboto-Regular.ttf"
var headerFontName = "samples/resources/fonts/Roboto-Black.ttf"

func NewGame() TarotGame {

	var t = tarot{}
	t.bus = bus.GetVorpalBus()
	t.bus.AddMouseListener(&t)
	t.bus.AddKeyEventListener(&t)
	t.tarotDeck = NewDeck()
	t.doStart()
	return t

}

func (t *tarot) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			t.tarotDeck.Shuffle()

			t.bus.SendAudioEvent(bus.NewAudioEvent("samples/resources/audio/shuffle.mp3").Play())
			t.drawEvent = bus.NewDrawEvent()
			t.drawEvent.AddImageLayer(bus.NewImageLayer().AddLayerData(bus.NewImageMetadata("samples/resources/tarot/table.png", 0, 0, 1920, 1080)))

			t.bus.SendTextEvent(t.textEvent.Reinitialize().AddText(""))
			t.doSendCard()
		} else if evt.GetKey().EqualsIgnoreCase("N") {
			t.doSendCard()
			//TODO Check for max cards and end game...
		}

	}

}

func (t *tarot) doStart() {
	t.currentCard = 0
	//The keys we are interested in.
	t.bus.SendKeysRegistrationEvent(bus.NewKeysRegistrationEvent("s", "n"))
	t.drawEvent = bus.NewDrawEvent()

	t.drawEvent.AddImageLayer(bus.NewImageLayer().AddLayerData(bus.NewImageMetadata("samples/resources/tarot/table.png", 0, 0, 1920, 1080)))

	t.bus.SendDrawEvent(t.drawEvent)
	t.textEvent = bus.NewTextEvent(fontName, 18, 0, 0).AddText("Press S to shuffle and N to deal next card.").SetX(120).SetY(100)
	t.bus.SendTextEvent(t.textEvent)

}

func (t *tarot) doSendCard() {
	card := t.tarotDeck.GetTopCard()
	//TODO iterate over it and split into lines.
	t.textEvent.Reinitialize().SetX(120).SetY(100)
	t.formatCardText(card)
	t.bus.SendTextEvent(t.textEvent)
	displayCard := card.GetCardImg()
	cardWidth := int32(200)
	cardHeight := int32(340)
	mainY := int32(360)
	mainX := int32(1000)

	rightY := int32(mainY + 350)
	rightX := int32(mainX + 500)

	switch t.currentCard {

	case 0:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, mainX, mainY, cardWidth, cardHeight))
	case 1:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, mainX+50, mainY, cardWidth, cardHeight)) //TODO rotate..
	case 2:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, mainX-cardWidth-50, mainY, cardWidth, cardHeight))
	case 3:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, mainX+cardWidth+75, mainY, cardWidth, cardHeight))
	case 4:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, mainX, mainY+cardHeight+10, cardWidth, cardHeight))
	case 5:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, mainX, mainY-cardHeight-10, cardWidth, cardHeight))
	case 6:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY, cardWidth, cardHeight))
	case 7:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY-cardHeight-10, cardWidth, cardHeight))
	case 8:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY-(2*cardHeight)-20, cardWidth, cardHeight))
	}
	t.bus.SendDrawEvent(t.drawEvent)

	if t.currentCard < 9 {
		t.currentCard++

	} else {
		t.doStart()
	}
}
func createImageLayer(displayCard string, mainX, mainY, cardWidth, cardHeight int32) bus.ImageLayer {
	imgData := bus.NewImageMetadata(displayCard, mainX, mainY, cardWidth, cardHeight)
	return bus.NewImageLayer().AddLayerData(imgData)
}
func (t *tarot) formatCardText(card TarotCard) {
	lineLength := 70

	sentences := strings.Split(card.GetCardText(), "\n")

	for lineNo, sentence := range sentences {
		sentence = strings.Trim(sentence, " ")
		if lineNo == 0 {
			t.textEvent.AddTextLine(bus.NewTextLine(sentence, headerFontName, 24))
		} else {
			words := strings.Split(sentence, " ")
			var line string
			for _, word := range words {
				line += word + " "
				if len(line) > int(lineLength) {
					t.textEvent.AddText(strings.Trim(line, " "))
					line = ""
				}

			}
			t.textEvent.AddText(strings.Trim(line, " "))
		}

	}
}
func (t *tarot) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		if evt.LeftButton().IsDown() {
			fmt.Println(evt)
		}
	}
}
