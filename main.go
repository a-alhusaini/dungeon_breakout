package main

import (
	"strings"

	imgui "github.com/gabstv/cimgui-go"
	ebimgui "github.com/gabstv/ebiten-imgui/v3"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1024, 768)

	gg := &Game{input: "",
		output:      "You are in front of a door. \n There is a key on the floor",
		hasLockPick: false}

	ebiten.RunGame(gg)

}

type Game struct {
	width, height int
	input         string
	output        string
	hasLockPick   bool
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebimgui.Draw(screen)
}

func (g *Game) Update() error {
	ebimgui.Update(1.0 / 60.0)

	ebimgui.BeginFrame()
	defer ebimgui.EndFrame()

	imgui.BeginV("shite", nil, imgui.WindowFlagsNoTitleBar|imgui.WindowFlagsNoMove)
	defer imgui.End()

	imgui.SetWindowPosVec2(imgui.NewVec2(0, 0))
	imgui.SetWindowSizeVec2(imgui.NewVec2(float32(g.width), float32(g.height)))

	imgui.NewLine()

	for _, s := range strings.Split(g.output, "\n") {
		imgui.Text(s)
	}

	imgui.SetCursorPosY(float32(g.height) - imgui.TextLineHeight()*2)

	if imgui.InputTextWithHint("input", "", &g.input, imgui.InputTextFlagsEnterReturnsTrue, callback) {
		ParseInput(g)
	}

	imgui.SetKeyboardFocusHereV(-1)

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	g.width = outsideWidth
	g.height = outsideHeight

	ebimgui.SetDisplaySize(float32(g.width), float32(g.height))
	return g.width, g.height
}

func callback(data imgui.InputTextCallbackData) int {

	return 0
}
