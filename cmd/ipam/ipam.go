// Package ipam implements CLI subcommands for the NetBox IPAM API.
package ipam

import (
	"context"
	"encoding/json"
	"fmt"

	netbox "github.com/netbox-community/go-netbox/v4"
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
		cmdutil.ListCmd("aggregates", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamAggregatesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("aggregate", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableAggregateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamAggregatesCreate(ctx).WritableAggregateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("aggregate", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableAggregateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamAggregatesPartialUpdate(ctx, id).PatchedWritableAggregateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("aggregate", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamAggregatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// asnsCmd -------------------------------------------------------

func asnsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "asns", Short: "Manage ASNs"}
	cmd.AddCommand(
		cmdutil.ListCmd("asns", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamAsnsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("asn", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ASNRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamAsnsCreate(ctx).ASNRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("asn", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedASNRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamAsnsPartialUpdate(ctx, id).PatchedASNRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("asn", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamAsnsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// asnRangesCmd -------------------------------------------------------

func asnRangesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "asn-ranges", Short: "Manage ASN ranges"}
	cmd.AddCommand(
		cmdutil.ListCmd("asn-ranges", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamAsnRangesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("asn-range", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ASNRangeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamAsnRangesCreate(ctx).ASNRangeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("asn-range", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedASNRangeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamAsnRangesPartialUpdate(ctx, id).PatchedASNRangeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("asn-range", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamAsnRangesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// fhrpGroupsCmd -------------------------------------------------------

func fhrpGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "fhrp-groups", Short: "Manage FHRP groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("fhrp-groups", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamFhrpGroupsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("fhrp-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.FHRPGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupsCreate(ctx).FHRPGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("fhrp-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedFHRPGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupsPartialUpdate(ctx, id).PatchedFHRPGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("fhrp-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamFhrpGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// fhrpGroupAssignmentsCmd -------------------------------------------------------

func fhrpGroupAssignmentsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "fhrp-group-assignments", Short: "Manage FHRP group assignments"}
	cmd.AddCommand(
		cmdutil.ListCmd("fhrp-group-assignments", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("fhrp-group-assignment", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.FHRPGroupAssignmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsCreate(ctx).FHRPGroupAssignmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("fhrp-group-assignment", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedFHRPGroupAssignmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamFhrpGroupAssignmentsPartialUpdate(ctx, id).PatchedFHRPGroupAssignmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("fhrp-group-assignment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamFhrpGroupAssignmentsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// ipAddressesCmd -------------------------------------------------------

func ipAddressesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ip-addresses", Short: "Manage IP addresses"}
	cmd.AddCommand(
		ipAddressesListCmd(),
		cmdutil.GetCmd("ip-address", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamIpAddressesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("ip-address", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIPAddressRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamIpAddressesCreate(ctx).WritableIPAddressRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ip-address", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIPAddressRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamIpAddressesPartialUpdate(ctx, id).PatchedWritableIPAddressRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ip-address", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamIpAddressesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func ipAddressesListCmd() *cobra.Command {
	var virtualMachine, device, address, search string
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all ip-addresses",
		Long:  "List IP addresses. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli ipam ip-addresses list --virtual-machine sqdazrtst01.fes.corp
  netbox-cli ipam ip-addresses list --device core-sw-01
  netbox-cli ipam ip-addresses list --address 192.0.2.1/24
  netbox-cli ipam ip-addresses list --search 10.0.1`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.IpamAPI.IpamIpAddressesList(cmd.Context()).Limit(limit)
			if search != "" {
				req = req.Q(search)
			}
			if virtualMachine != "" {
				req = req.VirtualMachine([]string{virtualMachine})
			}
			if device != "" {
				req = req.Device([]string{device})
			}
			if address != "" {
				req = req.Address([]string{address})
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search")
	cmd.Flags().StringVar(&virtualMachine, "virtual-machine", "", "filter by virtual machine name")
	cmd.Flags().StringVar(&device, "device", "", "filter by device name")
	cmd.Flags().StringVar(&address, "address", "", "filter by IP address (e.g. 192.0.2.1/24)")
	cmd.Flags().Int32Var(&limit, "limit", 0, "maximum number of records to return (default 0: return all)")
	return cmd
}

// ipRangesCmd -------------------------------------------------------

func ipRangesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "ip-ranges", Short: "Manage IP ranges"}
	cmd.AddCommand(
		cmdutil.ListCmd("ip-ranges", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamIpRangesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("ip-range", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableIPRangeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamIpRangesCreate(ctx).WritableIPRangeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("ip-range", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableIPRangeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamIpRangesPartialUpdate(ctx, id).PatchedWritableIPRangeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("ip-range", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamIpRangesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// prefixesCmd -------------------------------------------------------

func prefixesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "prefixes", Short: "Manage prefixes"}
	cmd.AddCommand(
		prefixesListCmd(),
		cmdutil.GetCmd("prefix", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.IpamAPI.IpamPrefixesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("prefix", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritablePrefixRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamPrefixesCreate(ctx).WritablePrefixRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("prefix", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritablePrefixRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamPrefixesPartialUpdate(ctx, id).PatchedWritablePrefixRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("prefix", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamPrefixesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func prefixesListCmd() *cobra.Command {
	var search, site string
	var limit int32
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all prefixes",
		Long:  "List IP prefixes. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli ipam prefixes list --search 10.0
  netbox-cli ipam prefixes list --site hq`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.IpamAPI.IpamPrefixesList(cmd.Context()).Limit(limit)
			if search != "" {
				req = req.Q(search)
			}
			if site != "" {
				req = req.Site([]string{site})
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search (prefix string or description)")
	cmd.Flags().StringVar(&site, "site", "", "filter by site slug")
	cmd.Flags().Int32Var(&limit, "limit", 0, "maximum number of records to return (default 0: return all)")
	return cmd
}

// rirsCmd -------------------------------------------------------

func rirsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rirs", Short: "Manage RIRs"}
	cmd.AddCommand(
		cmdutil.ListCmd("rirs", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamRirsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("rir", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.RIRRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamRirsCreate(ctx).RIRRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rir", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedRIRRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamRirsPartialUpdate(ctx, id).PatchedRIRRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rir", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamRirsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// rolesCmd -------------------------------------------------------

func rolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "roles", Short: "Manage IP/VLAN roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("roles", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamRolesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("role", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.RoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamRolesCreate(ctx).RoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("role", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamRolesPartialUpdate(ctx, id).PatchedRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamRolesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// routeTargetsCmd -------------------------------------------------------

func routeTargetsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "route-targets", Short: "Manage route targets"}
	cmd.AddCommand(
		cmdutil.ListCmd("route-targets", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamRouteTargetsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("route-target", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.RouteTargetRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamRouteTargetsCreate(ctx).RouteTargetRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("route-target", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedRouteTargetRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamRouteTargetsPartialUpdate(ctx, id).PatchedRouteTargetRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("route-target", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamRouteTargetsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// serviceTemplatesCmd -------------------------------------------------------

func serviceTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "service-templates", Short: "Manage service templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("service-templates", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamServiceTemplatesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("service-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableServiceTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamServiceTemplatesCreate(ctx).WritableServiceTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("service-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableServiceTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamServiceTemplatesPartialUpdate(ctx, id).PatchedWritableServiceTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("service-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamServiceTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// servicesCmd -------------------------------------------------------

func servicesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "services", Short: "Manage services"}
	cmd.AddCommand(
		cmdutil.ListCmd("services", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamServicesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("service", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableServiceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamServicesCreate(ctx).WritableServiceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("service", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableServiceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamServicesPartialUpdate(ctx, id).PatchedWritableServiceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("service", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamServicesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// vlanGroupsCmd -------------------------------------------------------

func vlanGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlan-groups", Short: "Manage VLAN groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlan-groups", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamVlanGroupsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("vlan-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.VLANGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlanGroupsCreate(ctx).VLANGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("vlan-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedVLANGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlanGroupsPartialUpdate(ctx, id).PatchedVLANGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("vlan-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamVlanGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// vlanTranslationPoliciesCmd -------------------------------------------------------

func vlanTranslationPoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlan-translation-policies", Short: "Manage VLAN translation policies"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlan-translation-policies", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("vlan-translation-policy", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.VLANTranslationPolicyRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesCreate(ctx).VLANTranslationPolicyRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("vlan-translation-policy", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedVLANTranslationPolicyRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationPoliciesPartialUpdate(ctx, id).PatchedVLANTranslationPolicyRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("vlan-translation-policy", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamVlanTranslationPoliciesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// vlanTranslationRulesCmd -------------------------------------------------------

func vlanTranslationRulesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlan-translation-rules", Short: "Manage VLAN translation rules"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlan-translation-rules", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("vlan-translation-rule", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.VLANTranslationRuleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesCreate(ctx).VLANTranslationRuleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("vlan-translation-rule", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedVLANTranslationRuleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlanTranslationRulesPartialUpdate(ctx, id).PatchedVLANTranslationRuleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("vlan-translation-rule", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamVlanTranslationRulesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// vlansCmd -------------------------------------------------------

func vlansCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vlans", Short: "Manage VLANs"}
	cmd.AddCommand(
		cmdutil.ListCmd("vlans", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamVlansList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("vlan", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVLANRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlansCreate(ctx).WritableVLANRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("vlan", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVLANRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVlansPartialUpdate(ctx, id).PatchedWritableVLANRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("vlan", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamVlansDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// vrfsCmd -------------------------------------------------------

func vrfsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "vrfs", Short: "Manage VRFs"}
	cmd.AddCommand(
		cmdutil.ListCmd("vrfs", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.IpamAPI.IpamVrfsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("vrf", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.VRFRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVrfsCreate(ctx).VRFRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("vrf", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedVRFRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.IpamAPI.IpamVrfsPartialUpdate(ctx, id).PatchedVRFRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("vrf", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.IpamAPI.IpamVrfsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
