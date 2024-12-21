package main

import ("fyne.io/fyne/v2")

type ScreenMan interface {
	SwitchScreen(string)
}

type StackScreenMan struct{
	container *fyne.Container
}

func NewStackScreenManager(container *fyne.Container) *StackScreenMan {
	return &StackScreenMan{container: container}
}

func (s *StackScreenMan) SwitchScreen(screen string) {
	switch screen {
	case "login":
		s.container.Objects = []fyne.CanvasObject{LoginWidget(s)}
	case "home":
		s.container.Objects = []fyne.CanvasObject{HomeWidget(s)}
	}
	s.container.Refresh()
}