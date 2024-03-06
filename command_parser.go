package main

import (
	"os"
	"strings"
)

type CommandInput struct {
	g      *Game
	inputs []string
}

var commands = map[string]func(CommandInput) string{
	"take": func(data CommandInput) string { return "unimplemented" },
	"quit": func(data CommandInput) string {
		os.Exit(0)
		return "exited"
	},
	"open": func(data CommandInput) string {
		g := data.g
		if len(data.inputs) < 2 {
			return "I can't open nothing"
		}

		if g.hasLockPick {
			return "You open the door and escape."
		} else {
			return "The door is locked"
		}
	},
}

func ParseInput(g *Game) {
	words := strings.Split(g.input, " ")
	v, exists := commands[words[0]]
	if exists {
		g.output = g.output + "\n" + v(CommandInput{g, words})
	} else {
		g.output = g.output + "\n" + "I don't understand."
	}

	g.input = ""
}
