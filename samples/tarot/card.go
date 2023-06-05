package tarot

type tarotCard struct {
	cardImg string
}
type TarotCard interface {
	GetCardImg() string
}

func (c *tarotCard) GetCardImg() string {
	return c.cardImg
}
