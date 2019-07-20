package pinbox

import (
	"testing"
)

func TestReadingConfig(t *testing.T) {
	filePath := "example/config.toml"
	config, err := ReadConfigFile(filePath)

	if err != nil {
		t.Fatalf("Failed to read config file '%s': %s", filePath, err)
	}

	if config.Port != 8000 {
		t.Fatalf("Unexpected port number read from config file '%s'", filePath)
	}
}

func TestReadingNonExistantConfig(t *testing.T) {
	filePath := "doesnotexist"
	_, err := ReadConfigFile(filePath)

	if err == nil {
		t.Fatalf("Expected error")
	}
}
