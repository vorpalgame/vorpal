package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"log"
)

type RenderTransaction struct {
	RenderImage   *rl.Image
	RenderTexture *rl.Texture2D
	DrawEvent     bus.DrawEvent
	TextEvent     bus.TextEvent
}

type renderingPipelineFunction = func(tx RenderTransaction)

type RenderPipelineData struct {
	pipeline []renderingPipelineFunction
}

type RenderPipeline interface {
	Add(function renderingPipelineFunction) RenderPipeline
	Execute(transaction RenderTransaction)
}

func (r RenderPipelineData) Add(function renderingPipelineFunction) RenderPipeline {
	r.pipeline = append(r.pipeline, function)
	return r
}

func (r RenderPipelineData) Execute(transaction RenderTransaction) {
	for _, function := range r.pipeline {
		function(transaction)
	}
}

var drawRenderer = func(tx RenderTransaction) {
	log.Println("Draw Renderer")
	log.Println(tx.RenderImage)
}

var textRenderer = func(tx RenderTransaction) {
	log.Println("Text Renderer")
}

var textureRenderer = func(tx RenderTransaction) {
	log.Println("Text Renderer")

}
