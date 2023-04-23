package handlers

import (
	"syscall"
	"whisprite/core"
)

var (
	shutdown = core.Handler{
		Name:        "shutdown",
		ModRequired: true,
		Run: func(dispatch core.Dispatch, event core.Event, self core.Handler) {
			event.Say(event.Channel, "Shutting down now! Goodnight :)")
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		},
	}

	ping = core.Handler{
		Name: "ping",
		Run: func(dispatch core.Dispatch, event core.Event, self core.Handler) {
			event.Say(event.Channel, "Pong!")
		},
	}
)

var System = []core.Handler{
	shutdown,
	ping,
}
