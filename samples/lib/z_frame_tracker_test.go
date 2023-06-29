package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrameTracker(t *testing.T) {
	maxFrame := int32(4)
	f := NewFrameData()
	f.SetMaxFrame(maxFrame)
	f.SetToLoop(false)
	assert.False(t, f.IsLoop())
	assert.Equal(t, int32(2), f.Increment().GetCurrentFrame())
	f.Increment().Increment()
	assert.Equal(t, maxFrame, f.GetCurrentFrame())
	//Because we are not looping, the  frame data should
	//stay at max frame.
	assert.Equal(t, maxFrame, f.Increment().Increment().GetCurrentFrame())
	f.SetToLoop(true)
	//Now it should loop back to beginning

	assert.Equal(t, int32(1), f.Increment().GetCurrentFrame())

}
