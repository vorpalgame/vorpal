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

func NewImageMetadata(img string, x, y, scale int32) ImageMetadata {
	return &imageMetadata{img, x, y, scale, false}
}

//TODO diferent types for scale versus fixed size rectangle.

type ImageMetadata interface {
	GetFileName() string
	GetPoint() (x, y int32)
	SetScale(scale int32) ImageMetadata
	GetScale() int32
	GetScalePercent() float32
	IsFlipHorizontal() bool
	SetFlipHorizontal(bool) ImageMetadata
	SetPoint(x, y int32) ImageMetadata
}

type imageMetadata struct {
	img            string
	x, y, scale    int32
	horizontalFlip bool
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

func (p *imageMetadata) GetPoint() (x, y int32) {
	return p.x, p.y
}

//TODO Refactor to use Point

func (p *imageMetadata) SetPoint(x, y int32) ImageMetadata {
	p.x, p.y = x, y
	return p
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
