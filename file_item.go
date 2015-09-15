package main

import (
	"os"

	"github.com/aoisensi/gxui-material-icon"
	"github.com/google/gxui"
)

func fileItem(theme gxui.Theme, file os.FileInfo) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	
	icon := materialicon.CreateIcon(theme, materialicon.IconFolder, 12)
	layout.AddChild(icon)

	name := theme.CreateLabel()
	name.SetText(file.Name())
	layout.AddChild(name)
	return layout
}
