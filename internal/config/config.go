// Package config handles loading the netbox-cli JSON configuration file and
// resolving the NetBox URL and API token. The token is read from the
// NETBOX_TOKEN environment variable first, falling back to the config file.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const defaultConfigFile = ".netbox_cli.json"

// Config holds the values loaded from the JSON configuration file.
type Config struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

// Load reads and parses the JSON config file at path. If path is empty the
// default location (~/.netbox_cli.json) is used. A missing file is not an
// error; an empty Config is returned instead.
func Load(path string) (*Config, error) {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("could not determine home directory: %w", err)
		}
		path = filepath.Join(home, defaultConfigFile)
	}

	data, err := os.ReadFile(path) //nolint:gosec // path is user-supplied config file location
	if errors.Is(err, os.ErrNotExist) {
		return &Config{}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading config file %s: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config file %s: %w", path, err)
	}

	return &cfg, nil
}

// ResolveToken returns the API token, preferring the NETBOX_TOKEN environment
// variable over the token field in the config file. An error is returned if
// no token is found in either location.
func (c *Config) ResolveToken() (string, error) {
	if token := os.Getenv("NETBOX_TOKEN"); token != "" {
		return token, nil
	}
	if c.Token != "" {
		return c.Token, nil
	}
	return "", errors.New(
		"no API token found: set the NETBOX_TOKEN environment variable " +
			"or add a \"token\" key to your config file",
	)
}

// ResolveURL returns the NetBox base URL from the config file. An error is
// returned if the url field is empty.
func (c *Config) ResolveURL() (string, error) {
	if c.URL == "" {
		return "", errors.New("no NetBox URL found: add a \"url\" key to your config file")
	}
	return c.URL, nil
}
