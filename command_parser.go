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
	"take": func(data CommandInput) string {
		if len(data.inputs) < 2 {
			return "I can't take nothing. Dumbass"
		}

		if data.inputs[1] == "key" {
			data.g.hasKey = true
			return "Taken"
		}

		return "something borked"
	},
	"quit": func(data CommandInput) string {
		os.Exit(0)
		return "exited"
	},
	"open": func(data CommandInput) string {
		g := data.g
		if len(data.inputs) < 2 {
			return "I can't open nothing"
		}

		if g.hasKey {
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
		g.Output = g.Output + "\n" + v(CommandInput{g, words})
	} else {
		g.Output = g.Output + "\n" + "I don't understand."
	}

	g.input = ""
}
