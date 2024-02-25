package main

import (
	imgui "github.com/gabstv/cimgui-go"
	ebimgui "github.com/gabstv/ebiten-imgui/v3"
	"github.com/hajimehoshi/ebiten/v2"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	ebiten.SetWindowSize(1024, 768)

	gg := &G{}

	ebiten.RunGame(gg)

}

type G struct {
	w, h int
}

func (g *G) Draw(screen *ebiten.Image) {

	ebimgui.Draw(screen)
}

func (g *G) Update() error {
	ebimgui.Update(1.0 / 60.0)

	ebimgui.BeginFrame()
	defer ebimgui.EndFrame()
	imgui.BeginV("shite", nil, imgui.WindowFlagsNoTitleBar|imgui.WindowFlagsNoMove)
	defer imgui.End()

	imgui.SetWindowPosVec2(imgui.NewVec2(0, 0))
	imgui.SetWindowSizeVec2(imgui.NewVec2(float32(g.w), float32(g.h)))

	if imgui.Button("Don't press me.. I have press phobia") {
		speech := htgotts.Speech{Folder: "audio", Language: voices.English, Handler: &handlers.Native{}}
		speech.Speak("Did you just press the button? When the label said not to press it. Cringe")
		speech.Speak("hello world")
		//	os.RemoveAll(speech.Folder)
	}

	return nil
}

func (g *G) Layout(outsideWidth, outsideHeight int) (int, int) {

	g.w = outsideWidth
	g.h = outsideHeight

	ebimgui.SetDisplaySize(float32(g.w), float32(g.h))
	return g.w, g.h
}
