package lib

// TODO We are moving from using structs and pointers to structs to interfaces.
// Early interfaces were abandoned due to marshaling issues. Structs were separated
// into data and tested.
type ImageDescriptor interface {
	GetImageFileName() string
	GetX() int32
	GetY() int32
	GetWidth() int32
	GetHeight() int32
	GetHorizontalFlip() bool
}

func (i *ImageMetadata) GetImageFileName() string {
	return i.ImageFileName
}

func (i *ImageMetadata) GetX() int32 {
	return i.X
}

func (i *ImageMetadata) GetY() int32 {
	return i.Y
}

func (i *ImageMetadata) GetWidth() int32 {
	return i.Width
}

func (i *ImageMetadata) GetHeight() int32 {
	return i.Height
}

func (i *ImageMetadata) GetHorizontalFlip() bool {
	return i.HorizontalFlip
}
