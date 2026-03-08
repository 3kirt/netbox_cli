// Package wireless implements CLI subcommands for the NetBox Wireless API.
package wireless

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the Wireless API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wireless",
		Short: "Commands for the NetBox Wireless API",
	}

	cmd.AddCommand(
		wirelessLanGroupsCmd(),
		wirelessLansCmd(),
		wirelessLinksCmd(),
	)

	return cmd
}

// wirelessLanGroupsCmd -------------------------------------------------------

func wirelessLanGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "wireless-lan-groups", Short: "Manage wireless LAN groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("wireless-lan-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.WirelessAPI.WirelessWirelessLanGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("wireless-lan-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.WirelessAPI.WirelessWirelessLanGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// wirelessLansCmd -------------------------------------------------------

func wirelessLansCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "wireless-lans", Short: "Manage wireless LANs"}
	cmd.AddCommand(
		cmdutil.ListCmd("wireless-lans", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.WirelessAPI.WirelessWirelessLansList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("wireless-lan", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.WirelessAPI.WirelessWirelessLansRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// wirelessLinksCmd -------------------------------------------------------

func wirelessLinksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "wireless-links", Short: "Manage wireless links"}
	cmd.AddCommand(
		cmdutil.ListCmd("wireless-links", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.WirelessAPI.WirelessWirelessLinksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("wireless-link", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.WirelessAPI.WirelessWirelessLinksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
