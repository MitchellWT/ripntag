package gui_components

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func SearchMethod(main *widgets.QWidget) {
	searchMethodContainer := widgets.NewQWidget(nil, 0)
	searchMethodContainer.SetLayout(widgets.NewQHBoxLayout())

	searchMethodLabel := widgets.NewQLabel2("Search Method", nil, core.Qt__Widget)

	searchMethodComboBox := widgets.NewQComboBox(nil)
	searchMethodComboBox.AddItems([]string{"Barcode", "Artist Album Name"})

	searchMethodContainer.Layout().AddWidget(searchMethodLabel)
	searchMethodContainer.Layout().AddWidget(searchMethodComboBox)
	main.Layout().AddWidget(searchMethodContainer)
}
