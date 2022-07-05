package gui

import (
	"bytes"
	"fmt"
	"os"

	"github.com/webview/webview"
	ripntag "gitlab.com/MitchellWT/ripntag/internal"
)

func getHtml() []byte {
	index, err := os.ReadFile("./web/index.html")
	ripntag.ErrorCheck(err)
	style, err := os.ReadFile("./web/style.css")
	ripntag.ErrorCheck(err)
	interactions, err := os.ReadFile("./web/interactions.js")
	ripntag.ErrorCheck(err)
	// Is there a better way to do this?
	style = append([]byte("<style>"), style...)
	style = append(style, []byte("</style>")...)
	interactions = append([]byte("<script>"), interactions...)
	interactions = append(interactions, []byte("</script>")...)
	index = bytes.ReplaceAll(index, []byte("<style-inject/>"), style)
	return bytes.ReplaceAll(index, []byte("<script-inject/>"), interactions)
}

func Execute() {
	index := getHtml()
	webview := webview.New(true)
	webview.SetTitle("Rip 'N Tag")
	webview.SetHtml(string(index))
	fmt.Println(string(index))
	webview.Run()
	webview.Destroy()
}
