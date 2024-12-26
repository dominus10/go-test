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
	size:= fyne.NewSize(200,48)
	
	img:= canvas.NewImageFromFile("./assets/bg.jpg")
	img.FillMode = canvas.ImageFillStretch

	// Create buttons
	continueButton := widget.NewButton("Continue", func() {})
	continueButton.Disable()
	continueButtonContainer:= container.New(
		layout.NewGridWrapLayout(size),
		continueButton,
	)
	
	newGameButton := container.New(
		layout.NewGridWrapLayout(size),
		widget.NewButton("New Game", func() {
			manager.SwitchScreen("home")
		}),
	)
	loadGameButton := widget.NewButton("Load Game", func() {})
	loadGameButton.Disable()
	loadButtonContainer := container.New(
		layout.NewGridWrapLayout(size),
		loadGameButton,
	)
	settingsButton := container.New(
		layout.NewGridWrapLayout(size),
		widget.NewButton("Settings", func() {}),
	)
	quitGameButton := container.New(
		layout.NewGridWrapLayout(size),
		widget.NewButton("Quit Game", func() {}),
	)

	// Create a horizontal layout for the buttons
	composed := container.New(layout.NewVBoxLayout(), continueButtonContainer, newGameButton, loadButtonContainer, settingsButton, quitGameButton)
	spacer := container.New(layout.NewGridWrapLayout(fyne.NewSize(100,100)),layout.NewSpacer())
	// Use a vertical box layout with a spacer above the buttons to push them to the bottom
	bottomLayout := container.New(layout.NewVBoxLayout(),layout.NewSpacer(),composed,spacer)
	ctx := container.New(layout.NewHBoxLayout(),layout.NewSpacer(), bottomLayout,layout.NewSpacer())
	stack := container.NewStack(img, ctx)

	return stack
}
