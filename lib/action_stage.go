package lib

import (
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

// Revisit the use of pointer for image.Image. There were isseus in marshaling so switched to value.
type ActionStageController interface {
	CheckBehaviorColorAt(x, y int32) color.Color
	SetControlMap(image *image.Image, width, height int) ActionStageController
	LoadControlMapFromFile(file string, width, height int) ActionStageController
}

type ActionStageControllerData struct {
	controlMap *image.Image
}

func (a *ActionStageControllerData) CheckBehaviorColorAt(x, y int32) color.Color {
	if a.controlMap != nil {
		img := *a.controlMap
		return img.At(int(x), int(y))
	}
	return color.White
}

func (a *ActionStageControllerData) SetControlMap(i *image.Image, width, height int) ActionStageController {

	// Resize:
	dest := resize.Resize(uint(width), uint(height), *i, resize.InterpolationFunction(resize.NearestNeighbor))
	a.controlMap = &dest

	return a
}

func writeTestFile(file string, dest image.Image) {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err = png.Encode(f, dest); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}

func (a *ActionStageControllerData) LoadControlMapFromFile(file string, width, height int) ActionStageController {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	a.SetControlMap(&img, width, height)
	return a
}
