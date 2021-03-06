package ripntag

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"
	"unicode"

	"github.com/irlndts/go-discogs"
)

var client discogs.Discogs

func getClient() discogs.Discogs {
	client, err := discogs.New(&discogs.Options{
		UserAgent: "ripntag",
		Token:     getToken(),
	})
	ErrorCheck(err)

	return client
}

func getToken() string {
	Setup()

	token, err := os.ReadFile(ConfigDir + "token")
	ErrorCheck(err)
	token = bytes.ReplaceAll(token, []byte("\n"), []byte(""))

	return string(token)
}

// BarcodeSearch searches the discogs database (using the provided barcode)
// and returns the release struct
func BarcodeSearch(barcode string, interactive bool) *discogs.Release {
	lazyLoadClient()
	searchReq := discogs.SearchRequest{
		Barcode: barcode,
		Type:    "release",
	}
	return runSearch(searchReq, interactive)
}

// ArtistAlbumSearch searches the discogs database (using the provided album
// and artist) and returns the release struct
func ArtistAlbumSearch(artist string, album string, interactive bool) *discogs.Release {
	lazyLoadClient()
	searchReq := discogs.SearchRequest{
		ReleaseTitle: album,
		Artist:       artist,
		Type:         "release",
	}
	return runSearch(searchReq, interactive)
}

func lazyLoadClient() {
	if client == nil {
		client = getClient()
	}
}

// runSearch is a helper function that runs the provided search request
func runSearch(seaReq discogs.SearchRequest, interactive bool) *discogs.Release {
	search, err := client.Search(seaReq)
	ErrorCheck(err)

	rel, err := client.Release(search.Results[0].ID)
	ErrorCheck(err)

	if interactive {
		rel = interactiveSelection(search.Results)
	}
	return rel
}

// interactiveSelection is a helper function that provides an interactive
// selection interface for picking releases from search results
func interactiveSelection(resArr []discogs.Result) *discogs.Release {
	var selRel *discogs.Release
	reader := bufio.NewReader(os.Stdin)

forLoop:
	for _, res := range resArr {
		rel, err := client.Release(res.ID)
		ErrorCheck(err)
		// Generate and present terminal prompt
		termPrompt := fmt.Sprintf("\nTitle:\t%s\n", rel.Title)
		termPrompt += fmt.Sprintf("Country:\t%s\n", rel.Country)
		termPrompt += fmt.Sprintf("Year:\t%d\n", rel.Year)
		termPrompt += "Artist(s):"
		for _, artist := range rel.Artists {
			termPrompt += fmt.Sprintf(" \t%s\n", artist.Name)
		}
		termPrompt += "TrackList:"
		for index, track := range rel.Tracklist {
			termPrompt += fmt.Sprintf(" \t%d) %s\n", index+1, track.Title)
		}
		termPrompt += "\n" + "Is this the album your looking for (y/N): "
		// Print termPrompt to screen with formatting
		tabWrite := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprint(tabWrite, termPrompt)
		tabWrite.Flush()
		// Clean input and collect first rune
		lineByte, _, err := reader.ReadLine()
		ErrorCheck(err)

		lineByte = bytes.TrimSpace(lineByte)
		ans := rune(lineByte[0])
		ans = unicode.ToUpper(ans)

		switch ans {
		case 'Y':
			selRel = rel
			break forLoop
		}
	}
	return selRel
}
