package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
)

type MediaCache interface {
	CacheImages(event bus.DrawLayersEvent) MediaCache
	CacheFonts(event bus.Font) MediaCache
	SetCurrentRenderImage(img *rl.Image) MediaCache
	DoCacheControl(event bus.ImageCacheEvent) MediaCache
	GetImage(img string) *rl.Image
	GetFont(fontName string) *rl.Font
	GetAudio(fileName string) *rl.Sound
	GetCurrentRenderImage() *rl.Image
}

type mediaCache struct {
	imageCache         map[string]*rl.Image
	fontCache          map[string]*rl.Font
	audioCache         map[string]*rl.Sound
	currentRenderImage *rl.Image
}

func NewMediaCache() MediaCache {
	cache := mediaCache{}
	cache.imageCache = make(map[string]*rl.Image)
	cache.fontCache = make(map[string]*rl.Font)
	cache.audioCache = make(map[string]*rl.Sound)
	return &cache
}
func (c *mediaCache) GetCurrentRenderImage() *rl.Image {
	return c.currentRenderImage
}
func (c *mediaCache) SetCurrentRenderImage(img *rl.Image) MediaCache {
	c.currentRenderImage = img
	return c
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

func (c *mediaCache) CacheFonts(evt bus.Font) MediaCache {
	c.doFontCache(evt.GetFont())
	//TODO Refactor as necessary.
	// for _, line := range evt.GetText() {
	// 	c.doFontCache(line.GetFont())
	// }
	return c
}
func (c *mediaCache) doFontCache(fontName string) MediaCache {
	if c.fontCache[fontName] == nil {
		font := rl.LoadFont(fontName)
		c.fontCache[fontName] = &font
	}
	return c
}

func (c *mediaCache) GetImage(img string) *rl.Image {
	return c.imageCache[img]
}

// Need an update mechanism when scale changes. Perhaps map key needs to be
// name+scale. This may also be where the image cache  event comes in.
// In any case, it appears that scaling per image draw is a bit to expensive.
func (c *mediaCache) CacheImages(evt bus.DrawLayersEvent) MediaCache {
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
	return c
}

func (c *mediaCache) DoCacheControl(evt bus.ImageCacheEvent) MediaCache {
	if evt != nil {
		for _, op := range evt.GetImageCacheOperations() {
			if op.GetOperation() == "add" {
				if c.imageCache[op.GetImage()] == nil {
					c.imageCache[op.GetImage()] = rl.LoadImage(op.GetImage())
				}
			}
		}
	}
	return c
}
