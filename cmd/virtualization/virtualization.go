package virtualization

import (
  "context"
  "encoding/json"
  "fmt"
  "os"

  "github.com/spf13/cobra"

  netbox "github.com/netbox-community/go-netbox/v4"
  "github.com/kirtis/netbox-cli/internal/clientctx"
)

func Command() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "virtualization",
    Short: "Commands for the NetBox virtualization API",
  }

  cmd.AddCommand(
    clusterTypesCmd(),
    clusterGroupsCmd(),
    clustersCmd(),
    virtualMachinesCmd(),
    interfacesCmd(),
  )

  return cmd
}

// outputJSON marshals v to indented JSON and writes it to stdout.
func outputJSON(v any) error {
  enc := json.NewEncoder(os.Stdout)
  enc.SetIndent("", "  ")
  return enc.Encode(v)
}

func apiError(err error) error {
  return fmt.Errorf("netbox API error: %w", err)
}

// clusterTypesCmd -------------------------------------------------------

func clusterTypesCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "cluster-types",
    Short: "Manage cluster types",
  }
  cmd.AddCommand(clusterTypesListCmd(), clusterTypesGetCmd())
  return cmd
}

func clusterTypesListCmd() *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all cluster types",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationClusterTypesList(cmd.Context()).
        Limit(0).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp.GetResults())
    },
  }
}

func clusterTypesGetCmd() *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a cluster type by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationClusterTypesRetrieve(cmd.Context(), id).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, "cluster type ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}

// clusterGroupsCmd -------------------------------------------------------

func clusterGroupsCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "cluster-groups",
    Short: "Manage cluster groups",
  }
  cmd.AddCommand(clusterGroupsListCmd(), clusterGroupsGetCmd())
  return cmd
}

func clusterGroupsListCmd() *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all cluster groups",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationClusterGroupsList(cmd.Context()).
        Limit(0).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp.GetResults())
    },
  }
}

func clusterGroupsGetCmd() *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a cluster group by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationClusterGroupsRetrieve(cmd.Context(), id).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, "cluster group ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}

// clustersCmd -------------------------------------------------------

func clustersCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "clusters",
    Short: "Manage clusters",
  }
  cmd.AddCommand(clustersListCmd(), clustersGetCmd())
  return cmd
}

func clustersListCmd() *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all clusters",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationClustersList(cmd.Context()).
        Limit(0).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp.GetResults())
    },
  }
}

func clustersGetCmd() *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a cluster by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationClustersRetrieve(cmd.Context(), id).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, "cluster ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}

// virtualMachinesCmd -------------------------------------------------------

func virtualMachinesCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "virtual-machines",
    Short: "Manage virtual machines",
  }
  cmd.AddCommand(virtualMachinesListCmd(), virtualMachinesGetCmd())
  return cmd
}

func virtualMachinesListCmd() *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all virtual machines",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationVirtualMachinesList(cmd.Context()).
        Limit(0).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp.GetResults())
    },
  }
}

func virtualMachinesGetCmd() *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a virtual machine by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationVirtualMachinesRetrieve(cmd.Context(), id).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, "virtual machine ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}

// interfacesCmd -------------------------------------------------------

func interfacesCmd() *cobra.Command {
  cmd := &cobra.Command{
    Use:   "interfaces",
    Short: "Manage virtual machine interfaces",
  }
  cmd.AddCommand(interfacesListCmd(), interfacesGetCmd())
  return cmd
}

func interfacesListCmd() *cobra.Command {
  return &cobra.Command{
    Use:   "list",
    Short: "List all VM interfaces",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationInterfacesList(cmd.Context()).
        Limit(0).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp.GetResults())
    },
  }
}

func interfacesGetCmd() *cobra.Command {
  var id int32
  cmd := &cobra.Command{
    Use:   "get",
    Short: "Get a VM interface by ID",
    RunE: func(cmd *cobra.Command, args []string) error {
      client, err := clientctx.Client(cmd.Context())
      if err != nil {
        return err
      }
      resp, _, err := client.VirtualizationAPI.
        VirtualizationInterfacesRetrieve(cmd.Context(), id).
        Execute()
      if err != nil {
        return apiError(err)
      }
      return outputJSON(resp)
    },
  }
  cmd.Flags().Int32Var(&id, "id", 0, "interface ID (required)")
  _ = cmd.MarkFlagRequired("id")
  return cmd
}

// WithClient and context helpers are in internal/clientctx — re-exported
// here for convenience so callers only need to import this package.

func WithClient(ctx context.Context, client *netbox.APIClient) context.Context {
  return clientctx.WithClient(ctx, client)
}
