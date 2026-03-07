package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kirtis/netbox-cli/internal/config"
)

// writeConfig writes content to a temp file and returns its path.
func writeConfig(t *testing.T, content string) string {
	t.Helper()
	f := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(f, []byte(content), 0o600); err != nil {
		t.Fatalf("writeConfig: %v", err)
	}
	return f
}

func TestLoad(t *testing.T) {
	t.Run("valid JSON", func(t *testing.T) {
		path := writeConfig(t, `{"url":"https://netbox.example.com","token":"abc123"}`)
		cfg, err := config.Load(path)
		if err != nil {
			t.Fatalf("got error %v, want nil", err)
		}
		if cfg.URL != "https://netbox.example.com" {
			t.Errorf("got URL %q, want %q", cfg.URL, "https://netbox.example.com")
		}
		if cfg.Token != "abc123" {
			t.Errorf("got Token %q, want %q", cfg.Token, "abc123")
		}
	})

	t.Run("missing file returns empty config", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "does-not-exist.json")
		cfg, err := config.Load(path)
		if err != nil {
			t.Fatalf("got error %v, want nil", err)
		}
		if cfg.URL != "" || cfg.Token != "" {
			t.Errorf("got non-empty config %+v, want empty", cfg)
		}
	})

	t.Run("invalid JSON returns error", func(t *testing.T) {
		path := writeConfig(t, `{not valid json}`)
		_, err := config.Load(path)
		if err == nil {
			t.Fatal("got nil error, want parse error")
		}
	})
}

func TestResolveToken(t *testing.T) {
	t.Run("env var takes precedence over config", func(t *testing.T) {
		t.Setenv("NETBOX_TOKEN", "envtoken")
		cfg := &config.Config{Token: "filetoken"}
		got, err := cfg.ResolveToken()
		if err != nil {
			t.Fatalf("got error %v, want nil", err)
		}
		if got != "envtoken" {
			t.Errorf("got %q, want %q", got, "envtoken")
		}
	})

	t.Run("falls back to config file token", func(t *testing.T) {
		t.Setenv("NETBOX_TOKEN", "")
		cfg := &config.Config{Token: "filetoken"}
		got, err := cfg.ResolveToken()
		if err != nil {
			t.Fatalf("got error %v, want nil", err)
		}
		if got != "filetoken" {
			t.Errorf("got %q, want %q", got, "filetoken")
		}
	})

	t.Run("error when no token available", func(t *testing.T) {
		t.Setenv("NETBOX_TOKEN", "")
		cfg := &config.Config{}
		_, err := cfg.ResolveToken()
		if err == nil {
			t.Fatal("got nil error, want error")
		}
	})
}

func TestResolveURL(t *testing.T) {
	t.Run("returns URL from config", func(t *testing.T) {
		cfg := &config.Config{URL: "https://netbox.example.com"}
		got, err := cfg.ResolveURL()
		if err != nil {
			t.Fatalf("got error %v, want nil", err)
		}
		if got != "https://netbox.example.com" {
			t.Errorf("got %q, want %q", got, "https://netbox.example.com")
		}
	})

	t.Run("error when URL is empty", func(t *testing.T) {
		cfg := &config.Config{}
		_, err := cfg.ResolveURL()
		if err == nil {
			t.Fatal("got nil error, want error")
		}
	})
}
