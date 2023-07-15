package util

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func LoadImage(imageFileName string) *image.Image {
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

func Write(buffer *image.RGBA) {
	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(out, buffer)
}
