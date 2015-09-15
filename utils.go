package main

import (
	"github.com/google/gxui"
)

func hLayout(theme gxui.Theme, controls ...gxui.Control) gxui.Control {
	l := theme.CreateLinearLayout()
	l.SetDirection(gxui.TopToBottom)
	for _, c := range controls {
		l.AddChild(c)
	}
	return l
}

func vLayout(theme gxui.Theme, controls ...gxui.Control) gxui.Control {
	l := theme.CreateLinearLayout()
	l.SetDirection(gxui.LeftToRight)
	for _, c := range controls {
		l.AddChild(c)
	}
	return l
}
