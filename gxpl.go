package main

import (
	"os"
	"path/filepath"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(800, 600, "gxpl")
	list := theme.CreateList()
	adapter := newFileAdapter()
	list.SetSize(math.Size{math.MaxInt, math.MaxInt})
	list.SetAdapter(adapter)
	list.OnItemClicked(func(e gxui.MouseEvent, item gxui.AdapterItem) {
		if e.Button != gxui.MouseButtonLeft {
			return
		}
		fi := item.(os.FileInfo)
		if fi.IsDir() {
			mvdir(fi.Name())
			adapter.Reload()
		}
	})

	reloadButton := theme.CreateButton()
	reloadButton.SetText("Reload")
	reloadButton.OnClick(func(e gxui.MouseEvent) {
		adapter.Reload()
	})
	backButton := theme.CreateButton()
	backButton.SetText("Back")
	backButton.OnClick(func(e gxui.MouseEvent) {
		mvdir("..")
		adapter.Reload()
	})
	layoutButton := theme.CreateLinearLayout()
	layoutButton.SetDirection(gxui.RightToLeft)
	layoutButton.AddChild(reloadButton)
	layoutButton.AddChild(backButton)
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	layout.SetSizeMode(gxui.Fill)
	layout.AddChild(layoutButton)
	layout.AddChild(list)
	window.AddChild(layout)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

func mvdir(dir string) {
	path, _ := os.Getwd()
	os.Chdir(filepath.Join(path, dir))
}
