package bus

import "github.com/vorpalgame/vorpal/lib"

type TextEventListener interface {
	OnTextEvent(textChannel <-chan TextEvent)
}

//TODO Builder pattern isn't quite right for this yet. Refactor to correct.
//TODO Set/Get Location and Reinitialize should be implemented
//in a generic way to return the actual event. For example,
//MultilineTextEvent is a TextEvent but it should return that
//interface and not TextEvent only. Otherwise the builder pattern
//is broken.

// ////////////////////////////////////////////////
// // TextEvent
// ////////////////////////////////////////////////

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

//////////////////////////////////////////////////
//// MultilineTextEvent
//////////////////////////////////////////////////

type MultilineTextEvent interface {
	TextEvent
	lib.Font
	GetText() []lib.TextLine
	AddTextLine(lib.TextLine) MultilineTextEvent
	AddText(string) MultilineTextEvent
	GetId() int32
}

// ////
// Text event can have a font specified for defaults.
// TODO switch to using Point for location.
type multilineTextEventData struct {
	textEventData
	lib.Font
	text []lib.TextLine
	id   int32
}

var nextTextEventId = int32(0)

func NewMultilineTextEvent(font string, fontSize, x, y int32) MultilineTextEvent {
	nextTextEventId++ //Oddly can't do this in the struct
	return &multilineTextEventData{textEventData{lib.NewPoint(x, y)}, lib.NewFont(font, fontSize), make([]lib.TextLine, 0), nextTextEventId}

}

// Reuse the text event with font/size information but clear the slice and update the id
// This is problematic for the builder pattern.

func (e *multilineTextEventData) Reinitialize() TextEvent {
	e.text = make([]lib.TextLine, 0)
	e.id = e.id + 1
	return e
}

func (e *multilineTextEventData) AddText(text string) MultilineTextEvent {
	e.text = append(e.text, lib.NewTextLine(text, e.Font))
	return e
}

// If one wishes to specify diferent font or font size.

func (e *multilineTextEventData) AddTextLine(text lib.TextLine) MultilineTextEvent {
	e.text = append(e.text, text)
	return e
}
func (e *multilineTextEventData) GetText() []lib.TextLine {
	return e.text
}

func (p *multilineTextEventData) GetId() int32 {
	return p.id
}
