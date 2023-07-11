package raylibengine

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestRenderPipeline(t *testing.T) {
	tx := RenderTransaction{}
	pipeline := RenderPipelineData{}
	pipeline.Add(drawRenderer).Add(textRenderer).Add(textureRenderer)
	pipeline.Execute(tx)
	log.Println(tx.RenderImage)
	assert.NotNil(t, tx.RenderImage)
}
