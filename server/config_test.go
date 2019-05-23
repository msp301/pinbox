package pinbox

import "testing"

func TestReadingConfig(t *testing.T) {
	filePath := "example/config.toml"
	_, err := readConfigFile(filePath)

	if err != nil {
		t.Fatalf("Failed to read config file '%s': %s", filePath, err)
	}
}

func TestReadingNonExistantConfig(t *testing.T) {
	filePath := "doesnotexist"
	_, err := readConfigFile(filePath)

	if err != nil {
		t.Fatalf("Expected error")
	}
}
