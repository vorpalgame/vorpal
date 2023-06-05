package cardz

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vorpalgame/vorpal/bus"
)

// These should probably be private...
type deck struct {
	cards       []*card
	currentCard int
	vorpalBus   bus.VorpalBus
}

type Deck interface {
	Shuffle()
	GetTopCard() Card
	Init()
}

var suits = []string{"clubs", "spades", "hearts", "diamonds"}
var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "jack", "queen", "king", "ace"}

var d = deck{}

// TODO Revamp to use standard lifecycle
func GetDeck() Deck {
	d.currentCard = 0
	d.vorpalBus = bus.GetVorpalBus()

	return &d
}

func (d *deck) Init() {
	loadImagesEvent := bus.NewImagesEvent()
	for _, suit := range suits {
		for _, value := range values {

			cardSuit := "resources/" + value + "_of_" + suit + ".png"

			d.cards = append(d.cards, &card{value, suit, cardSuit})
			loadImagesEvent.AddImage(cardSuit)

		}

	}
	//fmt.Println(len(loadImagesEvent.GetImages()))

}

func (d *deck) GetTopCard() Card {
	if d.currentCard > len(d.cards) {
		d.currentCard = 0
	} else {
		d.currentCard += 1
	}
	return d.cards[d.currentCard]
}

func (d *deck) Shuffle() {
	fmt.Println("Shuffle up...")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}
