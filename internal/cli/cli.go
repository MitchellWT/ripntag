package cli

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	ripntag "gitlab.com/MitchellWT/ripntag/internal"
	enums "gitlab.com/MitchellWT/ripntag/internal/enums"
)

var rootCmd = &cobra.Command{
	Use:   "ripntag [album directory]",
	Short: "Ripntag allows users to tag riped audio files with metadata",
	Long: "Allows for files ripped from a music CD to be tagged with \n" +
		"accurate metadata, also provides conversion from WAV.",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rootCommand(cmd, args)
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// checkDir performs some basic checks to ensure that the provided dir path
// is correct, these included checking last rune and If the dir exists
func checkDir(inputDir string) string {
	if rune(inputDir[len(inputDir)-1]) != '/' {
		inputDir = inputDir + "/"
	}
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	return inputDir
}

// relativeToAbsolute converts a provided relative path to an absolute path
func relativeToAbsolute(inputDir string) string {
	outDir, err := filepath.Abs(inputDir)
	ripntag.ErrorCheck(err)

	return outDir + "/"
}

func emptyCheck(s string) string {
	if len(s) < 1 {
		log.Fatal("input is empty")
	}
	return s
}

func rootCommand(cmd *cobra.Command, args []string) {
	nonInter := cmd.Flag("non-interactive").Changed
	seaMethod, err := enums.ToSearchMethod(cmd.Flag("search").Value.String())
	ripntag.ErrorCheck(err)
	tagType, err := enums.ToTagType(cmd.Flag("type").Value.String())
	ripntag.ErrorCheck(err)
	albumDir := relativeToAbsolute(checkDir(args[0]))

	switch seaMethod {
	case enums.Barcode:
		barcode := emptyCheck(cmd.Flag("barcode").Value.String())
		rel := ripntag.BarcodeSearch(barcode, nonInter)
		switch tagType {
		case enums.Rip:
			ripntag.TagDiscRip(rel, albumDir)
		case enums.FileName:
			ripntag.TagFileName(rel, albumDir)
		}
	case enums.ArtistAlbum:
		artist := emptyCheck(cmd.Flag("artist").Value.String())
		album := emptyCheck(cmd.Flag("album").Value.String())
		rel := ripntag.ArtistAlbumSearch(artist, album, nonInter)
		switch tagType {
		case enums.Rip:
			ripntag.TagDiscRip(rel, albumDir)
		case enums.FileName:
			ripntag.TagFileName(rel, albumDir)
		}
	}
}

func init() {
	rootCmd.Flags().BoolP("non-interactive", "n", false, "stops interactive selection, this forces the first \n"+
		"album search result to be used")
	rootCmd.Flags().StringP("search", "s", "barcode", "specifies the search method used on the discogs API, this \n"+
		"will change the flag requirements, the following search methods are supported: \n"+
		"- artist-album \n"+
		"- barcode")
	rootCmd.Flags().StringP("type", "t", "rip", "specifies the tag type of the album directory, the tag type can \n"+
		"be one of the following: \n"+
		"- file-name \n"+
		"- rip")
	rootCmd.Flags().StringP("barcode", "b", "", "barecode used for searching the discogs API")
	rootCmd.Flags().StringP("artist", "a", "", "artist used for searching the discogs API (album flag also required)")
	rootCmd.Flags().StringP("album", "l", "", "album used for searching the discogs API (artist flag also required)")
}

// Execute calls undelying 'Execute' function on the cobra command
func Execute() error {
	return rootCmd.Execute()
}
