package bus

//Add coordinates, layers, etc. as necessary..

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

type DrawEvent interface {
	GetImageLayers() []ImageLayer
	AddImageLayer(imgLayer ImageLayer)
	Reset()
}

type drawEvent struct {
	imageLayers []ImageLayer
}

func NewDrawEvent() DrawEvent {
	evt := drawEvent{}
	evt.imageLayers = make([]ImageLayer, 0, 100)
	return &evt
}

func (evt *drawEvent) Reset() {
	evt.imageLayers = make([]ImageLayer, 0, 100)
}

func (evt *drawEvent) AddImageLayer(img ImageLayer) {
	evt.imageLayers = append(evt.imageLayers, img)
}

func (evt *drawEvent) GetImageLayers() []ImageLayer {
	return evt.imageLayers
}

func NewImageLayer(img string, x, y, width, height int32) ImageLayer {
	return &imageLayer{img, x, y, width, height, false}
}

type ImageLayer interface {
	GetImage() string
	GetX() int32
	GetY() int32
	GetHeight() int32
	GetWidth() int32
	IsFlipHorizontal() bool
	SetFlipHorizontal(bool)
}

type imageLayer struct {
	img                 string
	x, y, width, height int32
	horizontalFlip      bool
}

func (e *imageLayer) SetFlipHorizontal(horizontalFlip bool) {
	e.horizontalFlip = horizontalFlip
}
func (e *imageLayer) IsFlipHorizontal() bool {
	return e.horizontalFlip
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
