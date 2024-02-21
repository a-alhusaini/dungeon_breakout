package scenes

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameScene struct{}

func (m *gameScene) Create() {

}
func (m *gameScene) Update() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	panelRect := rl.NewRectangle(0, 0, float32(rl.GetScreenHeight()), float32(rl.GetScreenWidth()))

	raygui.Panel(panelRect, "Game scene")

	text := "I \n am \n multiline \n text!"

	panelRect.Y = float32(raygui.GetFont().BaseSize * 2)
	panelRect.Height = 400

	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_ALIGNMENT_VERTICAL, int64(raygui.TEXT_ALIGN_TOP))

	rl.DrawRectangle(panelRect.ToInt32().X,
		panelRect.ToInt32().Y,
		panelRect.ToInt32().Width,
		panelRect.ToInt32().Height,
		rl.DarkBlue)

	raygui.Label(panelRect, text)

	inputBoxRect := rl.NewRectangle(0,
		float32(rl.GetScreenHeight())-48,
		float32(rl.GetScreenWidth()),
		48)
	inputText := "I am editable?"

	raygui.TextBox(inputBoxRect, &inputText, 32, true)

	rl.EndDrawing()
}
func (m *gameScene) Destroy() {}
