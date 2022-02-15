package main

import (
	ripntag "gitlab.com/MitchellWT/ripntag/internal"
	"gitlab.com/MitchellWT/ripntag/internal/cli"
)

func main() {
	cli.Execute()
	// Testing
	ripntag.AlbumArtistSearch("let's dance", "david bowie")
	ripntag.BarcodeSearch("5021456168484")
}
