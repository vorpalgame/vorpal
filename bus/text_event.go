package bus

type TextEventListener interface {
	OnTextEvent(textChannel <-chan TextEvent)
}

//TODO other elements like location and type size.

type TextEvent interface {
	GetText() string
	GetX() int32
	GetY() int32
}

type textEvent struct {
	text string
	x, y int32
}

func NewTextEvent(text string, x, y int32) TextEvent {
	return &textEvent{text, x, y}

}

func (e *textEvent) GetText() string {
	return e.text
}

func (p *textEvent) GetX() int32 {
	return p.x
}

func (p *textEvent) GetY() int32 {
	return p.y
}
