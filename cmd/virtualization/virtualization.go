// Package virtualization implements CLI subcommands for the NetBox virtualization API.
package virtualization

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the virtualization API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "virtualization",
		Short: "Commands for the NetBox virtualization API",
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
		cmdutil.ListCmd("cluster-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cluster-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// clusterTypesCmd -------------------------------------------------------

func clusterTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "cluster-types", Short: "Manage cluster types"}
	cmd.AddCommand(
		cmdutil.ListCmd("cluster-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cluster-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// clustersCmd -------------------------------------------------------

func clustersCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "clusters", Short: "Manage clusters"}
	cmd.AddCommand(
		cmdutil.ListCmd("clusters", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClustersList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cluster", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClustersRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// interfacesCmd -------------------------------------------------------

func interfacesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "interfaces", Short: "Manage VM interfaces"}
	cmd.AddCommand(
		cmdutil.ListCmd("interfaces", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationInterfacesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationInterfacesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// virtualDisksCmd -------------------------------------------------------

func virtualDisksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-disks", Short: "Manage virtual disks"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-disks", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualDisksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-disk", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualDisksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// virtualMachinesCmd -------------------------------------------------------

func virtualMachinesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-machines", Short: "Manage virtual machines"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-machines", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualMachinesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-machine", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualMachinesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
