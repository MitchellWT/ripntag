package ripntag

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func WavToFlac(rootDir string) {
	nameSplit := strings.Split(rootDir, "/")
	destDir := strings.Join(append(nameSplit[:len(nameSplit)-1], nameSplit[len(nameSplit)-1]+" - Tagged"), "/")
	err := os.MkdirAll(destDir, 0755)
	ErrorCheck(err)
	files, err := os.ReadDir(rootDir)
	ErrorCheck(err)

	for _, file := range files {
		inputFile := fmt.Sprintf("%s/%s", rootDir, file.Name())
		outputFile := fmt.Sprintf("%s/%s.flac", destDir, strings.Split(file.Name(), ".")[0])
		fmt.Println(inputFile)
		fmt.Println(outputFile)
		cmd := exec.Command("ffmpeg", "-i", inputFile, "-af", "aformat=s32:176000", outputFile)
		err := cmd.Run()
		ErrorCheck(err)
	}
}
