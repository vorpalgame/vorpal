package bus

//Add coordinates, layers, etc. as necessary..

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

type DrawEvent interface {
	GetImage() string
	GetX() int32
	GetY() int32
	GetZ() int32
	GetHeight() int32
	GetWidth() int32
}

type drawEvent struct {
	img                    string
	x, y, z, width, height int32
}

//TODO is there a Golang Rectangle???
func NewDrawEvent(img string, x, y, z, width, height int32) DrawEvent {
	return &drawEvent{img, x, y, z, width, height}

}

func (e *drawEvent) GetImage() string {
	return e.img
}

func (p *drawEvent) GetX() int32 {
	return p.x
}

func (p *drawEvent) GetY() int32 {
	return p.y
}

func (p *drawEvent) GetZ() int32 {
	return p.z
}

func (p *drawEvent) GetWidth() int32 {
	return p.width
}

func (p *drawEvent) GetHeight() int32 {
	return p.height
}
