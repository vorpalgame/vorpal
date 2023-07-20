package render

import (
	"github.com/vorpalgame/vorpal/bus"
	"strings"
)

type RenderEvent interface {
	GetDrawEvent() bus.DrawEvent
	AddComment(string)
	GetComments() string
}

type renderEvent struct {
	drawEvent bus.DrawEvent
	comments  []string
}

func (r *renderEvent) GetDrawEvent() bus.DrawEvent {
	return r.drawEvent
}

func (r *renderEvent) AddComment(comment string) {
	r.comments = append(r.comments, comment)
}

func (r *renderEvent) GetComments() string {
	return strings.Join(r.comments, "\n")
}

func NewEvent(evt bus.DrawEvent) RenderEvent {
	return &renderEvent{drawEvent: evt}
}
