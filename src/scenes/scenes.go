package scenes

var Manager manager

func SetupManager() {
	var MainMenu mainMenu
	scenes := make([]Scene, 0)
	scenes = append(scenes, &MainMenu)

	Manager.CurrentScene = &MainMenu
	Manager.scenes = scenes
	Manager.CurrentScene.Create()
}

type manager struct {
	CurrentScene Scene
	scenes       []Scene
}

func (m *manager) ChangeScene(id int) {
	m.scenes[id].Create()
	m.CurrentScene = m.scenes[id]
}

type Scene interface {
	Update()
	Create()
	Destroy()
}
