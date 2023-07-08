package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO Refactoring left Point sort of stuck and inconsistent.
func TestNavigator(t *testing.T) {
	startX := int32(500)
	startY := int32(600)
	maxIncrement := int32(2)
	zeroIncrement := int32(0)
	//p := NewPoint(startX, startY) //Make X,Y asymmetric to avoid any confusion.
	//Current mouse is at 800,800 so we should move maximum.
	largeEvt := NewPoint(800, 800)
	largeNegativeEvt := NewPoint(-800, -800)
	//X,Y are only 3 off so no change to point to move is expected.
	smallEvt := PointData{503, 603}

	//2 and -2 are the maximum mouse moves and 5 is the window that determines
	//when something is "dead". In other words, if the mouse hasn't move more than 5 in X or Y then
	//don't change the position.
	//We can test behavior map separately
	l := NewNavigator(smallEvt, -maxIncrement, -maxIncrement, 5, 5, nil)

	//MoveByIncrement calculations will return maximum allowed movement in X,Y directions
	//as specified in start up. If X or Y are less than the offset window size
	//then no change will happen. For exmaple if x or y mouse position are
	//only 1 or 2 pixels difference and the offest is 5 pixels, the 0 is returned
	//indicating no movement should be recorded.
	//In the test the small event is only 3 pixels different and that's less than
	//the specified size of 5 for the offset so zero is returned.
	//Note that the calculation doesn't change current location as that
	//is a separate operation.
	var move Point = l.CalculateMoveIncrement(largeEvt)
	assert.Equal(t, maxIncrement, move.GetX())
	assert.Equal(t, maxIncrement, move.GetY())

	move = l.CalculateMoveIncrement(&smallEvt)
	assert.Equal(t, zeroIncrement, move.GetX())
	assert.Equal(t, zeroIncrement, move.GetY())

	//Now let's move based on the change. Max change allowed
	//is 2 so that's what we expect.
	l.MoveByIncrement(l.CalculateMoveIncrement(largeEvt))
	assert.Equal(t, startX+2, l.GetX())
	assert.Equal(t, startY+2, l.GetY())

	//Will increment back in the other direction now.
	l.MoveByIncrement(l.CalculateMoveIncrement(largeNegativeEvt))
	assert.Equal(t, startX, l.GetX())
	assert.Equal(t, startY, l.GetY())

}
