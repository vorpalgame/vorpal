package bus

import "github.com/vorpalgame/vorpal/lib"

type TextEventListener interface {
	OnTextEvent(textChannel <-chan TextEvent)
}

//TODO Builder pattern isn't quite right for this yet. Refactor to correct.
//TODO Set/Get Location and Reinitialize should be implemented
//in a generic way to return the actual event. For example,
//MultilineTextEvent is a TextEvent but it should return that
//interface and not TextEvent only. Otherwis the builder pattern
//is broken.
//////////////////////////////////////////////////
//// TextEvent
//////////////////////////////////////////////////
type TextEvent interface {
	Reinitialize() TextEvent
	GetLocation() lib.Point
	SetLocation(x, y int32) TextEvent
}
type textEventData struct {
	lib.Point
}

func (ted *textEventData) GetLocation() lib.Point {
	return ted.Point
}
func (ted *textEventData) SetLocation(x, y int32) TextEvent {
	ted.Point = lib.NewPoint(x, y)
	return ted
}

func (ted *textEventData) Reinitialize() TextEvent {
	//More specific events can override in whatever way
	//make sense. At this level there isn't much to do
	//nil'ing the location is potentially dangerous.
	return ted
}

/////////////////////////////////////////////////////
//// Font
/////////////////////////////////////////////////////
func NewFont(font string, size int32) Font {
	return &fontData{font, size}
}

//Font probably belongs in lib
type Font interface {
	GetFont() string
	GetFontSize() int32
	SetFont(font string) Font
	SetFontSize(int32) Font
}

///// Font
//TODO Font probably belongs in lib...
type fontData struct {
	font     string
	fontSize int32
}

func (f *fontData) SetFont(font string) Font {
	f.font = font
	return f
}

func (f *fontData) SetFontSize(size int32) Font {
	f.fontSize = size
	return f
}

func (f *fontData) GetFont() string {
	return f.font
}

func (f *fontData) GetFontSize() int32 {
	return f.fontSize
}

//////////////////////////////////////////////////
//// MultilineTextEvent
//////////////////////////////////////////////////
type MultilineTextEvent interface {
	TextEvent
	Font
	GetText() []TextLine
	AddTextLine(TextLine) MultilineTextEvent
	AddText(string) MultilineTextEvent
	GetId() int32
}

type TextLine interface {
	Font
	GetText() string
}

//////
// Text event can have a font specified for defaults.
//TODO switch to using Point for location.
type multilineTextEventData struct {
	textEventData
	Font
	text []TextLine
	id   int32
}

var nextTextEventId = int32(0)

func NewMultilineTextEvent(font string, fontSize, x, y int32) MultilineTextEvent {
	nextTextEventId++ //Oddly can't do this in the struct
	return &multilineTextEventData{textEventData{lib.NewPoint(x, y)}, &fontData{font, fontSize}, make([]TextLine, 0), nextTextEventId}

}

// Reuse the text event with font/size information but clear the slice and update the id
//This is problematic for the builder pattern.
func (e *multilineTextEventData) Reinitialize() TextEvent {
	e.text = make([]TextLine, 0)
	e.id = e.id + 1
	return e
}

// Use default font and font size.
func (e *multilineTextEventData) AddText(text string) MultilineTextEvent {
	e.text = append(e.text, &textLine{text, e.Font})
	return e
}

// If one wishes to specify diferent font or font size.
func (e *multilineTextEventData) AddTextLine(text TextLine) MultilineTextEvent {
	e.text = append(e.text, text)
	return e
}
func (e *multilineTextEventData) GetText() []TextLine {
	return e.text
}

func (p *multilineTextEventData) GetId() int32 {
	return p.id
}

////// TextLine
func NewTextLine(text string, font Font) TextLine {
	return &textLine{text, font}
}

type textLine struct {
	text string
	Font
}

func (p *textLine) GetText() string {
	return p.text
}
