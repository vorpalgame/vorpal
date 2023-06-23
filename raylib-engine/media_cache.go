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
	GetAudio(fileName string) *rl.Sound
}

type mediaCache struct {
	imageCache map[string]*rl.Image
	fontCache  map[string]*rl.Font
	audioCache map[string]*rl.Sound
}

func NewMediaCache() MediaCache {
	cache := mediaCache{}
	cache.imageCache = make(map[string]*rl.Image)
	cache.fontCache = make(map[string]*rl.Font)
	cache.audioCache = make(map[string]*rl.Sound)
	return &cache
}

func (c *mediaCache) GetFont(fontName string) *rl.Font {
	return c.fontCache[fontName]

}
func (c *mediaCache) GetAudio(fileName string) *rl.Sound {
	if c.audioCache[fileName] == nil {
		audio := rl.LoadSound(fileName)
		c.audioCache[fileName] = &audio
	}
	return c.audioCache[fileName]

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
		for _, imgData := range evt.GetLayerData() {
			img := c.imageCache[imgData.GetImage()]
			if img == nil {
				newImg := rl.LoadImage(imgData.GetImage())

				width := int32(float32(newImg.Width) * imgData.GetScalePercent())
				height := int32(float32(newImg.Height) * imgData.GetScalePercent())

				rl.ImageResize(newImg, width, height)
				c.imageCache[imgData.GetImage()] = newImg
			}
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
