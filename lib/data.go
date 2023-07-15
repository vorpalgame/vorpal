package lib

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Common struct used by all that are also yaml annotated for marshaling.
// Marshaling creates issues when using interfaces.

type ImageLayerData struct {
	LayerMetadata []*ImageMetadata `yaml:"LayerMetadata"`
}

func UnmarshalImageLayer(fileName string) *ImageLayerData {
	value := ImageLayerData{}
	unmarshal(readFile(fileName), &value)
	return &value
}

func MarshalImageLayer(fileName string, value *ImageLayerData) {
	marshal(fileName, value)
}

/////////////////////////////////////////////////////////////////////////

type ImageMetadata struct {
	ImageFileName  string `yaml:"ImageFileName"`
	X              int32  `yaml:"X"`
	Y              int32  `yaml:"Y"`
	Width          int32  `yaml:"Width"`
	Height         int32  `yaml:"Height"`
	HorizontalFlip bool   `yaml:"HorizontalFlip"`
}

func UnmarshalImageMetadata(fileName string) *ImageMetadata {
	value := ImageMetadata{}
	unmarshal(readFile(fileName), &value)
	return &value
}

func MarshalImageMetadata(fileName string, imgLayer *ImageMetadata) {
	marshal(fileName, imgLayer)
}

// /////////////////////////////////////////////////////////////////////////////
type NavigatorData struct {
	X                     int32 `yaml:"CurrentX"`
	Y                     int32 `yaml:"CurrentY"`
	XMove                 int32 `yaml:"XMove"`
	YMove                 int32 `yaml:"YMove"`
	MaxXOffset            int32 `yaml:"MaxXOffset"`
	MaxYOffset            int32 `yaml:"MaxYOffset"`
	ActionStageController `yaml:"-"`
}

func UnmarshalNavigator(fileName string) *NavigatorData {
	value := NavigatorData{}
	unmarshal(readFile(fileName), &value)
	return &value
}

func MarshalNavigator(fileName string, value *NavigatorData) {
	marshal(fileName, value)
}

// ///////////////////////////////////////////////////////////////////////////////
type Scene struct {
	WindowTitle  string          `yaml:"WindowTitle"`
	WindowWidth  int             `yaml:"WindowWidth"`
	WindowHeight int             `yaml:"WindowHeight"`
	RegisterKeys []rune          `yaml:"RegisterKeys"`
	Background   *ImageLayerData `yaml:"Background"`
	BehaviorMap  *ImageLayerData `yaml:"BehaviorMap"`
	Foreground   *ImageLayerData `yaml:"Foreground"`
	Actors       []string        `yaml:"Actors"`
}

func UnmarshalScene(fileName string) *Scene {
	value := Scene{}
	unmarshal(readFile(fileName), &value)
	return &value
}

func MarshalScene(fileName string, value *Scene) {
	marshal(fileName, value)
}

// ///////////////////////////////////////////////////////////////////////////////
type FontData struct {
	Font     string `yaml:"Font"`
	FontSize int32  `yaml:"FontSize"`
}

// ///////////////////////////////////////////////////////////////////////////////
type TextLineData struct {
	Test string `yaml:"Text"`
	Font
}

/////////////////////////////////////////////////////////////////////////////////

func marshal(file string, data interface{}) {
	marshaled, _ := yaml.Marshal(data)

	e := os.WriteFile(file, marshaled, 0644)
	if e != nil {
		panic(e)
	}
}

func unmarshal(contents []byte, toType interface{}) {
	e := yaml.Unmarshal(contents, toType)
	if e != nil {
		panic(e)
	}
}

func readFile(fileName string) []byte {
	f, e := os.ReadFile(fileName)
	if e != nil {
		log.Default().Println(e)
		os.Exit(1)
	}
	return f
}
