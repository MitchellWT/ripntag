package ripntag

import (
	"os"
	"testing"

	"github.com/wtolson/go-taglib"
)

var testingDirOne = "testing_album_one/"
var testingDirTwo = "testing_album_two/"

func TestTagDiscRipLow(t *testing.T) {
	rootDir := "Low/"
	createTestDir(rootDir, testingDirOne)

	rel := BarcodeSearch("4943674277988", false)
	TagDiscRip(rel, rootDir)

	checkFile(rootDir+"1 - Speed Of Life.flac", t, "Speed Of Life", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 1, 2018)
	checkFile(rootDir+"2 - Breaking Glass.flac", t, "Breaking Glass", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 2, 2018)
	checkFile(rootDir+"3 - What In The World.flac", t, "What In The World", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 3, 2018)
	checkFile(rootDir+"4 - Sound And Vision.flac", t, "Sound And Vision", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 4, 2018)
	checkFile(rootDir+"5 - Always Crashing In The Same Car.flac", t, "Always Crashing In The Same Car", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 5, 2018)
	checkFile(rootDir+"6 - Be My Wife.flac", t, "Be My Wife", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 6, 2018)
	checkFile(rootDir+"7 - A New Career In A New Town.flac", t, "A New Career In A New Town", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 7, 2018)
	checkFile(rootDir+"8 - Warszawa.flac", t, "Warszawa", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 8, 2018)
	checkFile(rootDir+"9 - Art Decade.flac", t, "Art Decade", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 9, 2018)
	checkFile(rootDir+"10 - Weeping Wall.flac", t, "Weeping Wall", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 10, 2018)
	checkFile(rootDir+"11 - Subterraneans.flac", t, "Subterraneans", "Low", "David Bowie", "Electronic;Rock;Art Rock;Ambient", 11, 2018)

	os.RemoveAll(rootDir)
}

func TestTagDiscRipPrettyHateMachine(t *testing.T) {
	rootDir := "Pretty Hate Machine/"
	createTestDir(rootDir, testingDirOne)

	rel := BarcodeSearch("602527567730", false)
	TagDiscRip(rel, rootDir)

	checkFile(rootDir+"1 - Head Like A Hole.flac", t, "Head Like A Hole", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 1, 2010)
	checkFile(rootDir+"2 - Terrible Lie.flac", t, "Terrible Lie", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 2, 2010)
	checkFile(rootDir+"3 - Down In It.flac", t, "Down In It", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 3, 2010)
	checkFile(rootDir+"4 - Sanctified.flac", t, "Sanctified", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 4, 2010)
	checkFile(rootDir+"5 - Something I Can Never Have.flac", t, "Something I Can Never Have", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 5, 2010)
	checkFile(rootDir+"6 - Kinda I Want To.flac", t, "Kinda I Want To", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 6, 2010)
	checkFile(rootDir+"7 - Sin.flac", t, "Sin", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 7, 2010)
	checkFile(rootDir+"8 - That's What I Get.flac", t, "That's What I Get", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 8, 2010)
	checkFile(rootDir+"9 - The Only Time.flac", t, "The Only Time", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 9, 2010)
	checkFile(rootDir+"10 - Ringfinger.flac", t, "Ringfinger", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 10, 2010)
	checkFile(rootDir+"11 - Get Down Make Love.flac", t, "Get Down Make Love", "Pretty Hate Machine", "Nine Inch Nails", "Electronic;Industrial", 11, 2010)

	os.RemoveAll(rootDir)
}

func TestTagFileNameShroomeez(t *testing.T) {
	rootDir := "Shroomeez/"
	createTestDir(rootDir, testingDirTwo)

	os.Rename(rootDir+"Track 1.mp3", rootDir+"you WaNNa Stay.mp3")
	os.Rename(rootDir+"Track 2.mp3", rootDir+"LeFToverS   .mp3")
	os.Rename(rootDir+"Track 3.mp3", rootDir+"Back At It.mp3")
	os.Rename(rootDir+"Track 4.mp3", rootDir+"Ma Osim (Very cooool).mp3")

	rel := ArtistAlbumSearch("Infected Mushroom", "Shroomeez", false)
	TagFileName(rel, rootDir)

	checkFile(rootDir+"1 - You Wanna Stay.mp3", t, "You Wanna Stay", "Shroomeez", "Infected Mushroom", "Electronic;Psy-Trance", 1, 2021)
	checkFile(rootDir+"2 - Leftovers.mp3", t, "Leftovers", "Shroomeez", "Infected Mushroom", "Electronic;Psy-Trance", 2, 2021)
	checkFile(rootDir+"3 - Back At It.mp3", t, "Back At It", "Shroomeez", "Infected Mushroom", "Electronic;Psy-Trance", 3, 2021)
	checkFile(rootDir+"4 - Ma Osim.mp3", t, "Ma Osim", "Shroomeez", "Infected Mushroom", "Electronic;Psy-Trance", 4, 2021)

	os.RemoveAll(rootDir)
}

func TestTagFileNameSin(t *testing.T) {
	rootDir := "Sin/"
	createTestDir(rootDir, testingDirTwo)

	os.Rename(rootDir+"Track 1.mp3", rootDir+"Sin (Long).mp3")
	os.Rename(rootDir+"Track 2.mp3", rootDir+"Sin (Dub).mp3")
	os.Rename(rootDir+"Track 3.mp3", rootDir+"Get Down Make Love.mp3")
	os.Rename(rootDir+"Track 4.mp3", rootDir+"Sin (Short).mp3")

	rel := BarcodeSearch("016581261723", false)
	TagFileName(rel, rootDir)

	checkFile(rootDir+"1 - Sin Long.mp3", t, "Sin Long", "Sin (Long, Dub & Short)", "Nine Inch Nails", "Electronic;Rock;Industrial", 1, 1990)
	checkFile(rootDir+"2 - Sin Dub.mp3", t, "Sin Dub", "Sin (Long, Dub & Short)", "Nine Inch Nails", "Electronic;Rock;Industrial", 2, 1990)
	checkFile(rootDir+"3 - Get Down Make Love.mp3", t, "Get Down Make Love", "Sin (Long, Dub & Short)", "Nine Inch Nails", "Electronic;Rock;Industrial", 3, 1990)
	checkFile(rootDir+"4 - Sin Short.mp3", t, "Sin Short", "Sin (Long, Dub & Short)", "Nine Inch Nails", "Electronic;Rock;Industrial", 4, 1990)

	os.RemoveAll(rootDir)
}

func TestTagFileNameストライク_ザ_ブラッド(t *testing.T) {
	rootDir := "ストライク・ザ・ブラッド/"
	createTestDir(rootDir, testingDirTwo)

	os.Rename(rootDir+"Track 1.mp3", rootDir+"ストライク・ザ・ブラッド.mp3")
	os.Rename(rootDir+"Track 2.mp3", rootDir+"ハンゲツトウゲ.mp3")
	os.Rename(rootDir+"Track 3.mp3", rootDir+"ストライク・ザ・ブラッド-Instrumenral-.mp3")
	os.Rename(rootDir+"Track 4.mp3", rootDir+"ハンゲツトウゲ-Instrumenral-.mp3")

	rel := ArtistAlbumSearch("Kisida Kyoudan & The Akebosi Rockets", "ストライク・ザ・ブラッド", false)
	TagFileName(rel, rootDir)

	checkFile(rootDir+"1 - ストライク・ザ・ブラッド.mp3", t, "ストライク・ザ・ブラッド", "ストライク・ザ・ブラッド", "Kisida Kyoudan & The Akebosi Rockets", "Rock", 1, 2013)
	checkFile(rootDir+"2 - ハンゲツトウゲ.mp3", t, "ハンゲツトウゲ", "ストライク・ザ・ブラッド", "Kisida Kyoudan & The Akebosi Rockets", "Rock", 2, 2013)
	checkFile(rootDir+"3 - ストライク・ザ・ブラッド-Instrumenral-.mp3", t, "ストライク・ザ・ブラッド-Instrumenral-", "ストライク・ザ・ブラッド", "Kisida Kyoudan & The Akebosi Rockets", "Rock", 3, 2013)
	checkFile(rootDir+"4 - ハンゲツトウゲ-Instrumenral-.mp3", t, "ハンゲツトウゲ-Instrumenral-", "ストライク・ザ・ブラッド", "Kisida Kyoudan & The Akebosi Rockets", "Rock", 4, 2013)

	os.RemoveAll(rootDir)
}

func createTestDir(rootDir string, testingDir string) {
	os.Mkdir(rootDir, 0775)
	files, err := os.ReadDir(testingDir)
	ErrorCheck(err)

	for _, file := range files {
		fileData, err := os.ReadFile(testingDir + file.Name())
		ErrorCheck(err)
		os.WriteFile(rootDir+file.Name(), fileData, 0644)
	}
}

func checkFile(file string, t *testing.T, title string, album string, artist string, genre string, track int, year int) {
	tagFile, err := taglib.Read(file)
	ErrorCheck(err)

	if tagFile.Title() != title {
		t.Errorf(`Error: tagFile.Title() equals %s, should equal "%s"`, tagFile.Title(), title)
	}
	if tagFile.Album() != album {
		t.Errorf(`Error: tagFile.Album() equals %s, should equal "%s"`, tagFile.Album(), album)
	}
	if tagFile.Artist() != artist {
		t.Errorf(`Error: tagFile.Artist() equals %s, should equal "%s"`, tagFile.Artist(), artist)
	}
	if tagFile.Genre() != genre {
		t.Errorf(`Error: tagFile.Genre() equals %s, should equal "%s"`, tagFile.Genre(), genre)
	}
	if tagFile.Track() != track {
		t.Errorf(`Error: tagFile.Track() equals %d, should equal %d`, tagFile.Track(), track)
	}
	if tagFile.Year() != year {
		t.Errorf(`Error: tagFile.Year() equals %d, should equal %d`, tagFile.Year(), year)
	}
}
