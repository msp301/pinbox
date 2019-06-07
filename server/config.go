package pinbox

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Maildir string
	Port    int
	Inbox   string
	Bundle  []string
	Hidden  []string
}

func ReadConfigFile(path string) (Config, error) {
	var config = Config{}

	_, err := toml.DecodeFile(path, &config)

	return config, err
}
