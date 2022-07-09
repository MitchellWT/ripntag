package gui_components

import (
	"github.com/therecipe/qt/widgets"
)

func FormFooter(main *widgets.QWidget) {
	formFooterContainer := widgets.NewQWidget(nil, 0)
	formFooterContainer.SetLayout(widgets.NewQHBoxLayout())

	formSubmitButton := widgets.NewQPushButton2("BEGIN RIP 'N TAG", nil)

	formFooterContainer.Layout().AddWidget(formSubmitButton)
	main.Layout().AddWidget(formFooterContainer)
}
