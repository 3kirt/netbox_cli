package cmdutil

import (
  "encoding/json"
  "fmt"
  "os"
)

func OutputJSON(v any) error {
  enc := json.NewEncoder(os.Stdout)
  enc.SetIndent("", "  ")
  return enc.Encode(v)
}

func APIError(err error) error {
  return fmt.Errorf("netbox API error: %w", err)
}
