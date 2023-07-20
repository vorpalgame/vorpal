package media

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/media/render"
	"log"
)

func (data *renderPipelineData) OnDrawEvent(drawEventChannel <-chan bus.DrawEvent) {
	for evt := range drawEventChannel {
		log.Println("Received draw event")
		data.inputDrawEventChannel <- evt

	}
}
func NewRenderPipeline(returnChannel chan render.RenderEvent) RenderPipeline {
	data := renderPipelineData{inputDrawEventChannel: make(chan bus.DrawEvent, 1), inputChannel: createChannel(), outputChannel: createChannel(), returnChannel: returnChannel, cache: render.NewImageCache()}
	bus.GetVorpalBus().AddDrawEventListener(&data)
	go data.pumpInputEvents()
	go data.pumpOutputEvents()
	return &data
}

type Render func(event render.RenderEvent)

type RenderPipeline interface {
	Add(renderer Render) RenderPipeline
}
type renderPipelineData struct {
	inputDrawEventChannel                      chan bus.DrawEvent
	inputChannel, outputChannel, returnChannel chan render.RenderEvent
	head, tail                                 *renderer
	cache                                      render.ImageCache
}

func createChannel() chan render.RenderEvent {
	return make(chan render.RenderEvent, 1)
}

func (data *renderPipelineData) pumpInputEvents() {
	for evt := range data.inputDrawEventChannel {
		renderEvent := render.NewEvent(evt)
		data.inputChannel <- renderEvent
	}
}

func (data *renderPipelineData) pumpOutputEvents() {
	for evt := range data.outputChannel {
		data.returnChannel <- evt
	}
}

func (data *renderPipelineData) Add(render Render) RenderPipeline {
	newRenderer(data, render)
	return data
}

type Renderer interface {
	DoRender()
}
type renderer struct {
	inputChannel, outputChannel chan render.RenderEvent
	render                      Render
}

func (r *renderer) DoRender() {
	for evt := range r.inputChannel {
		r.render(evt)
		r.outputChannel <- evt
	}
}

func newRenderer(data *renderPipelineData, render Render) {

	if data.head == nil {
		data.head = &renderer{render: render, inputChannel: data.inputChannel, outputChannel: data.outputChannel}
		data.tail = data.head

	} else {
		bridge := createChannel()
		r := &renderer{render: render, inputChannel: bridge, outputChannel: data.outputChannel}
		data.tail.outputChannel = bridge
		data.tail = r

	}
	go data.tail.DoRender()
}
