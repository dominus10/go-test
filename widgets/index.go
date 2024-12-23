package widgets

import "game/injection"

func ScreenFactory()map[string]injection.ScreenFactory {
	factory := map[string]injection.ScreenFactory{
		"login": LoginWidget,
		"home":  HomeWidget,
	}

	return factory
}