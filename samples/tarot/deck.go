package tarot

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

// These should probably be private...
type TarotDeckData struct {
	CurrentCard int             `yaml:"CurrentCard"`
	TarotDeck   []TarotCardData `yaml:"TarotDeck"`
	bus         bus.VorpalBus   `yaml:"-"`
}

type TarotDeck interface {
	Shuffle()
	GetTopCard() TarotCard
}

func (d *TarotDeckData) GetTopCard() TarotCard {
	card := &d.TarotDeck[d.CurrentCard]
	d.CurrentCard += 1
	if d.CurrentCard >= len(d.TarotDeck) {
		d.CurrentCard = 0
	}
	return card
}

func (d *TarotDeckData) Shuffle() {
	fmt.Println("Shuffle up...")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.TarotDeck), func(i, j int) { d.TarotDeck[i], d.TarotDeck[j] = d.TarotDeck[j], d.TarotDeck[i] })
	d.CurrentCard = 0
}
