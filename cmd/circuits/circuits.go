// Package circuits implements CLI subcommands for the NetBox Circuits API.
package circuits

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the Circuits API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "circuits",
		Short: "Commands for the NetBox Circuits API",
	}

	cmd.AddCommand(
		circuitGroupAssignmentsCmd(),
		circuitGroupsCmd(),
		circuitTerminationsCmd(),
		circuitTypesCmd(),
		circuitsCmd(),
		providerAccountsCmd(),
		providerNetworksCmd(),
		providersCmd(),
		virtualCircuitTerminationsCmd(),
		virtualCircuitTypesCmd(),
		virtualCircuitsCmd(),
	)

	return cmd
}

// circuitGroupAssignmentsCmd -------------------------------------------------------

func circuitGroupAssignmentsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "circuit-group-assignments", Short: "Manage circuit group assignments"}
	cmd.AddCommand(
		cmdutil.ListCmd("circuit-group-assignments", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupAssignmentsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("circuit-group-assignment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupAssignmentsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// circuitGroupsCmd -------------------------------------------------------

func circuitGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "circuit-groups", Short: "Manage circuit groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("circuit-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("circuit-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// circuitTerminationsCmd -------------------------------------------------------

func circuitTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "circuit-terminations", Short: "Manage circuit terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("circuit-terminations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTerminationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("circuit-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// circuitTypesCmd -------------------------------------------------------

func circuitTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "circuit-types", Short: "Manage circuit types"}
	cmd.AddCommand(
		cmdutil.ListCmd("circuit-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("circuit-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// circuitsCmd -------------------------------------------------------

func circuitsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "circuits", Short: "Manage circuits"}
	cmd.AddCommand(
		cmdutil.ListCmd("circuits", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("circuit", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsCircuitsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// providerAccountsCmd -------------------------------------------------------

func providerAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "provider-accounts", Short: "Manage provider accounts"}
	cmd.AddCommand(
		cmdutil.ListCmd("provider-accounts", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsProviderAccountsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("provider-account", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsProviderAccountsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// providerNetworksCmd -------------------------------------------------------

func providerNetworksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "provider-networks", Short: "Manage provider networks"}
	cmd.AddCommand(
		cmdutil.ListCmd("provider-networks", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsProviderNetworksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("provider-network", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsProviderNetworksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// providersCmd -------------------------------------------------------

func providersCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "providers", Short: "Manage providers"}
	cmd.AddCommand(
		cmdutil.ListCmd("providers", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsProvidersList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("provider", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsProvidersRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// virtualCircuitTerminationsCmd -------------------------------------------------------

func virtualCircuitTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-circuit-terminations", Short: "Manage virtual circuit terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-circuit-terminations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTerminationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-circuit-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// virtualCircuitTypesCmd -------------------------------------------------------

func virtualCircuitTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-circuit-types", Short: "Manage virtual circuit types"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-circuit-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-circuit-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// virtualCircuitsCmd -------------------------------------------------------

func virtualCircuitsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-circuits", Short: "Manage virtual circuits"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-circuits", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-circuit", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
