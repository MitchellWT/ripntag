package ripntag

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// WavToFlac converts a folder of wav files to flac using ffmpeg
func WavToFlac(rootDir string) string {
	nameSplit := strings.Split(rootDir, "/")
	destDir := strings.Join(append(nameSplit[:len(nameSplit)-2], nameSplit[len(nameSplit)-2]+" - Tagged"), "/")
	err := os.MkdirAll(destDir, 0755)
	ErrorCheck(err)
	files, err := os.ReadDir(rootDir)
	ErrorCheck(err)

	for _, file := range files {
		inputFile := fmt.Sprintf("%s/%s", rootDir, file.Name())
		outputFile := fmt.Sprintf("%s/%s.flac", destDir, strings.Split(file.Name(), ".")[0])
		cmd := exec.Command("ffmpeg", "-i", inputFile, "-af", "aformat=s32:176000", outputFile)
		err := cmd.Run()
		ErrorCheck(err)
	}
	return destDir
}
