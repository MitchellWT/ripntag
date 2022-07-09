package gui_components

import "github.com/therecipe/qt/widgets"

func AlbumFolder(main *widgets.QWidget) {
	albumFolderContainer := widgets.NewQWidget(nil, 0)
	albumFolderContainer.SetLayout(widgets.NewQHBoxLayout())

	albumFolderButton := widgets.NewQPushButton2("Select Album Folder", nil)

	albumFolderFinder := widgets.NewQFileDialog2(nil, "Select Album Folder", "", "")
	albumFolderFinder.SetFileMode(widgets.QFileDialog__DirectoryOnly)

	albumFolderText := widgets.NewQLineEdit(nil)
	albumFolderText.SetReadOnly(true)

	albumFolderButton.ConnectClicked(func(bool) { albumFolderFinder.Show() })
	albumFolderFinder.ConnectFileSelected(func(albumFolder string) {
		albumFolderText.SetText(albumFolder)
	})

	albumFolderContainer.Layout().AddWidget(albumFolderButton)
	albumFolderContainer.Layout().AddWidget(albumFolderText)
	main.Layout().AddWidget(albumFolderContainer)
}
