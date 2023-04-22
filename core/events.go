package core

import (
	"github.com/Adeithe/go-twitch/irc"
	"github.com/anmitsu/go-shlex"

	"strings"
)

type Event struct {
	irc.ChatMessage

	Name          string
	Args          []string
	IsMod         bool
	IsVIP         bool
	IsBroadcaster bool

	Say  func(string, string) error
	Sayf func(string, string, ...interface{}) error
}

func (e *Event) Parse() (err error) {
	e.Name = strings.Split(e.Text, " ")[0][1:]
	args, err := shlex.Split(e.Text, true)
	if err != nil {
		return err
	}

	e.Args = args[1:]
	return
}
