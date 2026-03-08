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
	tests := []struct {
		name      string
		envToken  string
		cfgToken  string
		wantToken string
		wantErr   bool
	}{
		{
			name:      "env var takes precedence over config",
			envToken:  "envtoken",
			cfgToken:  "filetoken",
			wantToken: "envtoken",
		},
		{
			name:      "falls back to config file token",
			envToken:  "",
			cfgToken:  "filetoken",
			wantToken: "filetoken",
		},
		{
			name:     "error when no token available",
			envToken: "",
			cfgToken: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("NETBOX_TOKEN", tt.envToken)
			cfg := &config.Config{Token: tt.cfgToken}
			got, err := cfg.ResolveToken()
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.wantToken {
				t.Errorf("got %q, want %q", got, tt.wantToken)
			}
		})
	}
}

func TestResolveURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantURL string
		wantErr bool
	}{
		{
			name:    "returns URL from config",
			url:     "https://netbox.example.com",
			wantURL: "https://netbox.example.com",
		},
		{
			name:    "error when URL is empty",
			url:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{URL: tt.url}
			got, err := cfg.ResolveURL()
			if (err != nil) != tt.wantErr {
				t.Fatalf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.wantURL {
				t.Errorf("got %q, want %q", got, tt.wantURL)
			}
		})
	}
}
