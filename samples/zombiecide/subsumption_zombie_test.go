package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
)

func TestBasicSubsumption(t *testing.T) {
	zombie := newSubsumptionZombie()

	//Basic test to ensure that mouse event triggers the
	//creation of (1) an ImageLayer and (2) puts ImageData on that layer.
	img := zombie.CreateImageLayer(newMouseEvent(0, 0))
	assert.Greater(t, len(img.GetLayerData()), 0)
	imgData := img.GetLayerData()[0]
	assert.NotNil(t, imgData)

	assert.Greater(t, len(imgData.GetImage()), 0)

}

func newMouseEvent(x, y int32) bus.MouseEvent {
	return bus.NewMouseEvent(bus.NewMouseButtonState("left", false), bus.NewMouseButtonState("right", false), bus.NewMouseButtonState("middle", false), x, y)
}
