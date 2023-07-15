package util

///No x11 packages here...
import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
	"golang.org/x/image/draw"
	"image"
)

func NewLoadResizeCachePipeline(mediaCache *MediaCache, inputChannel chan bus.DrawLayersEvent) {
	go loadResizeCachePipeline(mediaCache, inputChannel)

}

type loadResizeCacheData struct {
	loadedImage   *image.Image
	imageMetadata *lib.ImageMetadata
	resizedImage  *image.RGBA
}

var loadResizeCachePipeline = func(mediaCache *MediaCache, inputChannel <-chan bus.DrawLayersEvent) {

	loadChan := make(chan *loadResizeCacheData, 100)
	go loadImageFunc(mediaCache, loadChan)

	for evt := range inputChannel {
		for _, layer := range evt.GetImageLayers() {
			for _, imgData := range layer.LayerMetadata {
				img := (*mediaCache).GetImage(imgData.ImageFileName)
				if img == nil {
					loadChan <- &loadResizeCacheData{imageMetadata: imgData}
				}
			}
		}

	}
}

var loadImageFunc = func(mediaCache *MediaCache, inputChannel chan *loadResizeCacheData) {
	resizeChan := make(chan *loadResizeCacheData, 100)
	go resizeImageFunc(mediaCache, resizeChan)
	for evt := range inputChannel {
		evt.loadedImage = LoadImage(evt.imageMetadata.ImageFileName)
		resizeChan <- evt
	}
}

var resizeImageFunc = func(mediaCache *MediaCache, inputChannel chan *loadResizeCacheData) {
	cacheChan := make(chan *loadResizeCacheData, 100)
	go cacheImageFunc(mediaCache, cacheChan)
	for evt := range inputChannel {

		toRect := image.Rect(0, 0, int(evt.imageMetadata.Width), int(evt.imageMetadata.Height))
		resizedImage := image.NewRGBA(toRect)
		draw.BiLinear.Scale(resizedImage, resizedImage.Rect, *evt.loadedImage, (*evt.loadedImage).Bounds(), draw.Over, nil)
		evt.resizedImage = resizedImage
		cacheChan <- evt
	}
}

var cacheImageFunc = func(cache *MediaCache, inputChannel chan *loadResizeCacheData) {
	for evt := range inputChannel {
		store := image.Image(evt.resizedImage)
		(*cache).CacheImage(evt.imageMetadata.ImageFileName, &store)
	}
}
