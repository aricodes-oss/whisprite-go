package main

import (
	"strings"
	"whisprite/query"
)

type Handler struct {
	Name        string
	Run         func(dispatch *Dispatch, event *Event, self *Handler)
	ModRequired bool
	VipRequired bool
}

func (h *Handler) RespondsTo(name string) bool {
	iname := strings.ToLower(name)
	a := query.Alias

	matchingAlias, err := a.Where(a.Name.Eq(iname), a.Target.Eq(h.Name)).First()
	if err != nil {
		log.Errorf("Error finding alias for %s: %v", name, err)
	}

	return iname == h.Name || matchingAlias != nil
}

func (h *Handler) Authenticate(event *Event) bool {
	if !h.ModRequired && !h.VipRequired {
		return true
	}

	vipValid := h.VipRequired && (event.IsMod || event.IsVIP)
	modValid := (h.ModRequired && !h.VipRequired) && (event.IsMod)

	// Streamer can run whatever command whenevr
	return (vipValid || modValid) || event.IsBroadcaster
}
