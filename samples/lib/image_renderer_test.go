package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
)

func TestImageRenderer(t *testing.T) {
	//Image renderer is really doing metadata rendering.

	ir := NewImageRenderer()
	fd := NewFrameData()
	fd.SetMaxFrame(5)
	cl := NewCurrentLocation(NewPoint(600, 500), -2, -2, 5, 5)
	file := "foo.png"
	scale := int32(50)
	evt := bus.NewMouseEvent(nil, nil, nil, 800, 800)
	ir.SetImageName(file).SetScale(scale)

	il := ir.CreateImageLayer(evt, fd, cl)
	firstInLayer := il.GetLayerData()[0]
	assert.Equal(t, fmt.Sprintf(file, int32(1)), firstInLayer.GetImage())
}
