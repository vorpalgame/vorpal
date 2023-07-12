package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

type MediaCache interface {
	CacheImages(event bus.DrawLayersEvent)
	CacheFonts(event lib.Font)
	DoCacheControl(event bus.ImageCacheEvent)
	GetImage(img string) *rl.Image
	GetFont(fontName string) *rl.Font
	GetAudio(fileName string) *rl.Sound
}

type MediaCacheData struct {
	imageCache map[string]*rl.Image
	fontCache  map[string]*rl.Font
	audioCache map[string]*rl.Sound
}

func NewMediaCache() MediaCache {
	cache := MediaCacheData{}
	cache.imageCache = make(map[string]*rl.Image)
	cache.fontCache = make(map[string]*rl.Font)
	cache.audioCache = make(map[string]*rl.Sound)
	return &cache
}

func (c *MediaCacheData) GetFont(fontName string) *rl.Font {
	return c.fontCache[fontName]

}
func (c *MediaCacheData) GetAudio(fileName string) *rl.Sound {
	if c.audioCache[fileName] == nil {
		audio := rl.LoadSound(fileName)
		c.audioCache[fileName] = &audio
	}
	return c.audioCache[fileName]

}

func (c *MediaCacheData) CacheFonts(evt lib.Font) {

	c.doFontCache(evt.GetFont())
	//TODO Refactor as necessary.
	// for _, line := range evt.GetText() {
	// 	c.doFontCache(line.GetFont())
	// }

}
func (c *MediaCacheData) doFontCache(fontName string) {
	if c.fontCache[fontName] == nil {
		font := rl.LoadFont(fontName)
		c.fontCache[fontName] = &font
	}
}

func (c *MediaCacheData) GetImage(img string) *rl.Image {
	return c.imageCache[img]
}

// TODO Need an update mechanism when scale changes. Perhaps map key needs to be
// name+scale. This may also be where the image cache  event comes in.
// In any case, it appears that scaling per image draw is a bit to expensive.

func (c *MediaCacheData) CacheImages(evt bus.DrawLayersEvent) {
	for _, evt := range evt.GetImageLayers() {
		for _, imgData := range evt.LayerMetadata {
			img := c.imageCache[imgData.ImageFileName]
			if img == nil {
				newImg := rl.LoadImage(imgData.ImageFileName)

				rl.ImageResize(newImg, imgData.Width, imgData.Height)
				c.imageCache[imgData.ImageFileName] = newImg
			}
		}

	}
}

func (c *MediaCacheData) DoCacheControl(evt bus.ImageCacheEvent) {
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
