package core

import (
	"strings"
	"whisprite/query"
)

type Handler struct {
	Name        string
	Run         func(dispatch Dispatch, event Event, self Handler)
	ModRequired bool
	VipRequired bool
}

func (h *Handler) RespondsTo(name string) bool {
	iname := strings.ToLower(name)

	// Return early if it matches outright, no need to hang for the db
	if iname == h.Name {
		return true
	}

	a := query.Q.CommandAlias

	// Recursively resolve aliases
	found, err := a.Where(a.Name.Eq(iname)).First()
	if found != nil {
		for found.Target != h.Name {
			found, err = a.Where(a.Name.Eq(found.Target)).First()
		}
	}

	if err != nil {
		log.Debugf("Error finding alias for %s: %v", name, err)
		log.Debug(found)
	}

	return found != nil
}

func (h *Handler) Authenticate(event Event) bool {
	if !h.ModRequired && !h.VipRequired {
		return true
	}

	vipValid := h.VipRequired && (event.IsMod || event.IsVIP)
	modValid := (h.ModRequired && !h.VipRequired) && (event.IsMod)

	// Streamer can run whatever command whenevr
	return (vipValid || modValid) || event.IsBroadcaster
}
