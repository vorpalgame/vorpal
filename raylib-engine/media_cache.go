package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// TODO Determine the right level of store/cache. Images might be kept in memory of system
// even out of scope but textures wouldn't be.
type MediaCache interface {
	CacheImages(event bus.DrawEvent)
	GetImage(img string) *rl.Image
	DoCacheControl(event bus.ImageCacheEvent)
}

type mediaCache struct {
	imageCache map[string]*rl.Image
}

func NewMediaCache() MediaCache {
	cache := mediaCache{}
	cache.imageCache = make(map[string]*rl.Image)
	return &cache
}
func (c *mediaCache) GetImage(img string) *rl.Image {
	return c.imageCache[img]
}
func (c *mediaCache) CacheImages(evt bus.DrawEvent) {
	for _, evt := range evt.GetImageLayers() {
		img := c.imageCache[evt.GetImage()]
		if img == nil {
			c.imageCache[evt.GetImage()] = rl.LoadImage(evt.GetImage())
			img = c.imageCache[evt.GetImage()]
			rl.ImageResize(img, evt.GetWidth(), evt.GetHeight())
		}

	}
}

func (c *mediaCache) DoCacheControl(evt bus.ImageCacheEvent) {
	if evt != nil {
		for _, op := range evt.GetImageCacheOperations() {
			if op.GetOperation() == "add" {
				if c.imageCache[op.GetImage()] == nil {
					c.imageCache[op.GetImage()] = rl.LoadImage(op.GetImage())
				}
			}
		}
	}
}
