// Package cmdutil provides shared helpers for building netbox-cli subcommands,
// including JSON output, error formatting, and cobra command constructors for
// the standard list and get actions.
package cmdutil

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
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

// ListCmd builds a cobra "list" subcommand. The run callback receives the
// context and the NetBox API client resolved from the command context.
func ListCmd(noun string, run func(context.Context, *netbox.APIClient) error) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all " + noun,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			return run(cmd.Context(), client)
		},
	}
}

// GetCmd builds a cobra "get" subcommand. The run callback receives the
// context, the NetBox API client resolved from the command context, and
// the value of the required --id flag.
func GetCmd(noun string, run func(context.Context, *netbox.APIClient, int32) error) *cobra.Command {
	var id int32
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a " + noun + " by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			return run(cmd.Context(), client, id)
		},
	}
	cmd.Flags().Int32Var(&id, "id", 0, noun+" ID (required)")
	// MarkFlagRequired only errors if the flag name is unknown, which cannot
	// happen here since we just registered "id" on this command.
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
