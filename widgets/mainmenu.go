package widgets

import (
	"game/injection"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainMenuWidget(manager injection.ScreenMan) fyne.CanvasObject {
	img:= canvas.NewImageFromFile("./assets/bg.jpg")
	img.FillMode = canvas.ImageFillStretch

	// Create buttons
	continueButton := widget.NewButton("Continue", func() {})
	newGameButton := widget.NewButton("New Game", func() {})
	loadGameButton := widget.NewButton("Load Game", func() {})
	settingsButton := widget.NewButton("Settings", func() {})
	quitGameButton := widget.NewButton("Quit Game", func() {})

	// Create a horizontal layout for the buttons
	composed := container.New(layout.NewHBoxLayout(), continueButton, newGameButton, loadGameButton, settingsButton, quitGameButton)
	spacer := container.New(layout.NewGridWrapLayout(fyne.NewSize(200,100)),layout.NewSpacer())
	// Use a vertical box layout with a spacer above the buttons to push them to the bottom
	bottomLayout := container.New(layout.NewVBoxLayout(),layout.NewSpacer(),composed,spacer)
	ctx := container.New(layout.NewHBoxLayout(),layout.NewSpacer(), bottomLayout,layout.NewSpacer())
	stack := container.NewStack(img, ctx)

	return stack
}
