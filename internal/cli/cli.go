package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "ripntag [command]",
	Short: "Ripntag allows users to tag riped audio files with metadata",
	Long: "Allows for files ripped from a music CD to be tagged with \n" +
		"accurate metadata, also provides conversion from WAV.",
}

// Execute calls undelying 'Execute' function on the cobra command
func Execute() error {
	return rootCmd.Execute()
}
