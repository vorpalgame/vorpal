package zombiecide

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
	"github.com/vorpalgame/vorpal/lib"
)

func TestBasicSubsumption(t *testing.T) {
	zombie := newSubsumptionZombie()

	//Basic test to ensure that mouse event triggers the
	//creation of (1) an ImageLayer and (2) puts ImageData on that layer.
	img := zombie.CreateImageLayer(newMouseEvent(0, 0))
	assert.Greater(t, len(img.GetLayerData()), 0)
	imgData := img.GetLayerData()[0]
	assert.NotNil(t, imgData)

	assert.Greater(t, len(imgData.GetFileName()), 0)

}

func newMouseEvent(x, y int32) bus.MouseEvent {
	return bus.NewMouseEvent(lib.NewMouseButtonState("left", false), lib.NewMouseButtonState("right", false), lib.NewMouseButtonState("middle", false), x, y)
}
