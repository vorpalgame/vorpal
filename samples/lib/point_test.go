package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {

	p := NewPoint(0, 0)
	x := int32(0)
	assert.Equal(t, p.GetX(), x)
	assert.Equal(t, p.GetY(), int32(0))

	p2 := NewPoint(5, -5)
	p.Add(p2)
	assert.Equal(t, p.GetX(), int32(5))
	assert.Equal(t, p.GetY(), int32(-5))

}
