package main

import (
	"dungeon_breakout/state_manager"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type CommandInput struct {
	g      *Game
	inputs []string
}

var builtinCommands = map[string]func(CommandInput) string{
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
	inputs := strings.Split(g.input, " ")

	builtinCommandFunc, isBuiltinCommand := builtinCommands[inputs[0]]
	if isBuiltinCommand {
		g.Output = g.Output + "\n" + builtinCommandFunc(CommandInput{g, inputs})
		g.input = ""
		return
	}

	foundMapCommand, commandID := isMapCommand(g.state, inputs)
	if foundMapCommand {
		g.Output = g.Output + "\n" + doMapCommand(g, inputs, commandID)
		g.input = ""
		return
	}

	g.Output = g.Output + "\n" + "I don't understand"

	g.input = ""
}

func isMapCommand(state string, inputs []string) (bool, int64) {
	didFindMapCommand := false
	var commandID int64 = 0
	state_manager.Get(state, "commands").ForEach(func(key, value gjson.Result) bool {
		if value.Get("name").String() == inputs[0] {
			didFindMapCommand = true
			commandID = value.Get("id").Int()
			return false
		}
		return true
	})
	return didFindMapCommand, commandID
}

func doMapCommand(g *Game, inputs []string, commandID int64) string {

	if len(inputs) < 2 {
		return "I don't understand"

	}

	objectCount := state_manager.Get(g.state, "rooms.0.objects.#").Int()

	for i := 0; i < int(objectCount); i++ {

		if inputs[1] == state_manager.Get(g.state, genObjectPath(0, i)+".name").String() {
			return state_manager.Get(
				g.state, genIneractionPath(0, i, 0)+".result.output").String()
		}
	}

	return "I don't understand"
}

func genRoomPath(roomID int) string {
	return "rooms." + strconv.Itoa(roomID)
}

func genObjectPath(roomID, objectID int) string {
	return genRoomPath(roomID) + ".objects." + strconv.Itoa(objectID)
}

func genIneractionPath(roomID, objectID, interactionID int) string {
	return genObjectPath(roomID, objectID) + ".interactions." + strconv.Itoa(interactionID)
}
