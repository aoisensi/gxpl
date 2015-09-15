package main

import (
	"os"

	"github.com/google/gxui"
)

func fileItem(theme gxui.Theme, file os.FileInfo) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	name := theme.CreateLabel()
	name.SetText(file.Name())
	layout.AddChild(name)
	return layout
}
