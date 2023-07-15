package util

import (
	"image"
	"sync"
)

// Minimal synchronization limited to cache access.
var (
	imageMapMutex = sync.RWMutex{}
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
	cache.imageCache = make(map[string]*image.Image, 100)

	return &cache
}
func (c *MediaCacheData) GetImage(fileName string) *image.Image {
	var img *image.Image
	imageMapMutex.Lock()
	img = c.imageCache[fileName]
	imageMapMutex.Unlock()
	return img

}

func (c *MediaCacheData) CacheImage(fileName string, img *image.Image) {
	imageMapMutex.Lock()
	c.imageCache[fileName] = img
	imageMapMutex.Unlock()
}
