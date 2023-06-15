package bus

type TextEventListener interface {
	OnTextEvent(textChannel <-chan TextEvent)
}

// This is currently being refactored from previous verious that had a single line of
// text and font/size to one that is slices of text lines which may override font/size
// TODO other elements like location and type size.
// All text in the event are assumed to be of the same font and size unless
// they are overridden for specific lines.
type TextEvent interface {
	GetFont() string
	GetFontSize() int32
	GetText() []TextLine
	AddTextLine(TextLine) TextEvent
	AddText(string) TextEvent
	GetId() int32
	GetX() int32
	GetY() int32
	SetX(int32) TextEvent
	SetY(int32) TextEvent
	Reinitialize() TextEvent
}

type TextLine interface {
	GetText() string
	GetFont() string
	GetFontSize() int32
}

// Text event can have a font specified for defaults.
type textEvent struct {
	text               []TextLine
	font               string
	fontSize, x, y, id int32
}

var nextTextEventId = int32(0)

func NewTextEvent(font string, fontSize, x, y int32) TextEvent {
	nextTextEventId++
	return &textEvent{make([]TextLine, 0), font, fontSize, x, y, nextTextEventId}

}

func NewTextLine(text, font string, fontSize int32) TextLine {
	return &textLine{text, font, fontSize}
}

// Reuse the text event with font/size information but clear the slice and update the id
func (e *textEvent) Reinitialize() TextEvent {
	e.text = make([]TextLine, 0)
	e.id = e.id + 1
	return e
}
func (p *textEvent) GetFontSize() int32 {
	return p.fontSize
}
func (p *textEvent) GetFont() string {
	return p.font
}

// Use default font and font size.
func (e *textEvent) AddText(text string) TextEvent {
	e.text = append(e.text, &textLine{text, e.font, e.fontSize})
	return e
}

// If one wishes to specify diferent font or font size.
func (e *textEvent) AddTextLine(text TextLine) TextEvent {
	e.text = append(e.text, text)
	return e
}
func (e *textEvent) GetText() []TextLine {
	return e.text
}
func (p *textEvent) GetX() int32 {
	return p.x
}
func (p *textEvent) GetY() int32 {
	return p.y
}
func (p *textEvent) SetX(x int32) TextEvent {
	p.x = x
	return p
}
func (p *textEvent) SetY(y int32) TextEvent {
	p.y = y
	return p
}
func (p *textEvent) GetId() int32 {
	return p.id
}

type textLine struct {
	text     string
	font     string
	fontSize int32
}

func (p *textLine) GetText() string {
	return p.text
}
func (p *textLine) GetFontSize() int32 {
	return p.fontSize
}
func (p *textLine) GetFont() string {
	return p.font
}
