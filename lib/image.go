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

//TODO We should refactor lib to make it general use the Point

type ImageMetadata interface {
	GetImage() string
	GetX() int32
	GetY() int32
	SetScale(scale int32) ImageMetadata
	GetScale() int32
	GetScalePercent() float32
	IsFlipHorizontal() bool
	SetFlipHorizontal(bool) ImageMetadata
	SetX(x int32) ImageMetadata
	SetY(y int32) ImageMetadata
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
func (e *imageMetadata) GetImage() string {
	return e.img
}

func (p *imageMetadata) GetX() int32 {
	return p.x
}

//TODO Refactor to use Point

func (p *imageMetadata) SetX(x int32) ImageMetadata {
	p.x = x
	return p
}
func (p *imageMetadata) GetY() int32 {
	return p.y
}
func (p *imageMetadata) SetY(y int32) ImageMetadata {
	p.y = y
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
