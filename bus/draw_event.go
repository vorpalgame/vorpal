package bus

import "github.com/vorpalgame/vorpal/lib"

///// Constructors //////////////////

func NewDrawLayersEvent() DrawLayersEvent {
	evt := drawEvent{}
	evt.imageLayers = make([]lib.ImageLayerData, 0)
	return &evt
}

/////////////////////////////////////

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

type DrawEvent interface {
	Reset()
}

type DrawLayersEvent interface {
	DrawEvent
	GetImageLayers() []lib.ImageLayerData
	AddImageLayer(imgLayer lib.ImageLayerData) DrawEvent
}

type drawEvent struct {
	imageLayers []lib.ImageLayerData
}

func (evt *drawEvent) Reset() {
	evt.imageLayers = make([]lib.ImageLayerData, 0)
}

func (evt *drawEvent) AddImageLayer(img lib.ImageLayerData) DrawEvent {
	evt.imageLayers = append(evt.imageLayers, img)
	return evt
}

func (evt *drawEvent) GetImageLayers() []lib.ImageLayerData {
	return evt.imageLayers
}
