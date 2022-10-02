package ripntag

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	downloadCover(rel, rootDir)

	for _, file := range files {
		nameSplit := strings.Split(strings.ToLower(file.Name()), ".")
		fileExt := nameSplit[len(nameSplit)-1]
		for _, ext := range musicExtensions {
			if fileExt == ext {
				re := regexp.MustCompile(`\d+`)
				trackPos, err := strconv.Atoi(re.FindString(file.Name()))
				ErrorCheck(err)

				track := rel.Tracklist[trackPos-1]
				tagFile(rel, track, rootDir, file.Name(), trackPos, fileExt)
				break
			}
		}
	}
}

func TagFileName(rel *discogs.Release, rootDir string) {
	files, err := os.ReadDir(rootDir)
	ErrorCheck(err)
	downloadCover(rel, rootDir)

	for _, file := range files {
		// Checks If file is an audio file from our defined list
		nameSplit := strings.Split(strings.ToLower(file.Name()), ".")
		fileExt := nameSplit[len(nameSplit)-1]
		for _, ext := range musicExtensions {
			if fileExt == ext {
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
				break
			}
		}
	}
}

func matchFileName(rel *discogs.Release, fileName string) (discogs.Track, int) {
	track := discogs.Track{}
	max := 0
	trackPos := -1
	for index, relTrack := range rel.Tracklist {
		re := regexp.MustCompile(prepForMatch(relTrack.Title))
		match := re.FindString(prepForMatch(fileName))
		if len(match) > max {
			max = len(match)
			track = relTrack
			trackPos = index + 1
		}
	}
	return track, trackPos
}

func tagFile(rel *discogs.Release, track discogs.Track, rootDir string,
	fileName string, trackPos int, fileExt string) {
	tagFile, err := taglib.Read(rootDir + fileName)
	ErrorCheck(err)

	tagFile.SetTitle(track.Title)
	tagFile.SetAlbum(rel.Title)
	tagFile.SetArtist(convertArtists(track.Artists, rel.Artists))
	tagFile.SetGenre(convertGenres(rel.Genres, rel.Styles))
	tagFile.SetTrack(trackPos)
	tagFile.SetYear(rel.Year)

	ErrorCheck(tagFile.Save())

	newName := fmt.Sprintf("%d - %s.%s", trackPos, track.Title, fileExt)
	os.Rename(rootDir+fileName, rootDir+newName)
	fmt.Println(fileName, ": Tagged and renamed to ", newName)
}

func prepForMatch(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "")
}

func convertArtists(trkArts []discogs.ArtistSource, relArts []discogs.ArtistSource) string {
	artStr := ""
	for _, artist := range trkArts {
		artStr += artist.Name + ";"
	}
	if len(artStr) > 0 {
		return removeLastRune(artStr)
	}
	for _, artist := range relArts {
		artStr += artist.Name + ";"
	}
	if len(artStr) > 0 {
		artStr = removeLastRune(artStr)
	}
	return artStr
}

func convertGenres(genres []string, styles []string) string {
	genStr := ""
	for _, genre := range genres {
		genStr += genre + ";"
	}
	for _, style := range styles {
		genStr += style + ";"
	}
	if len(genStr) > 0 {
		genStr = removeLastRune(genStr)
	}
	return genStr
}

func downloadCover(rel *discogs.Release, rootDir string) {
	url := rel.Images[0].ResourceURL
	urlSplit := strings.Split(url, ".")
	imageExt := urlSplit[len(urlSplit)-1]
	res, err := http.Get(url)
	ErrorCheck(err)

	defer res.Body.Close()
	imageData, err := ioutil.ReadAll(res.Body)
	ErrorCheck(err)
	err = os.WriteFile(rootDir+"cover."+imageExt, imageData, 0644)
	ErrorCheck(err)
}

func removeLastRune(input string) string {
	return input[:len(input)-1]
}
