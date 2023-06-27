package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vorpalgame/vorpal/bus"
)

func TestImageSequenceCreator(t *testing.T) {
	baseFile := "foo"
	imgCreator := NewImageSequenceCreator(baseFile)
	imgCreator.AddRenderer(NewFlipHorizontalTransformer())

	fd := NewFrameData()

	fd.SetMaxFrame(5)
	cl := NewCurrentLocation(NewPoint(600, 500), -2, -2, 5, 5)
	img := imgCreator.CreateImage(bus.NewMouseEvent(nil, nil, nil, 800, 800), cl, fd)

	assert.Equal(t, "foo1", img.GetImage())
	assert.False(t, img.IsFlipHorizontal())

	imgFlipped := imgCreator.CreateImage(bus.NewMouseEvent(nil, nil, nil, -800, -800), cl, fd)
	assert.Equal(t, "foo1", imgFlipped.GetImage())
	assert.True(t, imgFlipped.IsFlipHorizontal())
}
