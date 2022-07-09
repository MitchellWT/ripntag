package gui

import (
	"os"

	"github.com/therecipe/qt/widgets"
	gui_components "gitlab.com/MitchellWT/ripntag/internal/gui/components"
)

func Execute() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Rip 'N Tag")

	main := widgets.NewQWidget(nil, 0)
	main.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(main)

	gui_components.Form(main)

	window.Show()
	app.Exec()
}
