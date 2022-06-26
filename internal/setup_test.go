package ripntag

import (
	"os"
	"testing"
)

func TestPreSetupCheck(t *testing.T) {
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
}
