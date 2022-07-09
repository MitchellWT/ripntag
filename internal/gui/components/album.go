package gui_components

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func Album(main *widgets.QWidget) {
	albumContainer := widgets.NewQWidget(nil, 0)
	albumContainer.SetLayout(widgets.NewQHBoxLayout())

	albumLabel := widgets.NewQLabel2("Album Name", nil, core.Qt__Widget)
	albumText := widgets.NewQLineEdit(nil)

	albumContainer.Layout().AddWidget(albumLabel)
	albumContainer.Layout().AddWidget(albumText)
	main.Layout().AddWidget(albumContainer)
}
