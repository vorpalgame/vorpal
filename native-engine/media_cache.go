package native_engine

import (
	"github.com/vorpalgame/vorpal/bus"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"log"
	"os"
)

type MediaCache interface {
	GetImage(img string) *image.Image
	CacheImages(evt bus.DrawLayersEvent)
}
type MediaCacheData struct {
	imageCache map[string]*image.Image
}

func NewMediaCache() MediaCache {
	cache := MediaCacheData{}
	cache.imageCache = make(map[string]*image.Image)

	return &cache
}
func (c *MediaCacheData) GetImage(img string) *image.Image {
	return c.imageCache[img]
}

// Need a better mechanism than nil checks.
func (c *MediaCacheData) CacheImages(evt bus.DrawLayersEvent) {

	//log.Println(evt)
	//log.Println(evt.GetImageLayers())
	for _, layer := range evt.GetImageLayers() {
		for _, imgData := range layer.LayerMetadata {
			img := c.imageCache[imgData.ImageFileName]
			if img == nil {
				img = loadImage(imgData.ImageFileName)
				toRect := image.Rect(0, 0, int(imgData.Width), int(imgData.Height))
				resizedImage := image.NewRGBA(toRect)

				draw.NearestNeighbor.Scale(resizedImage, resizedImage.Rect, *img, (*img).Bounds(), draw.Over, nil)
				//This is goofy and I'm sure there's a better way but moving on...
				store := image.Image(resizedImage)
				c.imageCache[imgData.ImageFileName] = &store
			}

		}
	}
}

func loadImage(imageFileName string) *image.Image {
	log.Println(imageFileName)
	f, err := os.Open(imageFileName)
	if err != nil {
		log.Fatal(err)

	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return &img
}
