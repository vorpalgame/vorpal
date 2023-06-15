package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

// TODO Determine the right level of store/cache. Images might be kept in memory of system
// even out of scope but textures wouldn't be.
type MediaCache interface {
	CacheImages(event bus.DrawEvent)
	CacheFonts(event bus.TextEvent)
	GetImage(img string) *rl.Image
	DoCacheControl(event bus.ImageCacheEvent)
	GetFont(fontName string) *rl.Font
}

type mediaCache struct {
	imageCache map[string]*rl.Image
	fontCache  map[string]*rl.Font
}

func NewMediaCache() MediaCache {
	cache := mediaCache{}
	cache.imageCache = make(map[string]*rl.Image)
	cache.fontCache = make(map[string]*rl.Font)
	return &cache
}

func (c *mediaCache) GetFont(fontName string) *rl.Font {
	return c.fontCache[fontName]

}
func (c *mediaCache) CacheFonts(evt bus.TextEvent) {
	c.doFontCache(evt.GetFont())
	for _, line := range evt.GetText() {
		c.doFontCache(line.GetFont())
	}
}
func (c *mediaCache) doFontCache(fontName string) {
	if c.fontCache[fontName] == nil {
		font := rl.LoadFont(fontName)
		c.fontCache[fontName] = &font
	}
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
