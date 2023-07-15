package util

import (
	"image"
)

type MediaCache interface {
	GetImage(fileName string) *image.Image
	CacheImage(fileName string, img *image.Image)
}
type MediaCacheData struct {
	imageCache map[string]*image.Image
}

func NewMediaCache() MediaCache {
	cache := MediaCacheData{}
	cache.imageCache = make(map[string]*image.Image)

	return &cache
}
func (c *MediaCacheData) GetImage(fileName string) *image.Image {
	img := c.imageCache[fileName]

	return img
}

func (c *MediaCacheData) CacheImage(fileName string, img *image.Image) {
	c.imageCache[fileName] = img
}

//
//func loadImage(imgData *lib.ImageMetadata, c *MediaCacheData) {
//	img := LoadImage(imgData.ImageFileName)
//	toRect := image.Rect(0, 0, int(imgData.Width), int(imgData.Height))
//	resizedImage := image.NewRGBA(toRect)
//
//	draw.BiLinear.Scale(resizedImage, resizedImage.Rect, *img, (*img).Bounds(), draw.Over, nil)
//	//This is goofy and I'm sure there's a better way but moving on...
//	store := image.Image(resizedImage)
//	c.imageCache[imgData.ImageFileName] = &store
//}
