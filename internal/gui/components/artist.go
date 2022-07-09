package gui_components

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func Artist(main *widgets.QWidget) {
	artistContainer := widgets.NewQWidget(nil, 0)
	artistContainer.SetLayout(widgets.NewQHBoxLayout())

	artistLabel := widgets.NewQLabel2("Artist Name", nil, core.Qt__Widget)
	artistText := widgets.NewQLineEdit(nil)

	artistContainer.Layout().AddWidget(artistLabel)
	artistContainer.Layout().AddWidget(artistText)
	main.Layout().AddWidget(artistContainer)
}
