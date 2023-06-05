package cardz

type card struct {
	value   string
	suit    string
	cardImg string
}
type Card interface {
	GetValue() string
	GetSuit() string
	GetCardImg() string
}

func (c *card) GetSuit() string {
	return c.suit
}

func (c *card) GetValue() string {
	return c.suit
}

func (c *card) GetCardImg() string {
	return c.cardImg
}
