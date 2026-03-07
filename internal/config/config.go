package config

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
)

const defaultConfigFile = ".netbox_cli.json"

type Config struct {
  URL   string `json:"url"`
  Token string `json:"token"`
}

func Load(path string) (*Config, error) {
  if path == "" {
    home, err := os.UserHomeDir()
    if err != nil {
      return nil, fmt.Errorf("could not determine home directory: %w", err)
    }
    path = filepath.Join(home, defaultConfigFile)
  }

  data, err := os.ReadFile(path)
  if err != nil {
    if os.IsNotExist(err) {
      return &Config{}, nil
    }
    return nil, fmt.Errorf("reading config file %s: %w", path, err)
  }

  var cfg Config
  if err := json.Unmarshal(data, &cfg); err != nil {
    return nil, fmt.Errorf("parsing config file %s: %w", path, err)
  }

  return &cfg, nil
}

func (c *Config) ResolveToken() (string, error) {
  if token := os.Getenv("NETBOX_TOKEN"); token != "" {
    return token, nil
  }
  if c.Token != "" {
    return c.Token, nil
  }
  return "", fmt.Errorf(
    "no API token found: set the NETBOX_TOKEN environment variable " +
      "or add a \"token\" key to your config file",
  )
}

func (c *Config) ResolveURL() (string, error) {
  if c.URL == "" {
    return "", fmt.Errorf(
      "no NetBox URL found: add a \"url\" key to your config file",
    )
  }
  return c.URL, nil
}
