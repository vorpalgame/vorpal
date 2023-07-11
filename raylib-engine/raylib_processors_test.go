package raylibengine

import (
	"testing"
)

func TestRenderPipeline(t *testing.T) {
	tx := RenderTransaction{}
	pipeline := RenderPipelineData{}
	pipeline.Add(drawRenderer).Add(textRenderer).Add(textureRenderer)
	pipeline.Execute(tx)

}
