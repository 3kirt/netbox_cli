// Package vpn implements CLI subcommands for the NetBox VPN API.
package vpn

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the VPN API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Commands for the NetBox VPN API",
	}

	cmd.AddCommand(
		ikePoliciesCmd(),
		ikeProposalsCmd(),
		ipsecPoliciesCmd(),
		ipsecProfilesCmd(),
		ipsecProposalsCmd(),
		l2vpnTerminationsCmd(),
		l2vpnsCmd(),
		tunnelGroupsCmd(),
		tunnelTerminationsCmd(),
		tunnelsCmd(),
	)

	return cmd
}

// ikePoliciesCmd -------------------------------------------------------

func ikePoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ike-policies", Short: "Manage IKE policies"}
	cmd.AddCommand(
		cmdutil.ListCmd("ike-policies", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnIkePoliciesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ike-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIkePoliciesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// ikeProposalsCmd -------------------------------------------------------

func ikeProposalsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ike-proposals", Short: "Manage IKE proposals"}
	cmd.AddCommand(
		cmdutil.ListCmd("ike-proposals", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnIkeProposalsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ike-proposal", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIkeProposalsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// ipsecPoliciesCmd -------------------------------------------------------

func ipsecPoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ipsec-policies", Short: "Manage IPSec policies"}
	cmd.AddCommand(
		cmdutil.ListCmd("ipsec-policies", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnIpsecPoliciesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ipsec-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIpsecPoliciesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// ipsecProfilesCmd -------------------------------------------------------

func ipsecProfilesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ipsec-profiles", Short: "Manage IPSec profiles"}
	cmd.AddCommand(
		cmdutil.ListCmd("ipsec-profiles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnIpsecProfilesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ipsec-profile", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIpsecProfilesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// ipsecProposalsCmd -------------------------------------------------------

func ipsecProposalsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ipsec-proposals", Short: "Manage IPSec proposals"}
	cmd.AddCommand(
		cmdutil.ListCmd("ipsec-proposals", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnIpsecProposalsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ipsec-proposal", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIpsecProposalsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// l2vpnTerminationsCmd -------------------------------------------------------

func l2vpnTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "l2vpn-terminations", Short: "Manage L2VPN terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("l2vpn-terminations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnL2vpnTerminationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("l2vpn-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnL2vpnTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// l2vpnsCmd -------------------------------------------------------

func l2vpnsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "l2vpns", Short: "Manage L2VPNs"}
	cmd.AddCommand(
		cmdutil.ListCmd("l2vpns", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnL2vpnsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("l2vpn", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnL2vpnsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tunnelGroupsCmd -------------------------------------------------------

func tunnelGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tunnel-groups", Short: "Manage tunnel groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("tunnel-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnTunnelGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tunnel-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnTunnelGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tunnelTerminationsCmd -------------------------------------------------------

func tunnelTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tunnel-terminations", Short: "Manage tunnel terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("tunnel-terminations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnTunnelTerminationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tunnel-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnTunnelTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tunnelsCmd -------------------------------------------------------

func tunnelsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tunnels", Short: "Manage tunnels"}
	cmd.AddCommand(
		cmdutil.ListCmd("tunnels", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VpnAPI.VpnTunnelsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tunnel", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnTunnelsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
