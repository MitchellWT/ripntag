package ripntag

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var ConfigDir = getConfigDir()

func getConfigDir() string {
	homeDir, err := os.UserHomeDir()
	ErrorCheck(err)

	return homeDir + "/.config/ripntag/"
}

// Checks If setup must be ran
func preSetupCheck() bool {
	_, err := os.Stat(ConfigDir)
	if err != nil && os.IsNotExist(err) {
		return true
	}

	_, err = os.Stat(ConfigDir + "token")
	if err != nil && os.IsNotExist(err) {
		return true
	}

	return false
}

func createDir() {
	err := os.MkdirAll(ConfigDir, 0755)
	ErrorCheck(err)
}

// createToken asks user for token and stores it in a local file (plain text)
// TODO: need to add consideration when token is incorrect (test auth) and ask for token again
func createToken() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("To use ripntag we need to access the discogs API. This requires\n" +
		"a personal access token for your account. To generate this go to:\n\n" +
		"Dashboard > Settings > Developers\n\n" +
		"Please enter the generate token: ")
	lineByte, _, err := reader.ReadLine()
	token := bytes.TrimSpace(lineByte)
	ErrorCheck(err)

	err = os.WriteFile(ConfigDir+"token", token, 0644)
	ErrorCheck(err)
}

func Setup() bool {
	if !preSetupCheck() {
		return false
	}

	createDir()
	createToken()
	return true
}
