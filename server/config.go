// Copyright 2019 Martin Pritchard
//
// This file is part of Pinbox.
//
// Pinbox is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Pinbox is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Pinbox.  If not, see <https://www.gnu.org/licenses/>.

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
