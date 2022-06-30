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
	// Can't be tagged!!!
	// "wav",
	// "wave",
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
		fileExt := nameSplit[len(nameSplit)-1]
		for _, ext := range musicExtensions {
			if fileExt == ext {
				exit = false
				break
			}
		}
		if exit {
			continue
		}
		re := regexp.MustCompile(`\d+`)
		trackPos, err := strconv.Atoi(re.FindString(file.Name()))
		ErrorCheck(err)

		track := rel.Tracklist[trackPos-1]
		tagFile(rel, track, rootDir, file.Name(), trackPos, fileExt)
	}
}

func TagFileName(rel *discogs.Release, rootDir string) {
	files, err := os.ReadDir(rootDir)
	ErrorCheck(err)

	for _, file := range files {
		// Checks If file is an audio file from our defined list
		exit := true
		nameSplit := strings.Split(strings.ToLower(file.Name()), ".")
		fileExt := nameSplit[len(nameSplit)-1]
		for _, ext := range musicExtensions {
			if fileExt == ext {
				exit = false
			}
		}
		if exit {
			continue
		}
		track, trackPos := matchFileName(rel, file.Name())
		if trackPos == -1 && strings.ContainsAny(strings.ToLower(file.Name()), "()[]") {
			replacer := strings.NewReplacer(
				"(", "",
				")", "",
				"[", "",
				"]", "",
			)
			fixedFileName := replacer.Replace(strings.ToLower(file.Name()))
			track, trackPos = matchFileName(rel, fixedFileName)
		}
		if trackPos == -1 {
			log.Fatalf("Can't find track for %s!", file.Name())
		}
		tagFile(rel, track, rootDir, file.Name(), trackPos, fileExt)
	}
}

func matchFileName(rel *discogs.Release, fileName string) (discogs.Track, int) {
	var track discogs.Track
	max := 0
	trackPosition := -1
	for index, relTrack := range rel.Tracklist {
		re := regexp.MustCompile(prepForMatch(relTrack.Title))
		match := re.FindString(prepForMatch(fileName))
		if len(match) > max {
			max = len(match)
			track = relTrack
			trackPosition = index + 1
		}
	}
	return track, trackPosition
}

func tagFile(rel *discogs.Release, track discogs.Track, rootDir string,
	fileName string, trackPos int, fileExt string) {
	tagFile, err := taglib.Read(rootDir + fileName)
	ErrorCheck(err)

	tagFile.SetTitle(track.Title)
	tagFile.SetAlbum(rel.Title)
	tagFile.SetArtist(convertArtists(track.Artists, rel.Artists))
	tagFile.SetGenre(convertGenres(rel.Genres))
	tagFile.SetTrack(trackPos)
	tagFile.SetYear(rel.Year)

	ErrorCheck(tagFile.Save())

	newName := fmt.Sprintf("%d - %s.%s", trackPos, track.Title, fileExt)
	os.Rename(rootDir+fileName, rootDir+newName)
}

func prepForMatch(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "")
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
