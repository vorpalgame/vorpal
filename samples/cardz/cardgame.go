package cardz

//TODO
import (
	"fmt"

	"github.com/vorpalgame/vorpal/bus"
)

type cardGame struct {
	bus  bus.VorpalBus
	deck Deck
}
type CardGame interface {
	Init()
}

var cards = cardGame{}

func GetCardGame() CardGame {
	cards.bus = bus.GetVorpalBus()
	cards.bus.AddKeyEventListener(&cards)
	cards.bus.AddMouseListener(&cards)
	return &cards
}

func (g *cardGame) Init() {
	fmt.Println("Card game init")

	g.deck = GetDeck()
	g.deck.Init()

}

func (g *cardGame) OnKeyEvent(keyChannel <-chan bus.KeyEvent) {
	for evt := range keyChannel {

		fmt.Println(evt.GetKey().ToString())

		if evt.GetKey().ToString() == "s" {
			g.deck.Shuffle()
		}

		//fmt.Println("Send card image name...")
		//	g.bus.SendDrawEvent(events.NewDrawEvent(g.deck.GetTopCard().GetCardImg()))

	}

}

func (g *cardGame) OnMouseEvent(mouseChannel <-chan bus.MouseEvent) {
	for evt := range mouseChannel {
		_ = evt
	}
}
