package gui_components

import "github.com/therecipe/qt/widgets"

func NonInteractive(main *widgets.QWidget) {
	nonInteractiveContainer := widgets.NewQWidget(nil, 0)
	nonInteractiveContainer.SetLayout(widgets.NewQHBoxLayout())

	nonInteractiveCheckbox := widgets.NewQCheckBox2("Non-interactive", nil)

	nonInteractiveContainer.Layout().AddWidget(nonInteractiveCheckbox)
	main.Layout().AddWidget(nonInteractiveContainer)
}
