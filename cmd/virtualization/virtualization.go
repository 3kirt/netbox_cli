// Package virtualization implements CLI subcommands for the NetBox virtualization API.
package virtualization

import (
	"context"
	"encoding/json"
	"fmt"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the virtualization API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "virtualization",
		Short: "Commands for the NetBox Virtualization API",
	}

	cmd.AddCommand(
		clusterGroupsCmd(),
		clusterTypesCmd(),
		clustersCmd(),
		interfacesCmd(),
		virtualDisksCmd(),
		virtualMachinesCmd(),
	)

	return cmd
}

// clusterGroupsCmd -------------------------------------------------------

func clusterGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "cluster-groups", Short: "Manage cluster groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("cluster-groups", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterGroupsList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("cluster-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("cluster-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ClusterGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationClusterGroupsCreate(ctx).ClusterGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("cluster-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedClusterGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationClusterGroupsPartialUpdate(ctx, id).PatchedClusterGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("cluster-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VirtualizationAPI.VirtualizationClusterGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// clusterTypesCmd -------------------------------------------------------

func clusterTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "cluster-types", Short: "Manage cluster types"}
	cmd.AddCommand(
		cmdutil.ListCmd("cluster-types", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterTypesList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("cluster-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("cluster-type", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ClusterTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationClusterTypesCreate(ctx).ClusterTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("cluster-type", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedClusterTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationClusterTypesPartialUpdate(ctx, id).PatchedClusterTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("cluster-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VirtualizationAPI.VirtualizationClusterTypesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// clustersCmd -------------------------------------------------------

func clustersCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "clusters", Short: "Manage clusters"}
	cmd.AddCommand(
		clustersListCmd(),
		cmdutil.GetCmd("cluster", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClustersRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("cluster", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableClusterRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationClustersCreate(ctx).WritableClusterRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("cluster", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableClusterRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationClustersPartialUpdate(ctx, id).PatchedWritableClusterRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("cluster", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VirtualizationAPI.VirtualizationClustersDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func clustersListCmd() *cobra.Command {
	var name, site, search string
	var limit int32
	var fields []string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all clusters",
		Long:  "List clusters. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli virtualization clusters list --name prod-vsphere-01
  netbox-cli virtualization clusters list --site azure-canada-central
  netbox-cli virtualization clusters list --search vsphere`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.VirtualizationAPI.VirtualizationClustersList(cmd.Context()).Limit(limit)
			if search != "" {
				req = req.Q(search)
			}
			if name != "" {
				req = req.Name([]string{name})
			}
			if site != "" {
				req = req.Site([]string{site})
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search")
	cmd.Flags().StringVar(&name, "name", "", "filter by exact name")
	cmd.Flags().StringVar(&site, "site", "", "filter by site slug")
	cmd.Flags().Int32Var(&limit, "limit", 0, "maximum number of records to return (default 0: return all)")
	cmd.Flags().StringSliceVar(&fields, "fields", nil, "comma-separated top-level fields to include in output (e.g. id,name,status)")
	return cmd
}

// interfacesCmd -------------------------------------------------------

func interfacesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "interfaces", Short: "Manage VM interfaces"}
	cmd.AddCommand(
		vmInterfacesListCmd(),
		cmdutil.GetCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationInterfacesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("interface", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVMInterfaceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationInterfacesCreate(ctx).WritableVMInterfaceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVMInterfaceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationInterfacesPartialUpdate(ctx, id).PatchedWritableVMInterfaceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VirtualizationAPI.VirtualizationInterfacesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func vmInterfacesListCmd() *cobra.Command {
	var virtualMachine, search string
	var limit int32
	var fields []string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all interfaces",
		Long:  "List VM interfaces. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli virtualization interfaces list --virtual-machine web-01
  netbox-cli virtualization interfaces list --search eth`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.VirtualizationAPI.VirtualizationInterfacesList(cmd.Context()).Limit(limit)
			if search != "" {
				req = req.Q(search)
			}
			if virtualMachine != "" {
				req = req.VirtualMachine([]string{virtualMachine})
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search")
	cmd.Flags().StringVar(&virtualMachine, "virtual-machine", "", "filter by virtual machine name")
	cmd.Flags().Int32Var(&limit, "limit", 0, "maximum number of records to return (default 0: return all)")
	cmd.Flags().StringSliceVar(&fields, "fields", nil, "comma-separated top-level fields to include in output (e.g. id,name,status)")
	return cmd
}

// virtualDisksCmd -------------------------------------------------------

func virtualDisksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-disks", Short: "Manage virtual disks"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-disks", func(ctx context.Context, client *netbox.APIClient, limit int32, fields []string) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualDisksList(ctx).Limit(limit).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		}),
		cmdutil.GetCmd("virtual-disk", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualDisksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("virtual-disk", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.VirtualDiskRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationVirtualDisksCreate(ctx).VirtualDiskRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-disk", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedVirtualDiskRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationVirtualDisksPartialUpdate(ctx, id).PatchedVirtualDiskRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-disk", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VirtualizationAPI.VirtualizationVirtualDisksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// virtualMachinesCmd -------------------------------------------------------

func virtualMachinesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-machines", Short: "Manage virtual machines"}
	cmd.AddCommand(
		virtualMachinesListCmd(),
		virtualMachinesGetCmd(),
		cmdutil.CreateCmd("virtual-machine", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVirtualMachineWithConfigContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationVirtualMachinesCreate(ctx).WritableVirtualMachineWithConfigContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-machine", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVirtualMachineWithConfigContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.VirtualizationAPI.VirtualizationVirtualMachinesPartialUpdate(ctx, id).PatchedWritableVirtualMachineWithConfigContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-machine", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.VirtualizationAPI.VirtualizationVirtualMachinesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func virtualMachinesListCmd() *cobra.Command {
	var name, nameContains, site, role, cluster, status, search string
	var tags []string
	var limit int32
	var fields []string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all virtual-machines",
		Long:  "List virtual machines. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli virtualization virtual-machines list --name web-01
  netbox-cli virtualization virtual-machines list --name-contains surveys
  netbox-cli virtualization virtual-machines list --site lon01 --status active
  netbox-cli virtualization virtual-machines list --cluster prod-vsphere-01
  netbox-cli virtualization virtual-machines list --search prod-web
  netbox-cli virtualization virtual-machines list --tag k8s-node`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.VirtualizationAPI.VirtualizationVirtualMachinesList(cmd.Context()).Limit(limit)
			if search != "" {
				req = req.Q(search)
			}
			if name != "" {
				req = req.Name([]string{name})
			}
			if nameContains != "" {
				req = req.NameIc([]string{nameContains})
			}
			if site != "" {
				req = req.Site([]string{site})
			}
			if role != "" {
				req = req.Role([]string{role})
			}
			if cluster != "" {
				req = req.Cluster([]string{cluster})
			}
			if status != "" {
				req = req.Status([]string{status})
			}
			if len(tags) > 0 {
				req = req.Tag(tags)
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSONFields(resp.GetResults(), fields)
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search")
	cmd.Flags().StringVar(&name, "name", "", "filter by exact name")
	cmd.Flags().StringVar(&nameContains, "name-contains", "", "filter by case-insensitive name substring")
	cmd.Flags().StringVar(&site, "site", "", "filter by site slug")
	cmd.Flags().StringVar(&role, "role", "", "filter by role slug")
	cmd.Flags().StringVar(&cluster, "cluster", "", "filter by cluster name")
	cmd.Flags().StringVar(&status, "status", "", "filter by status (active, staged, offline, planned, decommissioning)")
	cmd.Flags().StringSliceVar(&tags, "tag", nil, "filter by tag slug; multiple values require ALL tags to be present (comma-separated or repeated: --tag a,b or --tag a --tag b)")
	cmd.Flags().Int32Var(&limit, "limit", 0, "maximum number of records to return (default 0: return all)")
	cmd.Flags().StringSliceVar(&fields, "fields", nil, "comma-separated top-level fields to include in output (e.g. id,name,status)")
	return cmd
}

func virtualMachinesGetCmd() *cobra.Command {
	var id int32
	var name string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a virtual-machine by ID or name",
		Long:  "Get a virtual machine by numeric ID or exact name. Exactly one of --id or --name is required.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			if id != 0 && name != "" {
				return fmt.Errorf("only one of --id or --name may be specified")
			}
			if name != "" {
				resp, _, err := client.VirtualizationAPI.
					VirtualizationVirtualMachinesList(cmd.Context()).Name([]string{name}).Execute()
				if err != nil {
					return cmdutil.APIError(err)
				}
				results := resp.GetResults()
				switch len(results) {
				case 0:
					return fmt.Errorf("no virtual-machine found with name %q", name)
				case 1:
					return cmdutil.OutputJSON(results[0])
				default:
					return fmt.Errorf("%d virtual-machines matched name %q; use --id for an exact lookup", len(results), name)
				}
			}
			if id == 0 {
				return fmt.Errorf("either --id or --name is required")
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualMachinesRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		},
	}
	cmd.Flags().Int32Var(&id, "id", 0, "virtual-machine ID")
	cmd.Flags().StringVar(&name, "name", "", "virtual-machine name (exact match)")
	return cmd
}
