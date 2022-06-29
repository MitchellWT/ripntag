package ripntag

import (
	"os"
	"testing"
)

func TestPreSetupCheck(t *testing.T) {
	token, _ := os.ReadFile(ConfigDir + "token")

	os.RemoveAll(ConfigDir)
	if !preSetupCheck() {
		t.Error("Error: preSetupCheck equals false, should equal true")
	}

	os.MkdirAll(ConfigDir, 0755)
	if !preSetupCheck() {
		t.Error("Error: preSetupCheck equals false, should equal true")
	}

	os.WriteFile(ConfigDir+"token", nil, 0755)
	if preSetupCheck() {
		t.Error("Error: preSetupCheck equals true, should equal false")
	}

	// Clean up
	os.RemoveAll(ConfigDir)
	// Re-create token If it previously existed
	if len(token) > 0 {
		os.MkdirAll(ConfigDir, 0755)
		err := os.WriteFile(ConfigDir+"token", token, 0644)
		ErrorCheck(err)
	}
}
