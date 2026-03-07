package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"

  netbox "github.com/netbox-community/go-netbox/v4"
  "github.com/kirtis/netbox-cli/cmd/virtualization"
  "github.com/kirtis/netbox-cli/internal/clientctx"
  "github.com/kirtis/netbox-cli/internal/config"
)

var configPath string

var rootCmd = &cobra.Command{
  Use:   "netbox-cli",
  Short: "Command-line interface for the NetBox API",
  PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
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

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}

func init() {
  rootCmd.PersistentFlags().StringVar(&configPath, "config", "", "path to JSON config file (default: ~/.netbox_cli.json)")
  rootCmd.AddCommand(virtualization.Command())
}
