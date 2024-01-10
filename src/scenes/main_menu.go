package scenes

import (
	"os"

	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	menuX      float32
	menuY      float32
	menuWidth  float32
	menuHeight float32
)

type mainMenu struct{}

func (m *mainMenu) Create() {
	menuX = float32((rl.GetScreenWidth() / 2) - rl.GetScreenWidth()/6)
	menuY = float32(rl.GetScreenHeight() / 5)
	menuWidth = float32(rl.GetScreenWidth() / 3)
	menuHeight = float32(rl.GetScreenHeight() / 2)
}

func (m *mainMenu) Update() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	drawUI()

	rl.EndDrawing()
}

func (m *mainMenu) Destroy() {}

func drawUI() {

	menuPanelRec := rl.NewRectangle(menuX, menuY, menuWidth, menuHeight)
	raygui.Panel(menuPanelRec, "Dungeon Breakout")

	startBtnRec := rl.NewRectangle(menuX+25, menuY+50, menuWidth-50, 40)
	startBtn := raygui.Button(startBtnRec, "Play")

	if startBtn {
		Manager.ChangeScene(1)
	}

	quitBtnRec := rl.NewRectangle(
		menuX+25,
		startBtnRec.Y+startBtnRec.Height+25,
		startBtnRec.Width,
		startBtnRec.Height)

	quitBtn := raygui.Button(quitBtnRec, "Quit")

	if quitBtn {
		os.Exit(0)
	}
}
