package gui_components

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func Barcode(main *widgets.QWidget) {
	barcodeContainer := widgets.NewQWidget(nil, 0)
	barcodeContainer.SetLayout(widgets.NewQHBoxLayout())

	barcodeLabel := widgets.NewQLabel2("Album Barcode", nil, core.Qt__Widget)
	barcodeText := widgets.NewQLineEdit(nil)

	barcodeContainer.Layout().AddWidget(barcodeLabel)
	barcodeContainer.Layout().AddWidget(barcodeText)
	main.Layout().AddWidget(barcodeContainer)
}
