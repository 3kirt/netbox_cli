// Package virtualization implements CLI subcommands for the NetBox virtualization API.
package virtualization

import (
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
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
		cmdutil.ListCmd("cluster-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cluster-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("cluster-types", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterTypesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cluster-type", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClusterTypesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("clusters", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClustersList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cluster", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationClustersRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("interfaces", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationInterfacesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("interface", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationInterfacesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("virtual-disks", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualDisksList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-disk", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualDisksRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("virtual-machines", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualMachinesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-machine", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.VirtualizationAPI.
				VirtualizationVirtualMachinesRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
