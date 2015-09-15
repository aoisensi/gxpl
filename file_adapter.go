package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

func newFileAdapter() *fileAdapter {
	return new(fileAdapter)
}

type fileAdapter struct {
	gxui.AdapterBase
	fs []os.FileInfo
}

func (a *fileAdapter) Count() int {
	return len(a.fs)
}

func (a *fileAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.fs[index]
}

func (a *fileAdapter) ItemIndex(item gxui.AdapterItem) int {
	for i, f := range a.fs {
		if f == item {
			return i
		}
	}
	return -1
}

func (a *fileAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: 500, H: 20}
}

func (a *fileAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	return fileItem(theme, a.fs[index])
}

func (a *fileAdapter) Reload() error {
	dirname, err := filepath.Abs(".")

	if err != nil {
		return err
	}

	fs, err := ioutil.ReadDir(dirname)

	if err != nil {
		return err
	}
	a.fs = fs
	a.DataChanged(false)
	return nil
}
