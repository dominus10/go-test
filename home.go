package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func HomeWidget(manager ScreenMan) fyne.CanvasObject {

	
	playerBar := container.New(layout.NewVBoxLayout())

	leftContainer := container.New(layout.NewVBoxLayout(),playerBar)
	ctx:= container.New(layout.NewHBoxLayout(),leftContainer)

	return ctx
}