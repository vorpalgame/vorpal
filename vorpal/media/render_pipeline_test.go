package media

import (
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/media/render"
	"log"
	"os"
	"testing"
)

func TestRendererPipeline(t *testing.T) {
	resultChannel := make(chan *render.Event, 1)
	renderPipeline := NewRenderPipeline(resultChannel)
	renderPipeline.AddRenderFunction(renderFuncOne)
	renderPipeline.AddRenderFunction(renderFuncTwo)
	go readResults(resultChannel)
	bus.GetVorpalBus().SendDrawEvent(bus.NewDrawLayersEvent())
	log.Println("Sent draw event...")

}
func readResults(resultChannel <-chan *render.Event) {
	for evt := range resultChannel {
		log.Println("Final result: ", evt)
		os.Exit(0)
	}
}

var renderFuncOne = func(inputChannel <-chan *render.Event, outputChannel chan<- *render.Event) {
	for evt := range inputChannel {
		log.Println(evt)
		outputChannel <- evt
	}
}

var renderFuncTwo = func(inputChannel <-chan *render.Event, outputChannel chan<- *render.Event) {
	for evt := range inputChannel {
		log.Println(evt)
		outputChannel <- evt
	}
}
