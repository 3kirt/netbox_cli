// Package cmdutil provides shared helpers for building netbox-cli subcommands,
// including JSON output, error formatting, and cobra command constructors for
// the standard list and get actions.
package cmdutil

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// OutputJSON marshals v to indented JSON and writes it to stdout.
func OutputJSON(v any) error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// APIError wraps a NetBox API error with a descriptive prefix.
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
	// MarkFlagRequired only errors if the flag name is unknown, which cannot
	// happen here since we just registered "id" on this command.
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
