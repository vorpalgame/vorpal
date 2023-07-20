package media

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/media/render"
	"log"
)

func (data *renderPipelineData) OnDrawEvent(inputChannel <-chan bus.DrawEvent) {
	for evt := range inputChannel {
		log.Println("Received draw event")
		renderEvent := render.NewEvent(evt)
		data.inputChannel <- &renderEvent

	}
}
func NewRenderPipeline(outputChannel chan *render.Event) RenderPipeline {
	data := renderPipelineData{inputChannel: createRenderChannel(), outputChannel: outputChannel, renderPipeline: []Renderer{}, cache: render.NewImageCache()}
	bus.GetVorpalBus().AddDrawEventListener(&data)
	return &data
}

type Renderer func(inputChannel <-chan *render.Event, outputChannel chan<- *render.Event)

type RenderPipeline interface {
	AddRenderFunction(renderer Renderer)
}
type renderPipelineData struct {
	inputChannel, outputChannel chan *render.Event
	renderPipeline              []Renderer
	cache                       render.ImageCache
}

func (data *renderPipelineData) AddRenderFunction(renderer Renderer) {
	data.renderPipeline = append(data.renderPipeline, renderer)
	go renderer(createRenderChannel(), createRenderChannel())
}

func createRenderChannel() chan *render.Event {
	return make(chan *render.Event, 1)
}
