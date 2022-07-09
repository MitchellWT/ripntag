package gui_components

import "github.com/therecipe/qt/widgets"

func Form(main *widgets.QWidget) {
	AlbumFolder(main)
	NonInteractive(main)
	SearchMethod(main)
	TagType(main)
	// Need to use search method for conditional rendering
	Barcode(main)
	Album(main)
	Artist(main)
	FormFooter(main)
}
