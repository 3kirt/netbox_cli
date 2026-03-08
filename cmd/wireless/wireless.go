// Package wireless implements CLI subcommands for the NetBox Wireless API.
package wireless

import (
	"context"
	"encoding/json"
	"fmt"

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
		cmdutil.CreateCmd("wireless-lan-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableWirelessLANGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.WirelessAPI.WirelessWirelessLanGroupsCreate(ctx).WritableWirelessLANGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("wireless-lan-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableWirelessLANGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.WirelessAPI.WirelessWirelessLanGroupsPartialUpdate(ctx, id).PatchedWritableWirelessLANGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("wireless-lan-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.WirelessAPI.WirelessWirelessLanGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("wireless-lan", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableWirelessLANRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.WirelessAPI.WirelessWirelessLansCreate(ctx).WritableWirelessLANRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("wireless-lan", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableWirelessLANRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.WirelessAPI.WirelessWirelessLansPartialUpdate(ctx, id).PatchedWritableWirelessLANRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("wireless-lan", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.WirelessAPI.WirelessWirelessLansDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("wireless-link", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableWirelessLinkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.WirelessAPI.WirelessWirelessLinksCreate(ctx).WritableWirelessLinkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("wireless-link", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableWirelessLinkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.WirelessAPI.WirelessWirelessLinksPartialUpdate(ctx, id).PatchedWritableWirelessLinkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("wireless-link", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.WirelessAPI.WirelessWirelessLinksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
