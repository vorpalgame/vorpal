package bus

import "github.com/vorpalgame/vorpal/lib"

///// Constructors //////////////////

func NewDrawLayersEvent() DrawLayersEvent {
	evt := drawEvent{}
	evt.imageLayers = make([]lib.ImageLayer, 0)
	return &evt
}

/////////////////////////////////////

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

type DrawEventProcessor interface {
	ProcessDrawEvent(evt DrawEvent)
}

type DrawEvent interface {
	Reset()
}

type DrawLayersEvent interface {
	DrawEvent
	GetImageLayers() []lib.ImageLayer
	AddImageLayer(imgLayer lib.ImageLayer) DrawEvent
}

type drawEvent struct {
	imageLayers []lib.ImageLayer
}

func (evt *drawEvent) Reset() {
	evt.imageLayers = make([]lib.ImageLayer, 0)
}

func (evt *drawEvent) AddImageLayer(img lib.ImageLayer) DrawEvent {
	evt.imageLayers = append(evt.imageLayers, img)
	return evt
}

func (evt *drawEvent) GetImageLayers() []lib.ImageLayer {
	return evt.imageLayers
}
