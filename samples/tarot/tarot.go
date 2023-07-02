package tarot

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

type TarotGame interface {
}

// TODO This should match with the tarot_bootstrap.yaml so
// that it can be unmarshaled
type Tarot struct {
	Title           string        `yaml:"Title"`
	HeaderFont      string        `yaml:"HeaderFont"`
	TextFont        string        `yaml:"TextFont"`
	ShuffleAudio    string        `yaml:"ShuffleAudio"`
	BackgroundImage string        `yaml:"BackgroundImage"`
	TarotDeck       TarotDeckData `yaml:"TarotDeck"`
	currentCard     int32
	shuffled        bool
	bus             bus.VorpalBus
	drawEvent       bus.DrawEvent
	textEvent       bus.TextEvent
}

func NewGame() TarotGame {

	currentGame := Tarot{}
	viper.Unmarshal(&currentGame)
	currentGame.bus = bus.GetVorpalBus()

	currentGame.bus.AddMouseListener(&currentGame)
	currentGame.bus.AddKeyEventListener(&currentGame)

	currentGame.shuffled = false
	currentGame.doStartupScreen()
	return currentGame

}

func (t *Tarot) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			t.bus.SendAudioEvent(bus.NewAudioEvent(t.ShuffleAudio).Stop())
			t.drawEvent.AddImageLayer(bus.NewImageLayer().AddLayerData(bus.NewImageMetadata(t.BackgroundImage, 0, 0, 100)))
			t.TarotDeck.Shuffle()
			t.shuffled = true
			t.bus.SendAudioEvent(bus.NewAudioEvent(t.ShuffleAudio).Play())
			t.currentCard = 0
			t.doSendCard()
		} else if evt.GetKey().EqualsIgnoreCase("N") && t.shuffled {
			t.doSendCard()
		}

	}

}

func (t *Tarot) doStartupScreen() {
	t.currentCard = 0
	t.shuffled = false
	//The keys we are interested in.
	//These shoudl be registered from the yaml configuration
	//TODO Get key bindings from yaml file...
	t.bus.SendControlEvent(bus.NewWindowTitleEvent(t.Title))
	t.bus.SendKeysRegistrationEvent(bus.NewKeysRegistrationEvent(lib.NewKeys([]string{"s", "n", "S", "N"})))
	t.drawEvent = bus.NewDrawEvent()

	t.drawEvent.AddImageLayer(bus.NewImageLayer().AddLayerData(bus.NewImageMetadata(t.BackgroundImage, 0, 0, 100)))

	t.bus.SendDrawEvent(t.drawEvent)
	//Get intro text from Yaml file.
	t.textEvent = bus.NewTextEvent(t.TextFont, 18, 0, 0).AddText("Press S to shuffle and N to deal next card.").SetX(120).SetY(100)
	t.bus.SendTextEvent(t.textEvent)

}

func (t *Tarot) doSendCard() {
	card := t.TarotDeck.GetTopCard()
	//TODO iterate over it and split into lines.
	t.textEvent.Reinitialize().SetX(120).SetY(100)
	t.formatCardText(card)
	t.bus.SendTextEvent(t.textEvent)
	displayCard := card.GetCardImg()
	actualWidth := float64(1500)
	actualHeight := float64(2550)
	scale := float64(0.1)

	centralY := int32(360)
	centralX := int32(900)
	cardWidth := int32(actualWidth * scale)
	cardHeight := int32(actualHeight * scale)
	rightY := int32(810)
	rightX := int32(1300)

	switch t.currentCard {

	case 0:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, centralX, centralY))
	case 1:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, centralX+50, centralY)) //TODO rotate..
	case 2:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, centralX-cardWidth-25, centralY))
	case 3:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, centralX+cardWidth+75, centralY))
	case 4:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, centralX, centralY+cardHeight+10))
	case 5:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, centralX, centralY-cardHeight-10))
	case 6:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY))
	case 7:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY-cardHeight-10))
	case 8:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY-(2*cardHeight)-20))
	case 9:
		t.drawEvent.AddImageLayer(createImageLayer(displayCard, rightX, rightY-(3*cardHeight)-30))
	}
	t.bus.SendDrawEvent(t.drawEvent)

	if t.currentCard < 10 {
		t.currentCard++
	} else {
		t.doStartupScreen()
	}
}

// TODO Current seetting this to 25% hard coded...
func createImageLayer(displayCard string, mainX, mainY int32) bus.ImageLayer {
	imgData := bus.NewImageMetadata(displayCard, mainX, mainY, 10)
	return bus.NewImageLayer().AddLayerData(imgData)
}

// TODO Put formatting and font in Yaml
// TODO Text and Lines shyould be in the lib and not
// bus package.
func (t *Tarot) formatCardText(card TarotCard) {
	lineLength := 70
	t.textEvent.AddTextLine(bus.NewTextLine(card.GetCardTitle(), t.HeaderFont, 24))
	sentences := strings.Split(card.GetCardText(), "\n")

	for _, sentence := range sentences {
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
func (t *Tarot) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		if evt.LeftButton().IsDown() {
			//fmt.Println(evt)
		}
	}
}
