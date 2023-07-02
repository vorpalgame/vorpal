package lib

/////////////////////////////////////////////////////
//// Font
/////////////////////////////////////////////////////

func NewFont(font string, size int32) Font {
	return &fontData{font, size}
}

type Font interface {
	GetFont() string
	GetFontSize() int32
	SetFont(font string) Font
	SetFontSize(int32) Font
}

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

/////////////////////////////////////////////////////
////// Text Line
////////////////////////////////////////////////////

type TextLine interface {
	Font
	GetText() string
}

// //// TextLine

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
