package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"gopkg.in/yaml.v3"
)

func main() {

	var chContainer CharactersContainer
	data, _:= os.ReadFile("./save.yaml")

	err := yaml.Unmarshal([]byte(data), &chContainer)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// Dependency Injection
	processor := &ConsoleProcessor{}
	for _, character := range chContainer.Characters {
		processor.Process(character)
	}


	a := app.New()
	w := a.NewWindow("Sample App")
	w.Resize(fyne.NewSize(1366, 768))

	stack := container.NewStack()
	manager := NewStackScreenManager(stack)
	manager.SwitchScreen("login")

	w.SetContent(stack)
	w.ShowAndRun()
}


