package main

import (
	"github.com/nicklaw5/helix/v2"
)

type Dispatch struct {
	handlers []*Handler
	Twitch   *helix.Client
}

func (d *Dispatch) Register(handler *Handler) {
	d.handlers = append(d.handlers, handler)
}

func (d *Dispatch) Handle(event *Event) {
	for idx, handler := range d.handlers {
		if handler.RespondsTo(event.Name) {
			log.Debugf("Command %s (position %v) matched handler %s, firing", event.Name, idx, handler.Name)
			handler.Run(d, event, handler)
		}
	}
}
