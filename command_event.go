package main

import (
	"github.com/Adeithe/go-twitch/irc"
	"github.com/anmitsu/go-shlex"

	"github.com/nicklaw5/helix/v2"
	"strings"
)

type CommandEvent struct {
	irc.ChatMessage

	Name   string
	Args   []string
	Twitch *helix.Client
}

func (c *CommandEvent) Parse() (err error) {
	c.Name = strings.Split(c.Text, " ")[0][1:]
	args, err := shlex.Split(c.Text, true)
	if err != nil {
		return err
	}

	c.Args = args
	return
}
