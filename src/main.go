package main

import (
	"dungeon_breakout/src/scenes"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "Dungeon Breakout")
	raygui.LoadStyle("./assets/style_terminal.rgs")
	scenes.SetupManager()

	for !rl.WindowShouldClose() {

		scenes.Manager.CurrentScene.Update()
	}
}
