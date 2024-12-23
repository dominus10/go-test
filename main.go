package main

import (
	"game/injection"
	"game/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {

	// var chContainer CharactersContainer
	// data, _:= os.ReadFile("./save.yaml")

	// err := yaml.Unmarshal([]byte(data), &chContainer)
	// if err != nil {
	// 	log.Fatalf("Error parsing YAML: %v", err)
	// }

	// // Dependency Injection
	// processor := &ConsoleProcessor{}
	// for _, character := range chContainer.Characters {
	// 	processor.Process(character)
	// }

	// currentScene := &scene.SceneContainer{}
	// previousScene := &scene.SceneContainer{}

	// processor := scene.NewSceneProcessor(currentScene)
	// prevProcessor := scene.NewSceneProcessor(previousScene)
	// ctx:= context.Background()

	a := app.New()
	w := a.NewWindow("Sample App")
	w.Resize(fyne.NewSize(1366, 768))

	stack := container.NewStack()
	screenFactories := widgets.ScreenFactory()
	manager := injection.NewStackScreenManager(stack, screenFactories)
	manager.SwitchScreen("login")

	w.SetContent(stack)
	w.ShowAndRun()
}


