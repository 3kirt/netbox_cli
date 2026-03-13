// Package vpn implements CLI subcommands for the NetBox VPN API.
package vpn

import (
	"context"
	"encoding/json"
	"fmt"

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
		cmdutil.ListCmd("ike-policies", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnIkePoliciesList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("ike-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIkePoliciesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("ike-policy", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIKEPolicyRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIkePoliciesCreate(ctx).WritableIKEPolicyRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ike-policy", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIKEPolicyRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIkePoliciesPartialUpdate(ctx, id).PatchedWritableIKEPolicyRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ike-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnIkePoliciesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// ikeProposalsCmd -------------------------------------------------------

func ikeProposalsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ike-proposals", Short: "Manage IKE proposals"}
	cmd.AddCommand(
		cmdutil.ListCmd("ike-proposals", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnIkeProposalsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("ike-proposal", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIkeProposalsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("ike-proposal", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIKEProposalRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIkeProposalsCreate(ctx).WritableIKEProposalRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ike-proposal", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIKEProposalRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIkeProposalsPartialUpdate(ctx, id).PatchedWritableIKEProposalRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ike-proposal", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnIkeProposalsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// ipsecPoliciesCmd -------------------------------------------------------

func ipsecPoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ipsec-policies", Short: "Manage IPSec policies"}
	cmd.AddCommand(
		cmdutil.ListCmd("ipsec-policies", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnIpsecPoliciesList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("ipsec-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIpsecPoliciesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("ipsec-policy", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIPSecPolicyRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIpsecPoliciesCreate(ctx).WritableIPSecPolicyRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ipsec-policy", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIPSecPolicyRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIpsecPoliciesPartialUpdate(ctx, id).PatchedWritableIPSecPolicyRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ipsec-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnIpsecPoliciesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// ipsecProfilesCmd -------------------------------------------------------

func ipsecProfilesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ipsec-profiles", Short: "Manage IPSec profiles"}
	cmd.AddCommand(
		cmdutil.ListCmd("ipsec-profiles", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnIpsecProfilesList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("ipsec-profile", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIpsecProfilesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("ipsec-profile", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIPSecProfileRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIpsecProfilesCreate(ctx).WritableIPSecProfileRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ipsec-profile", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIPSecProfileRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIpsecProfilesPartialUpdate(ctx, id).PatchedWritableIPSecProfileRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ipsec-profile", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnIpsecProfilesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// ipsecProposalsCmd -------------------------------------------------------

func ipsecProposalsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ipsec-proposals", Short: "Manage IPSec proposals"}
	cmd.AddCommand(
		cmdutil.ListCmd("ipsec-proposals", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnIpsecProposalsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("ipsec-proposal", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnIpsecProposalsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("ipsec-proposal", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIPSecProposalRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIpsecProposalsCreate(ctx).WritableIPSecProposalRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ipsec-proposal", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIPSecProposalRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnIpsecProposalsPartialUpdate(ctx, id).PatchedWritableIPSecProposalRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ipsec-proposal", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnIpsecProposalsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// l2vpnTerminationsCmd -------------------------------------------------------

func l2vpnTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "l2vpn-terminations", Short: "Manage L2VPN terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("l2vpn-terminations", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnL2vpnTerminationsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("l2vpn-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnL2vpnTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("l2vpn-termination", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.L2VPNTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnL2vpnTerminationsCreate(ctx).L2VPNTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("l2vpn-termination", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedL2VPNTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnL2vpnTerminationsPartialUpdate(ctx, id).PatchedL2VPNTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("l2vpn-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnL2vpnTerminationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// l2vpnsCmd -------------------------------------------------------

func l2vpnsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "l2vpns", Short: "Manage L2VPNs"}
	cmd.AddCommand(
		cmdutil.ListCmd("l2vpns", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnL2vpnsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("l2vpn", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnL2vpnsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("l2vpn", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableL2VPNRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnL2vpnsCreate(ctx).WritableL2VPNRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("l2vpn", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableL2VPNRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnL2vpnsPartialUpdate(ctx, id).PatchedWritableL2VPNRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("l2vpn", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnL2vpnsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// tunnelGroupsCmd -------------------------------------------------------

func tunnelGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tunnel-groups", Short: "Manage tunnel groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("tunnel-groups", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnTunnelGroupsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("tunnel-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnTunnelGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("tunnel-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.TunnelGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnTunnelGroupsCreate(ctx).TunnelGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("tunnel-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedTunnelGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnTunnelGroupsPartialUpdate(ctx, id).PatchedTunnelGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("tunnel-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnTunnelGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// tunnelTerminationsCmd -------------------------------------------------------

func tunnelTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tunnel-terminations", Short: "Manage tunnel terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("tunnel-terminations", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnTunnelTerminationsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("tunnel-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnTunnelTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("tunnel-termination", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableTunnelTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnTunnelTerminationsCreate(ctx).WritableTunnelTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("tunnel-termination", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableTunnelTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnTunnelTerminationsPartialUpdate(ctx, id).PatchedWritableTunnelTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("tunnel-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnTunnelTerminationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// tunnelsCmd -------------------------------------------------------

func tunnelsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tunnels", Short: "Manage tunnels"}
	cmd.AddCommand(
		cmdutil.ListCmd("tunnels", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VpnAPI.VpnTunnelsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("tunnel", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VpnAPI.VpnTunnelsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("tunnel", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableTunnelRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnTunnelsCreate(ctx).WritableTunnelRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("tunnel", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableTunnelRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VpnAPI.VpnTunnelsPartialUpdate(ctx, id).PatchedWritableTunnelRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("tunnel", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VpnAPI.VpnTunnelsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
