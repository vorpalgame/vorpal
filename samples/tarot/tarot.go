package tarot

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

// TODO Revisit all the sizeing as we've switched from scaled to absolute.
// Scaling has to take place on game side.
type TarotGame interface {
}

//We use the struct embededded when we need to marshal/unmarshal or
//Yaml/Viper get confused. Not a serialization technology.

type Tarot struct {
	Title           string             `yaml:"Title"`
	HeaderFont      string             `yaml:"HeaderFont"`
	TextFont        string             `yaml:"TextFont"`
	ShuffleAudio    lib.AudioStateData `yaml:"ShuffleAudio"`
	BackgroundImage string             `yaml:"BackgroundImage"`
	TarotDeck       TarotDeckData      `yaml:"TarotDeck"`
	currentCard     int32
	shuffled        bool
	bus             bus.VorpalBus
	drawEvent       bus.DrawLayersEvent
	textEvent       bus.MultilineTextEvent
}

func NewGame() TarotGame {

	currentGame := Tarot{}
	//TODO Check marshaling and bail out when bad
	e := viper.Unmarshal(&currentGame)
	if e != nil {
		panic(e)
	}
	currentGame.bus = bus.GetVorpalBus()

	currentGame.bus.AddMouseListener(&currentGame)
	currentGame.bus.AddKeyEventListener(&currentGame)

	currentGame.shuffled = false
	currentGame.doStartupScreen()
	return currentGame

}

// TODO We are currently constructing and passing with structs due to issue disecvoered
// with the marshaling. The interfaces and constructors shoudl be wired back in now that those
// are straightened out.
func (t *Tarot) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		if evt.GetKey().EqualsIgnoreCase("S") {
			//t.ShuffleAudio).Stop()
			t.bus.SendAudioEvent(bus.NewStopAudioEvent(&t.ShuffleAudio))
			layer := lib.ImageLayerData{}
			metatdata := lib.ImageMetadata{t.BackgroundImage, 0, 0, 1920, 1080, false}
			layer.LayerMetadata = append(layer.LayerMetadata, &metatdata)
			t.drawEvent.AddImageLayer(layer)
			t.TarotDeck.Shuffle()
			t.shuffled = true
			t.bus.SendAudioEvent(bus.NewPlayAudioEvent(&t.ShuffleAudio))
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
	//These should be registered from the yaml configuration
	//TODO Get key bindings from yaml file...
	t.bus.SendControlEvent(bus.NewWindowTitleEvent(t.Title))
	t.bus.SendKeysRegistrationEvent(bus.NewKeysRegistrationEvent(lib.NewKeys([]string{"s", "n", "S", "N"})))
	t.drawEvent = bus.NewDrawLayersEvent()

	layer := lib.ImageLayerData{}
	metatdata := lib.ImageMetadata{t.BackgroundImage, 0, 0, 1920, 1080, false}
	layer.LayerMetadata = append(layer.LayerMetadata, &metatdata)
	t.drawEvent.AddImageLayer(layer)
	t.bus.SendDrawEvent(t.drawEvent)
	//Get intro text from Yaml file.
	t.textEvent = bus.NewMultilineTextEvent(t.TextFont, 18, 0, 0).AddText("Press S to shuffle and N to deal next card.")
	t.textEvent.SetLocation(120, 100)
	t.bus.SendTextEvent(t.textEvent)

}

func (t *Tarot) doSendCard() {
	card := t.TarotDeck.GetTopCard()
	//TODO iterate over it and split into lines.
	t.textEvent.Reinitialize().SetLocation(120, 100)
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

func createImageLayer(displayCard string, mainX, mainY int32) lib.ImageLayerData {
	imgData := lib.ImageMetadata{displayCard, mainX, mainY, 155, 255, false}
	layer := lib.ImageLayerData{}
	layer.LayerMetadata = append(layer.LayerMetadata, &imgData)
	return layer
}

// TODO Put formatting and font in Yaml
// TODO Text and Lines shyould be in the lib and not
// bus package.
func (t *Tarot) formatCardText(card TarotCard) {
	lineLength := 70
	t.textEvent.AddTextLine(lib.NewTextLine(card.GetCardTitle(), lib.NewFont(t.HeaderFont, 24)))
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
