// Package ipam implements CLI subcommands for the NetBox IPAM API.
package ipam

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the IPAM API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ipam",
		Short: "Commands for the NetBox IPAM API",
	}

	cmd.AddCommand(
		aggregatesCmd(),
		asnsCmd(),
		asnRangesCmd(),
		fhrpGroupsCmd(),
		fhrpGroupAssignmentsCmd(),
		ipAddressesCmd(),
		ipRangesCmd(),
		prefixesCmd(),
		rirsCmd(),
		rolesCmd(),
		routeTargetsCmd(),
		serviceTemplatesCmd(),
		servicesCmd(),
		vlanGroupsCmd(),
		vlanTranslationPoliciesCmd(),
		vlanTranslationRulesCmd(),
		vlansCmd(),
		vrfsCmd(),
	)

	return cmd
}

// aggregatesCmd -------------------------------------------------------

func aggregatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "aggregates", Short: "Manage aggregates"}
	cmd.AddCommand(
		cmdutil.ListCmd("aggregates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamAggregatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("aggregate", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamAggregatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// asnsCmd -------------------------------------------------------

func asnsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "asns", Short: "Manage ASNs"}
	cmd.AddCommand(
		cmdutil.ListCmd("asns", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamAsnsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("asn", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamAsnsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// asnRangesCmd -------------------------------------------------------

func asnRangesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "asn-ranges", Short: "Manage ASN ranges"}
	cmd.AddCommand(
		cmdutil.ListCmd("asn-ranges", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamAsnRangesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("asn-range", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamAsnRangesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// fhrpGroupsCmd -------------------------------------------------------

func fhrpGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "fhrp-groups", Short: "Manage FHRP groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("fhrp-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamFhrpGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("fhrp-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamFhrpGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// fhrpGroupAssignmentsCmd -------------------------------------------------------

func fhrpGroupAssignmentsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "fhrp-group-assignments", Short: "Manage FHRP group assignments"}
	cmd.AddCommand(
		cmdutil.ListCmd("fhrp-group-assignments", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("fhrp-group-assignment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// ipAddressesCmd -------------------------------------------------------

func ipAddressesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ip-addresses", Short: "Manage IP addresses"}
	cmd.AddCommand(
		cmdutil.ListCmd("ip-addresses", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamIpAddressesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ip-address", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamIpAddressesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// ipRangesCmd -------------------------------------------------------

func ipRangesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ip-ranges", Short: "Manage IP ranges"}
	cmd.AddCommand(
		cmdutil.ListCmd("ip-ranges", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamIpRangesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ip-range", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamIpRangesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// prefixesCmd -------------------------------------------------------

func prefixesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "prefixes", Short: "Manage prefixes"}
	cmd.AddCommand(
		cmdutil.ListCmd("prefixes", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamPrefixesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("prefix", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamPrefixesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// rirsCmd -------------------------------------------------------

func rirsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rirs", Short: "Manage RIRs"}
	cmd.AddCommand(
		cmdutil.ListCmd("rirs", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamRirsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rir", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamRirsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// rolesCmd -------------------------------------------------------

func rolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "roles", Short: "Manage IP/VLAN roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("roles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamRolesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamRolesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// routeTargetsCmd -------------------------------------------------------

func routeTargetsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "route-targets", Short: "Manage route targets"}
	cmd.AddCommand(
		cmdutil.ListCmd("route-targets", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamRouteTargetsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("route-target", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamRouteTargetsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// serviceTemplatesCmd -------------------------------------------------------

func serviceTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "service-templates", Short: "Manage service templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("service-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamServiceTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("service-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamServiceTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// servicesCmd -------------------------------------------------------

func servicesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "services", Short: "Manage services"}
	cmd.AddCommand(
		cmdutil.ListCmd("services", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamServicesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("service", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamServicesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// vlanGroupsCmd -------------------------------------------------------

func vlanGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlan-groups", Short: "Manage VLAN groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlan-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamVlanGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamVlanGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// vlanTranslationPoliciesCmd -------------------------------------------------------

func vlanTranslationPoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlan-translation-policies", Short: "Manage VLAN translation policies"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlan-translation-policies", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan-translation-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// vlanTranslationRulesCmd -------------------------------------------------------

func vlanTranslationRulesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlan-translation-rules", Short: "Manage VLAN translation rules"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlan-translation-rules", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan-translation-rule", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// vlansCmd -------------------------------------------------------

func vlansCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlans", Short: "Manage VLANs"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlans", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamVlansList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamVlansRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// vrfsCmd -------------------------------------------------------

func vrfsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vrfs", Short: "Manage VRFs"}
	cmd.AddCommand(
		cmdutil.ListCmd("vrfs", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.IpamAPI.IpamVrfsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vrf", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamVrfsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
