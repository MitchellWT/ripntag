package gui_components

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func TagType(main *widgets.QWidget) {
	tagTypeContainer := widgets.NewQWidget(nil, 0)
	tagTypeContainer.SetLayout(widgets.NewQHBoxLayout())

	tagTypeLabel := widgets.NewQLabel2("Tag Type", nil, core.Qt__Widget)

	tagTypeComboBox := widgets.NewQComboBox(nil)
	tagTypeComboBox.AddItems([]string{"CD Rip", "File Name"})

	tagTypeContainer.Layout().AddWidget(tagTypeLabel)
	tagTypeContainer.Layout().AddWidget(tagTypeComboBox)
	main.Layout().AddWidget(tagTypeContainer)
}
