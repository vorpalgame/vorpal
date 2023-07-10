package lib

func NewImageLayer() ImageLayer {
	imageLayer := imageLayer{make([]ImageMetadata, 0)}
	return &imageLayer
}

//Use builder pattern methods

type ImageLayer interface {
	GetLayerData() []ImageMetadata
	AddLayerData(imgMetadata ImageMetadata) ImageLayer
	Reset()
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

func NewImageMetadata(img string, x, y, width, height int32) ImageMetadata {
	return &imageMetadata{img, x, y, width, height, false}
}

//TODO diferent types for scale versus fixed size rectangle.

type ImageMetadata interface {
	GetFileName() string
	GetRectangle() (x, y, width, height int32)
	SetRectangle(x, y, width, height int32) ImageMetadata
	IsFlipHorizontal() bool
	SetFlipHorizontal(bool) ImageMetadata
}

type imageMetadata struct {
	img                 string
	x, y, width, height int32
	horizontalFlip      bool
}
type imageLayer struct {
	images []ImageMetadata
}

func (e *imageMetadata) SetFlipHorizontal(horizontalFlip bool) ImageMetadata {
	e.horizontalFlip = horizontalFlip
	return e
}
func (e *imageMetadata) IsFlipHorizontal() bool {
	return e.horizontalFlip
}
func (e *imageMetadata) GetFileName() string {
	return e.img
}

func (p *imageMetadata) GetRectangle() (x, y, width, height int32) {
	return p.x, p.y, p.width, p.height
}

//TODO Refactor to use Point

func (p *imageMetadata) SetRectangle(x, y, width, height int32) ImageMetadata {
	p.x, p.y, p.width, p.height = x, y, width, height
	return p
}
