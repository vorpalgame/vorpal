package main

import (
	"fmt"
	"math/rand"
	"time"
)

// These should probably be private...
type TarotDeckData struct {
	currentCard int             `yaml:"-"`
	TarotDeck   []TarotCardData `yaml:"TarotDeck"`
}

type TarotDeck interface {
	Shuffle()
	GetTopCard() TarotCard
}

func (d *TarotDeckData) GetTopCard() TarotCard {

	card := &d.TarotDeck[d.currentCard]
	d.currentCard += 1
	if d.currentCard >= len(d.TarotDeck) {
		d.currentCard = 0
	}
	return card
}

func (d *TarotDeckData) Shuffle() {
	fmt.Println("Shuffle up...")
	rand.Seed(time.Now().UnixNano()) //TODO fix to remove deperecated...
	rand.Shuffle(len(d.TarotDeck), func(i, j int) { d.TarotDeck[i], d.TarotDeck[j] = d.TarotDeck[j], d.TarotDeck[i] })
	d.currentCard = 0
}
