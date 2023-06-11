package bus

//Add coordinates, layers, etc. as necessary..

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

//TODO review to ensure use of pointers/addresses and not end up with copies....
type DrawEvent interface {
	GetImageLayers() []ImageLayer
	AddImageLayer(imgLayer ImageLayer)
	GetId() int32
}

type drawEvent struct {
	imageLayers []ImageLayer
	id          int32
}

//TODO This should use a UUID but that's external
var id int32 = 0

func NewDrawEvent() DrawEvent {
	evt := drawEvent{}
	evt.imageLayers = make([]ImageLayer, 0, 100)
	id = id + 1
	evt.id = id
	return &evt
}

func (evt *drawEvent) GetId() int32 {
	return evt.id
}

func (evt *drawEvent) AddImageLayer(img ImageLayer) {
	evt.imageLayers = append(evt.imageLayers, img)
}

func (evt *drawEvent) GetImageLayers() []ImageLayer {
	return evt.imageLayers
}

func NewImageLayer(img string, x, y, width, height int32) ImageLayer {
	return &imageLayer{img, x, y, width, height}
}

type ImageLayer interface {
	GetImage() string
	GetX() int32
	GetY() int32
	GetHeight() int32
	GetWidth() int32
}

type imageLayer struct {
	img                 string
	x, y, width, height int32
}

func (e *imageLayer) GetImage() string {
	return e.img
}

func (p *imageLayer) GetX() int32 {
	return p.x
}

func (p *imageLayer) GetY() int32 {
	return p.y
}

func (p *imageLayer) GetWidth() int32 {
	return p.width
}

func (p *imageLayer) GetHeight() int32 {
	return p.height
}
