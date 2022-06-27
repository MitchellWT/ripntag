package ripntag

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/irlndts/go-discogs"
	"github.com/wtolson/go-taglib"
)

var musicExtensions = []string{
	"mp3",
	"ogg",
	"spx",
	"mpc",
	"mp+",
	"mpp",
	"ape",
	"flac",
	"wv",
	"wvc",
	"tta",
	"wma",
	"m4a",
	"m4b",
	"m4p",
	"m4r",
	"m4v",
	"mp4",
	"wav",
	"wave",
	"aif",
	"aiff",
	"aifc",
	"opus",
}

func TagDiscRip(rel *discogs.Release, rootDir string) {
	files, err := os.ReadDir(rootDir)
	ErrorCheck(err)

	for _, file := range files {
		// Checks If file is an audio file from our defined list
		exit := true
		nameSplit := strings.Split(strings.ToLower(file.Name()), ".")
		for _, ext := range musicExtensions {
			if nameSplit[len(nameSplit)-1] == ext {
				exit = false
			}
		}
		if exit {
			continue
		}

		re := regexp.MustCompile(`\d+`)
		trackPosition, err := strconv.Atoi(re.FindString(file.Name()))
		ErrorCheck(err)

		track := rel.Tracklist[trackPosition-1]

		tagFile, err := taglib.Read(rootDir + file.Name())
		ErrorCheck(err)

		tagFile.SetTitle(track.Title)
		tagFile.SetAlbum(rel.Title)
		tagFile.SetArtist(convertArtists(track.Artists, rel.Artists))
		tagFile.SetGenre(convertGenres(rel.Genres))
		tagFile.SetTrack(trackPosition)
		tagFile.SetYear(rel.Year)

		ErrorCheck(tagFile.Save())

		newName := fmt.Sprintf("%d - %s.%s", trackPosition, track.Title, nameSplit[len(nameSplit)-1])
		os.Rename(rootDir+file.Name(), rootDir+newName)
	}
}

// TODO: Make more robust (does not work with 'instramentals' and 'remixes')
func TagFileName(rel *discogs.Release, rootDir string) {
	files, err := os.ReadDir(rootDir)
	ErrorCheck(err)

	for _, file := range files {
		// Checks If file is an audio file from our defined list
		exit := true
		nameSplit := strings.Split(strings.ToLower(file.Name()), ".")
		for _, ext := range musicExtensions {
			if nameSplit[len(nameSplit)-1] == ext {
				exit = false
			}
		}
		if exit {
			continue
		}

		trackPosition := -1
		var track discogs.Track
		for index, tra := range rel.Tracklist {
			re := regexp.MustCompile(`.+` + strings.ToLower(tra.Title) + `.+`)
			match := re.MatchString(strings.ToLower(file.Name()))

			if match {
				track = tra
				trackPosition = index
			}
		}
		if trackPosition == -1 {
			log.Fatalf("Can't find track for %s!", file.Name())
		}

		tagFile, err := taglib.Read(rootDir + file.Name())
		ErrorCheck(err)

		tagFile.SetTitle(track.Title)
		tagFile.SetAlbum(rel.Title)
		tagFile.SetArtist(convertArtists(track.Artists, rel.Artists))
		tagFile.SetGenre(convertGenres(rel.Genres))
		tagFile.SetTrack(trackPosition)
		tagFile.SetYear(rel.Year)

		ErrorCheck(tagFile.Save())

		newName := fmt.Sprintf("%d - %s.%s", trackPosition, track.Title, nameSplit[len(nameSplit)-1])
		os.Rename(rootDir+file.Name(), rootDir+newName)
	}
}

func convertArtists(trkArts []discogs.ArtistSource, relArts []discogs.ArtistSource) string {
	artStr := ""
	for _, artist := range trkArts {
		artStr += artist.Name + "/"
	}
	if len(artStr) > 0 {
		return removeLastRune(artStr)
	}
	for _, artist := range relArts {
		artStr += artist.Name + "/"
	}
	if len(artStr) > 0 {
		artStr = removeLastRune(artStr)
	}
	return artStr
}

func convertGenres(genres []string) string {
	genStr := ""
	for _, genre := range genres {
		genStr += genre + "/"
	}
	if len(genStr) > 0 {
		genStr = removeLastRune(genStr)
	}
	return genStr
}

func removeLastRune(input string) string {
	return input[:len(input)-1]
}
