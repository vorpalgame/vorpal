package media

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/media/render"
	"log"
	"testing"
	"time"
)

func TestRendererPipeline(t *testing.T) {
	resultChannel := make(chan render.RenderEvent, 1)
	renderPipeline := NewRenderPipeline(resultChannel)
	renderPipeline.Add(renderFuncOne)
	renderPipeline.Add(renderFuncTwo)
	bus.GetVorpalBus().SendDrawEvent(bus.NewDrawLayersEvent())
	log.Println("Sent draw event...")
	go readResults(resultChannel)

	time.Sleep(5 * time.Second)
}
func readResults(resultChannel <-chan render.RenderEvent) {
	for evt := range resultChannel {
		log.Println("Final result: ", evt.GetComments())
	}
}

var renderFuncOne = func(event render.RenderEvent) {
	log.Println("Invoked one...")
	event.AddComment("Render function one...")
}

var renderFuncTwo = func(event render.RenderEvent) {
	log.Println("Invoked two....")
	event.AddComment("Render function two...")
}
