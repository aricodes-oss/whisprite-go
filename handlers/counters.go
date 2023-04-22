package handlers

import "whisprite/core"

var newCounter = &core.Handler{
	Name:        "newcounter",
	ModRequired: true,
	Run: func(dispatch *core.Dispatch, event *core.Event, self *core.Handler) {
		log.Info("Hit it!")
	},
}

var Counters = []*core.Handler{
	newCounter,
}
