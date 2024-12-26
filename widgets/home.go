package widgets

import (
	"game/injection"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func HomeWidget(manager injection.ScreenMan) fyne.CanvasObject {
	// playerBar := container.New(layout.NewVBoxLayout())
	// optionsBar:= container.New(layout.NewVBoxLayout(),canvas.NewLine(color.White),widget.NewButton("123",func() {
	// 	manager.SwitchScreen("mainmenu")
	// }),canvas.NewLine(color.White))

	leftContainer := CreateCustomContainer(widget.NewLabel("a"))
	rightContainer := CreateCustomContainer(widget.NewLabel("a"))

	ctx:= container.New(layout.NewHBoxLayout(),leftContainer,rightContainer)

	return ctx
}

func CreateCustomContainer(components fyne.CanvasObject) *fyne.Container {
	// Create the border rectangle
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeWidth = 1
	border.StrokeColor = color.White
	border.CornerRadius = 8

	// Create the inner content
	innerContent := container.New(
		layout.NewVBoxLayout(),
		container.New(
			layout.NewGridWrapLayout(fyne.NewSize(200, 0)),
			layout.NewSpacer(),
		),
		components,
	)

	// Stack the border and padded content
	content := container.NewStack(
		border,
		container.NewPadded(innerContent),
	)

	// Add the content to a bordered container
	return content
}