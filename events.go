package main

import (
	"github.com/Adeithe/go-twitch/irc"
	"github.com/anmitsu/go-shlex"

	"strings"
)

type Event struct {
	irc.ChatMessage

	Name string
	Args []string
}

func (c *Event) Parse() (err error) {
	c.Name = strings.Split(c.Text, " ")[0][1:]
	args, err := shlex.Split(c.Text, true)
	if err != nil {
		return err
	}

	c.Args = args
	return
}
