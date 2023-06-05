package bus

import "fmt"

//"fmt"

type LoadImagesListener interface {
	OnLoadImagesEvent(imageChannel <-chan ImagesEvent)
}

// ////Image Events
type imagesEvent struct {
	images []string
}

type ImagesEvent interface {
	GetImages() []string
	AddImage(img string)
}

func (i *imagesEvent) GetImages() []string {
	fmt.Println("Length of images in event: ", len(i.images))
	return i.images
}

func (i *imagesEvent) AddImage(img string) {
	//fmt.Println("Add image: " + img)
	i.images = append(i.images, img)
}
func NewImagesEvent() ImagesEvent {
	return &imagesEvent{}
}
