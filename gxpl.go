package main

import (
	"os"
	"path/filepath"

	"github.com/aoisensi/gxui-material-icon"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	materialicon.SetDriver(driver)
	theme := dark.CreateTheme(driver)
	window := theme.CreateWindow(800, 600, "gxpl")
	list := theme.CreateList()
	adapter := newFileAdapter()
	list.SetAdapter(adapter)

	reloadButton := theme.CreateButton()
	reloadButton.SetText("Reload")
	reloadButton.OnClick(func(e gxui.MouseEvent) {
		adapter.Reload()
	})
	backButton := theme.CreateButton()
	backButton.SetText("Back")
	backButton.OnClick(func(e gxui.MouseEvent) {
		path, _ := os.Getwd()
		os.Chdir(filepath.Join(path, ".."))
		adapter.Reload()
	})

	window.AddChild(hLayout(theme, vLayout(theme, reloadButton, backButton), list))
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
