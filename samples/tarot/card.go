package tarot

type TarotCardData struct {
	CardImg  string `yaml:"CardImg"`
	Title    string `yaml:"Title"`
	CardText string `yaml:"CardText"`
}
type TarotCard interface {
	GetCardImg() string
	GetCardText() string
	GetCardTitle() string
}

func (c *TarotCardData) GetCardImg() string {
	return c.CardImg
}

func (c *TarotCardData) GetCardText() string {
	return c.CardText
}

func (c *TarotCardData) GetCardTitle() string {
	return c.Title
}
