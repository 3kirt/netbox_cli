package cmdutil

import (
  "encoding/json"
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

func OutputJSON(v any) error {
  enc := json.NewEncoder(os.Stdout)
  enc.SetIndent("", "  ")
  return enc.Encode(v)
}

func APIError(err error) error {
  return fmt.Errorf("netbox API error: %w", err)
}

// ListCmd builds a cobra "list" subcommand that calls run with the command.
func ListCmd(noun string, run func(cmd *cobra.Command) error) *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all " + noun,
    RunE: func(cmd *cobra.Command, args []string) error {
      return run(cmd)
    },
  }
}

// GetCmd builds a cobra "get" subcommand that calls run with the command and parsed --id flag.
func GetCmd(noun string, run func(cmd *cobra.Command, id int32) error) *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a " + noun + " by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      return run(cmd, id)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, noun+" ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}
