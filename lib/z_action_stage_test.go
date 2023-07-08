package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestActionStage(t *testing.T) {
	ac := ActionStageControllerData{}
	ac.LoadControlMapFromFile("../samples/resources/zombiecide/behaviormap.png", 1920, 1080)
	assert.NotNil(t, ac)
	pd := &PointData{100, 100}
	currentColor := ac.CheckBehaviorColorAt(pd)
	assert.NotNil(t, currentColor)

	//Walk diagonally and check different points.
	for i := int32(1); i < 100; i++ {
		pd.X = pd.X + 10
		pd.Y = pd.Y + 10
		currentColor = ac.CheckBehaviorColorAt(pd)
		r, g, b, a := currentColor.RGBA()
		log.Default().Println(fmt.Sprintf("Red: %d, Green: %d, Blue: %d, Alpha: %d", r, g, b, a))

		log.Default().Println(fmt.Sprintf("Red: %d, Green: %d, Blue: %d, Alpha: %d", r/256, g/256, b/256, a/256))
	}

}
