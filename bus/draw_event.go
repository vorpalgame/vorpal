package bus

//Add coordinates, layers, etc. as necessary..

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

type DrawEvent interface {
	GetImageLayers() []ImageLayer
	AddImageLayer(imgLayer ImageLayer) DrawEvent
	Reset()
}

type drawEvent struct {
	imageLayers []ImageLayer
}

func NewDrawEvent() DrawEvent {
	evt := drawEvent{}
	evt.imageLayers = make([]ImageLayer, 0)
	return &evt
}

func (evt *drawEvent) Reset() {
	evt.imageLayers = make([]ImageLayer, 0)
}

func (evt *drawEvent) AddImageLayer(img ImageLayer) DrawEvent {
	evt.imageLayers = append(evt.imageLayers, img)
	return evt
}

func (evt *drawEvent) GetImageLayers() []ImageLayer {
	return evt.imageLayers
}

func NewImageLayer() ImageLayer {
	imageLayer := imageLayer{make([]ImageMetadata, 0)}
	return &imageLayer
}

func (i *imageLayer) Reset() {
	i.images = make([]ImageMetadata, 0)
}
func (i *imageLayer) GetLayerData() []ImageMetadata {
	return i.images
}

func (i *imageLayer) AddLayerData(img ImageMetadata) ImageLayer {
	i.images = append(i.images, img)
	return i
}

func NewImageMetadata(img string, x, y, scale int32) ImageMetadata {
	return &imageMetadata{img, x, y, scale, false}
}

type ImageMetadata interface {
	GetImage() string
	GetX() int32
	GetY() int32
	SetScale(scale int32) ImageMetadata
	GetScale() int32
	GetScalePercent() float32
	IsFlipHorizontal() bool
	SetFlipHorizontal(bool)
}

//Use builder pattern methods
type ImageLayer interface {
	GetLayerData() []ImageMetadata
	AddLayerData(imgMetadata ImageMetadata) ImageLayer
	Reset()
}

type imageMetadata struct {
	img            string
	x, y, scale    int32
	horizontalFlip bool
}
type imageLayer struct {
	images []ImageMetadata
}

func (e *imageMetadata) SetFlipHorizontal(horizontalFlip bool) {
	e.horizontalFlip = horizontalFlip
}
func (e *imageMetadata) IsFlipHorizontal() bool {
	return e.horizontalFlip
}
func (e *imageMetadata) GetImage() string {
	return e.img
}

func (p *imageMetadata) GetX() int32 {
	return p.x
}

func (p *imageMetadata) GetY() int32 {
	return p.y
}
func (p *imageMetadata) SetScale(scale int32) ImageMetadata {
	p.scale = scale
	return p
}
func (p *imageMetadata) GetScale() int32 {
	return p.scale
}

func (p *imageMetadata) GetScalePercent() float32 {
	return float32(p.scale) / 100
}
