package tarot

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

// These should probably be private...
type tarotDeck struct {
	cards       []*tarotCard
	currentCard int
	bus         bus.VorpalBus
}

type TarotDeck interface {
	Shuffle()
	GetTopCard() TarotCard
}

var d = tarotDeck{}

func NewDeck() TarotDeck {
	d.currentCard = 0
	d.bus = bus.GetVorpalBus()

	dir := "samples/resources/tarot/"
	for i := 0; i < 40; i++ {
		base := dir + strconv.Itoa(i)
		content, _ := os.ReadFile(base + ".txt")
		d.cards = append(d.cards, &tarotCard{base + ".png", d.formatText(string(content))})
	}

	return &d
}

func (d *tarotDeck) formatText(text string) string {
	//TODO shoudl we strip \n as well?
	text = strings.ReplaceAll(text, "\r", "")
	// return strings.Join(sentences, ".\n")
	return text
}
func (d *tarotDeck) GetTopCard() TarotCard {
	if d.currentCard >= len(d.cards) {
		d.currentCard = 0
	} else {
		d.currentCard += 1
	}
	return d.cards[d.currentCard]
}

func (d *tarotDeck) Shuffle() {
	fmt.Println("Shuffle up...")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
	d.currentCard = 0
}
