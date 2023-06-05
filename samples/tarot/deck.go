package tarot

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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
	Init()
}

var d = tarotDeck{}

func NewDeck() TarotDeck {
	d.currentCard = 0
	d.bus = bus.GetVorpalBus()

	return &d
}

func (d *tarotDeck) Init() {
	log.Print("Load tarot images...")
	files, err := ioutil.ReadDir("samples/resources/tarot")
	if err != nil {
		log.Fatal(err)
	}
	loadImagesEvent := bus.NewImagesEvent()
	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())

		resource := "samples/resources/tarot/" + file.Name()
		d.cards = append(d.cards, &tarotCard{resource})
		loadImagesEvent.AddImage(resource)

	}

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
