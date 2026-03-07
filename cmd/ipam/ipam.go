// Package ipam implements CLI subcommands for the NetBox IPAM API.
package ipam

import (
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
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
		cmdutil.ListCmd("aggregates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamAggregatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("aggregate", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamAggregatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("asns", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamAsnsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("asn", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamAsnsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("asn-ranges", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamAsnRangesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("asn-range", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamAsnRangesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("fhrp-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("fhrp-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("fhrp-group-assignments", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("fhrp-group-assignment", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("ip-addresses", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamIpAddressesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ip-address", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamIpAddressesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("ip-ranges", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamIpRangesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("ip-range", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamIpRangesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("prefixes", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamPrefixesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("prefix", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamPrefixesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("rirs", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamRirsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rir", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamRirsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("roles", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamRolesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("role", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamRolesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("route-targets", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamRouteTargetsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("route-target", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamRouteTargetsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("service-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamServiceTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("service-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamServiceTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("services", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamServicesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("service", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamServicesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("vlan-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlanGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlanGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("vlan-translation-policies", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan-translation-policy", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("vlan-translation-rules", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan-translation-rule", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("vlans", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlansList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vlan", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVlansRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("vrfs", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVrfsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("vrf", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.IpamAPI.IpamVrfsRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
