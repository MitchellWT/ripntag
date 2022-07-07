package gui

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func Execute() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.Show()
	app.Exec()
}
