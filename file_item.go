package main

import (
	"os"

	"github.com/aoisensi/gxui-material-icon"
	"github.com/google/gxui"
)

func fileItem(theme gxui.Theme, file os.FileInfo) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)
	iconID := materialicon.IconInsertDriveFile
	if file.IsDir() {
		iconID = materialicon.IconFolderOpen
	}
	icon := materialicon.CreateIcon(theme, iconID, 12)
	layout.AddChild(icon)

	name := theme.CreateLabel()
	name.SetText(file.Name())
	layout.AddChild(name)
	return layout
}
