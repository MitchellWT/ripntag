package ripntag

import (
	"os"
	"strings"
	"testing"
)

var testing_album_three = "testing_album_three/"

func TestWavToFlac(t *testing.T) {
	rootDir := "Death Magnetic/"
	createTestDir(rootDir, testing_album_three)
	destDir := WavToFlac(rootDir)
	files, err := os.ReadDir(destDir)
	ErrorCheck(err)

	for _, file := range files {
		nameSplit := strings.Split(file.Name(), ".")
		if nameSplit[len(nameSplit)-1] != "flac" {
			t.Errorf(`Error: file name extension was %s, should be flac`, nameSplit[len(nameSplit)-1])
		}
	}
	os.RemoveAll(rootDir)
	os.RemoveAll(destDir)
}
