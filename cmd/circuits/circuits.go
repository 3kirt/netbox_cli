// Package circuits implements CLI subcommands for the NetBox Circuits API.
package circuits

import (
	"context"
	"encoding/json"
	"fmt"

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
		cmdutil.CreateCmd("circuit-group-assignment", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableCircuitGroupAssignmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupAssignmentsCreate(ctx).WritableCircuitGroupAssignmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("circuit-group-assignment", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableCircuitGroupAssignmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupAssignmentsPartialUpdate(ctx, id).PatchedWritableCircuitGroupAssignmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("circuit-group-assignment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsCircuitGroupAssignmentsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("circuit-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.CircuitGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupsCreate(ctx).CircuitGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("circuit-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedCircuitGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitGroupsPartialUpdate(ctx, id).PatchedCircuitGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("circuit-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsCircuitGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("circuit-termination", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.CircuitTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTerminationsCreate(ctx).CircuitTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("circuit-termination", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedCircuitTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTerminationsPartialUpdate(ctx, id).PatchedCircuitTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("circuit-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsCircuitTerminationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("circuit-type", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.CircuitTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTypesCreate(ctx).CircuitTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("circuit-type", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedCircuitTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitTypesPartialUpdate(ctx, id).PatchedCircuitTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("circuit-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsCircuitTypesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("circuit", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableCircuitRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitsCreate(ctx).WritableCircuitRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("circuit", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableCircuitRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsCircuitsPartialUpdate(ctx, id).PatchedWritableCircuitRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("circuit", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsCircuitsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("provider-account", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ProviderAccountRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsProviderAccountsCreate(ctx).ProviderAccountRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("provider-account", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedProviderAccountRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsProviderAccountsPartialUpdate(ctx, id).PatchedProviderAccountRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("provider-account", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsProviderAccountsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("provider-network", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ProviderNetworkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsProviderNetworksCreate(ctx).ProviderNetworkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("provider-network", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedProviderNetworkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsProviderNetworksPartialUpdate(ctx, id).PatchedProviderNetworkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("provider-network", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsProviderNetworksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("provider", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ProviderRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsProvidersCreate(ctx).ProviderRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("provider", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedProviderRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsProvidersPartialUpdate(ctx, id).PatchedProviderRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("provider", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsProvidersDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("virtual-circuit-termination", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVirtualCircuitTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTerminationsCreate(ctx).WritableVirtualCircuitTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-circuit-termination", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVirtualCircuitTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTerminationsPartialUpdate(ctx, id).PatchedWritableVirtualCircuitTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-circuit-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsVirtualCircuitTerminationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("virtual-circuit-type", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.VirtualCircuitTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTypesCreate(ctx).VirtualCircuitTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-circuit-type", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedVirtualCircuitTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitTypesPartialUpdate(ctx, id).PatchedVirtualCircuitTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-circuit-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsVirtualCircuitTypesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("virtual-circuit", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVirtualCircuitRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitsCreate(ctx).WritableVirtualCircuitRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-circuit", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVirtualCircuitRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CircuitsAPI.CircuitsVirtualCircuitsPartialUpdate(ctx, id).PatchedWritableVirtualCircuitRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-circuit", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CircuitsAPI.CircuitsVirtualCircuitsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
