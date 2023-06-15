package tarot

type tarotCard struct {
	cardImg  string
	cardText string
}
type TarotCard interface {
	GetCardImg() string
	GetCardText() string
}

func (c *tarotCard) GetCardImg() string {
	return c.cardImg
}

func (c *tarotCard) GetCardText() string {
	return c.cardText
}
