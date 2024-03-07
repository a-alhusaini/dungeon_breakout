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
	"quit": func(data CommandInput) string {
		os.Exit(0)
		return "exited"
	},
	"clear": func(data CommandInput) string {
		data.g.Output = ""
		return ""
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
