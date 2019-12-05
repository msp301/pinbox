package pinbox

import (
	"github.com/BurntSushi/toml"
)

// Config holds any settings used to define the behaviour of the service.
// For example to set the port to host the API on etc.
type Config struct {
	Maildir         string
	TLS             bool
	CertificateFile string `toml:"certificate_file"`
	CertificateKey  string `toml:"certificate_key"`
	Port            int
	Inbox           string
	Bundle          []string
	Hidden          []string
}

// ReadConfigFile loads the configuration file from disk.
// Returns a Config object.
func ReadConfigFile(path string) (Config, error) {
	var config = Config{}

	_, err := toml.DecodeFile(path, &config)

	return config, err
}
