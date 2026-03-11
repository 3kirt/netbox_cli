// Package cmd wires together the root cobra command, configuration loading,
// and NetBox API client initialisation. Each NetBox API area is registered as
// a subcommand group from its own package under cmd/.
package cmd

import (
	"fmt"
	"os"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/cmd/circuits"
	"github.com/kirtis/netbox-cli/cmd/core"
	"github.com/kirtis/netbox-cli/cmd/dcim"
	"github.com/kirtis/netbox-cli/cmd/extras"
	"github.com/kirtis/netbox-cli/cmd/ipam"
	"github.com/kirtis/netbox-cli/cmd/tenancy"
	"github.com/kirtis/netbox-cli/cmd/users"
	"github.com/kirtis/netbox-cli/cmd/virtualization"
	"github.com/kirtis/netbox-cli/cmd/vpn"
	"github.com/kirtis/netbox-cli/cmd/wireless"
	"github.com/kirtis/netbox-cli/internal/clientctx"
	"github.com/kirtis/netbox-cli/internal/config"
)

// version is the build version. It is overridden at link time by GoReleaser
// and by `make build` via -ldflags "-X github.com/kirtis/netbox-cli/cmd.version=<tag>".
var version = "dev"

var configPath string

var rootCmd = &cobra.Command{
	Use:     "netbox-cli",
	Short:   "Command-line interface for the NetBox API",
	Version: version,
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		cfg, err := config.Load(configPath)
		if err != nil {
			return err
		}

		url, err := cfg.ResolveURL()
		if err != nil {
			return err
		}

		token, err := cfg.ResolveToken()
		if err != nil {
			return err
		}

		client := netbox.NewAPIClientFor(url, token)
		cmd.SetContext(clientctx.WithClient(cmd.Context(), client))
		return nil
	},
}

// Execute runs the root command and exits with a non-zero status on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// SilenceErrors prevents cobra from printing errors itself; Execute() owns
	// the output so errors are printed exactly once.
	// SilenceUsage suppresses the usage block on errors so the error message
	// is not buried in noise.
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "path to JSON config file (default: ~/.netbox_cli.json)")
	rootCmd.AddCommand(circuits.Command())
	rootCmd.AddCommand(core.Command())
	rootCmd.AddCommand(dcim.Command())
	rootCmd.AddCommand(extras.Command())
	rootCmd.AddCommand(ipam.Command())
	rootCmd.AddCommand(tenancy.Command())
	rootCmd.AddCommand(users.Command())
	rootCmd.AddCommand(virtualization.Command())
	rootCmd.AddCommand(vpn.Command())
	rootCmd.AddCommand(wireless.Command())
}
