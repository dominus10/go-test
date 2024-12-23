package injection

import (
	"fyne.io/fyne/v2"
)

type ScreenFactory func(manager ScreenMan) fyne.CanvasObject

type ScreenMan interface {
	SwitchScreen(string)
}

type StackScreenMan struct {
	container    *fyne.Container
	screenFactories map[string]ScreenFactory
}

func NewStackScreenManager(container *fyne.Container, factories map[string]ScreenFactory) *StackScreenMan {
	return &StackScreenMan{
		container: container,
		screenFactories: factories,
	}
}

func (s *StackScreenMan) SwitchScreen(screen string) {
	if factory, exists := s.screenFactories[screen]; exists {
		s.container.Objects = []fyne.CanvasObject{factory(s)}
		s.container.Refresh()
	}
}
