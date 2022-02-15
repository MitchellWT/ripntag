package ripntag

import (
	"fmt"

	"github.com/irlndts/go-discogs"
)

func AlbumArtistSearch(album string, artist string) {
	seaReq := discogs.SearchRequest{
		ReleaseTitle: album,
		Artist:       artist,
	}
	seaRes := runSearch(seaReq)

	for _, res := range seaRes {
		// STUB TODO: may need to add interactive selection?
		fmt.Println(res.Title)
		fmt.Println(res.ResourceURL)
		fmt.Println()
	}
}

func BarcodeSearch(barcode string) {
	seaReq := discogs.SearchRequest{
		Barcode: barcode,
	}
	seaRes := runSearch(seaReq)

	for _, res := range seaRes {
		// STUB TODO: may need to add interactive selection?
		fmt.Println(res.Title)
		fmt.Println(res.ResourceURL)
		fmt.Println()
	}
}

func runSearch(seaReq discogs.SearchRequest) []discogs.Result {
	client, err := discogs.New(&discogs.Options{
		UserAgent: "ripntag",
		Token:     "TODO: needs to be pulled from a file on the machine. Maybe, ~/.config/ripntag/token",
	})
	ErrorCheck(err)

	sea, err := client.Search(seaReq)
	ErrorCheck(err)

	return sea.Results
}
