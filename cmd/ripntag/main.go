package main

import (
	ripntag "gitlab.com/MitchellWT/ripntag/internal"
	"gitlab.com/MitchellWT/ripntag/internal/cli"
)

func main() {
	ripntag.Setup()
	cli.Execute()
}
