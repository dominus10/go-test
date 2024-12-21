package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sample App")
	w.Resize(fyne.NewSize(1366, 768))

	stack := container.NewStack()
	manager := NewStackScreenManager(stack)
	manager.SwitchScreen("login")

	w.SetContent(stack)
	w.ShowAndRun()
}


