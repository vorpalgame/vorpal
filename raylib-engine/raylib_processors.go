package raylibengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vorpalgame/vorpal/bus"
	"log"
)

type RenderTransaction struct {
	DrawEvent     bus.DrawEvent
	TextEvent     bus.TextEvent
	MediaCache    MediaCache
	RenderImage   *rl.Image //Mot used currently but will be in future instead of cache.
	RenderTexture *rl.Texture2D
}

func NewRenderTransaction(drawEvent bus.DrawEvent, textEvent bus.TextEvent, mediaCache MediaCache) RenderTransaction {
	return RenderTransaction{drawEvent, textEvent, mediaCache, nil, nil}
}

type renderingPipelineFunction = func(tx RenderTransaction)

func NewRendererPipeline() RendererPipeline {
	pipeline := RenderPipelineData{}
	pipeline.Add(drawRenderer).Add(textRenderer).Add(textureRenderer)
	return &pipeline
}

type RenderPipelineData struct {
	pipeline []renderingPipelineFunction
}

type RendererPipeline interface {
	Add(function renderingPipelineFunction) RendererPipeline
	Execute(transaction RenderTransaction)
}

func (r *RenderPipelineData) Add(function renderingPipelineFunction) RendererPipeline {
	r.pipeline = append(r.pipeline, function)
	return r
}

func (r *RenderPipelineData) Execute(transaction RenderTransaction) {
	//log.Println(r.pipeline)
	for _, function := range r.pipeline {
		function(transaction)
	}
}

var drawRenderer = func(tx RenderTransaction) {
	//log.Println("Drawing renderer")
	raylibProcessDrawEvent(tx.DrawEvent, tx.MediaCache)
}

var textRenderer = func(tx RenderTransaction) {
	//log.Println("Text Renderer")
	raylibProcessTextEvent(tx.TextEvent, tx.MediaCache)
}

var textureRenderer = func(tx RenderTransaction) {
	log.Println("Text Renderer")

}
