package render

import "github.com/vorpalgame/vorpal/bus"

type Event interface {
	GetDrawEvent() bus.DrawEvent
}

type renderEvent struct {
	drawEvent bus.DrawEvent
}

func (r *renderEvent) GetDrawEvent() bus.DrawEvent {
	return r.drawEvent
}

func NewEvent(evt bus.DrawEvent) Event {
	return &renderEvent{evt}
}
