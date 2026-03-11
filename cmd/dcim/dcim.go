// Package dcim implements CLI subcommands for the NetBox DCIM API.
package dcim

import (
	"context"
	"encoding/json"
	"fmt"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the DCIM API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dcim",
		Short: "Commands for the NetBox DCIM API",
	}

	cmd.AddCommand(
		cableTerminationsCmd(),
		cablesCmd(),
		consolePortTemplatesCmd(),
		consolePortsCmd(),
		consoleServerPortTemplatesCmd(),
		consoleServerPortsCmd(),
		deviceBayTemplatesCmd(),
		deviceBaysCmd(),
		deviceRolesCmd(),
		deviceTypesCmd(),
		devicesCmd(),
		frontPortTemplatesCmd(),
		frontPortsCmd(),
		interfaceTemplatesCmd(),
		interfacesCmd(),
		inventoryItemRolesCmd(),
		inventoryItemTemplatesCmd(),
		inventoryItemsCmd(),
		locationsCmd(),
		macAddressesCmd(),
		manufacturersCmd(),
		moduleBayTemplatesCmd(),
		moduleBaysCmd(),
		moduleTypeProfilesCmd(),
		moduleTypesCmd(),
		modulesCmd(),
		platformsCmd(),
		powerFeedsCmd(),
		powerOutletTemplatesCmd(),
		powerOutletsCmd(),
		powerPanelsCmd(),
		powerPortTemplatesCmd(),
		powerPortsCmd(),
		rackReservationsCmd(),
		rackRolesCmd(),
		rackTypesCmd(),
		racksCmd(),
		rearPortTemplatesCmd(),
		rearPortsCmd(),
		regionsCmd(),
		siteGroupsCmd(),
		sitesCmd(),
		virtualChassisCmd(),
		virtualDeviceContextsCmd(),
	)

	return cmd
}

// cableTerminationsCmd -------------------------------------------------------

func cableTerminationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "cable-terminations", Short: "Manage cable terminations"}
	cmd.AddCommand(
		cmdutil.ListCmd("cable-terminations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimCableTerminationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cable-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimCableTerminationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("cable-termination", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.CableTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimCableTerminationsCreate(ctx).CableTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("cable-termination", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedCableTerminationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimCableTerminationsPartialUpdate(ctx, id).PatchedCableTerminationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("cable-termination", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimCableTerminationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// cablesCmd -------------------------------------------------------

func cablesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "cables", Short: "Manage cables"}
	cmd.AddCommand(
		cmdutil.ListCmd("cables", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimCablesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cable", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimCablesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("cable", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableCableRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimCablesCreate(ctx).WritableCableRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("cable", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableCableRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimCablesPartialUpdate(ctx, id).PatchedWritableCableRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("cable", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimCablesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// consolePortTemplatesCmd -------------------------------------------------------

func consolePortTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "console-port-templates", Short: "Manage console port templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("console-port-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimConsolePortTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimConsolePortTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("console-port-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableConsolePortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsolePortTemplatesCreate(ctx).WritableConsolePortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("console-port-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableConsolePortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsolePortTemplatesPartialUpdate(ctx, id).PatchedWritableConsolePortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("console-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimConsolePortTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// consolePortsCmd -------------------------------------------------------

func consolePortsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "console-ports", Short: "Manage console ports"}
	cmd.AddCommand(
		cmdutil.ListCmd("console-ports", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimConsolePortsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimConsolePortsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("console-port", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableConsolePortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsolePortsCreate(ctx).WritableConsolePortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("console-port", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableConsolePortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsolePortsPartialUpdate(ctx, id).PatchedWritableConsolePortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("console-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimConsolePortsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// consoleServerPortTemplatesCmd -------------------------------------------------------

func consoleServerPortTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "console-server-port-templates", Short: "Manage console server port templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("console-server-port-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimConsoleServerPortTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-server-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimConsoleServerPortTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("console-server-port-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableConsoleServerPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortTemplatesCreate(ctx).WritableConsoleServerPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("console-server-port-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableConsoleServerPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortTemplatesPartialUpdate(ctx, id).PatchedWritableConsoleServerPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("console-server-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimConsoleServerPortTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// consoleServerPortsCmd -------------------------------------------------------

func consoleServerPortsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "console-server-ports", Short: "Manage console server ports"}
	cmd.AddCommand(
		cmdutil.ListCmd("console-server-ports", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimConsoleServerPortsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-server-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimConsoleServerPortsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("console-server-port", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableConsoleServerPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortsCreate(ctx).WritableConsoleServerPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("console-server-port", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableConsoleServerPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortsPartialUpdate(ctx, id).PatchedWritableConsoleServerPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("console-server-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimConsoleServerPortsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// deviceBayTemplatesCmd -------------------------------------------------------

func deviceBayTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "device-bay-templates", Short: "Manage device bay templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("device-bay-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimDeviceBayTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-bay-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimDeviceBayTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("device-bay-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.DeviceBayTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceBayTemplatesCreate(ctx).DeviceBayTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("device-bay-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedDeviceBayTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceBayTemplatesPartialUpdate(ctx, id).PatchedDeviceBayTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("device-bay-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimDeviceBayTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// deviceBaysCmd -------------------------------------------------------

func deviceBaysCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "device-bays", Short: "Manage device bays"}
	cmd.AddCommand(
		cmdutil.ListCmd("device-bays", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimDeviceBaysList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-bay", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimDeviceBaysRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("device-bay", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.DeviceBayRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceBaysCreate(ctx).DeviceBayRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("device-bay", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedDeviceBayRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceBaysPartialUpdate(ctx, id).PatchedDeviceBayRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("device-bay", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimDeviceBaysDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// deviceRolesCmd -------------------------------------------------------

func deviceRolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "device-roles", Short: "Manage device roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("device-roles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimDeviceRolesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimDeviceRolesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("device-role", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableDeviceRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceRolesCreate(ctx).WritableDeviceRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("device-role", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableDeviceRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceRolesPartialUpdate(ctx, id).PatchedWritableDeviceRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("device-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimDeviceRolesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// deviceTypesCmd -------------------------------------------------------

func deviceTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "device-types", Short: "Manage device types"}
	cmd.AddCommand(
		cmdutil.ListCmd("device-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimDeviceTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimDeviceTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("device-type", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableDeviceTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceTypesCreate(ctx).WritableDeviceTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("device-type", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableDeviceTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDeviceTypesPartialUpdate(ctx, id).PatchedWritableDeviceTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("device-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimDeviceTypesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// devicesCmd -------------------------------------------------------

func devicesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "devices", Short: "Manage devices"}
	cmd.AddCommand(
		devicesListCmd(),
		cmdutil.GetCmd("device", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimDevicesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("device", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableDeviceWithConfigContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDevicesCreate(ctx).WritableDeviceWithConfigContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("device", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableDeviceWithConfigContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimDevicesPartialUpdate(ctx, id).PatchedWritableDeviceWithConfigContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("device", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimDevicesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func devicesListCmd() *cobra.Command {
	var name, nameContains, site, status, role, search string
	var tags []string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all devices",
		Long:  "List devices. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli dcim devices list --site hq --status active
  netbox-cli dcim devices list --role access-switch
  netbox-cli dcim devices list --name-contains web
  netbox-cli dcim devices list --search core
  netbox-cli dcim devices list --tag ansible-managed`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.DcimAPI.DcimDevicesList(cmd.Context()).Limit(0)
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
			if status != "" {
				req = req.Status([]string{status})
			}
			if role != "" {
				req = req.Role([]string{role})
			}
			if len(tags) > 0 {
				req = req.Tag(tags)
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search")
	cmd.Flags().StringVar(&name, "name", "", "filter by exact name")
	cmd.Flags().StringVar(&nameContains, "name-contains", "", "filter by case-insensitive name substring")
	cmd.Flags().StringVar(&site, "site", "", "filter by site slug")
	cmd.Flags().StringVar(&status, "status", "", "filter by status (active, staged, offline, planned, decommissioning)")
	cmd.Flags().StringVar(&role, "role", "", "filter by role slug")
	cmd.Flags().StringSliceVar(&tags, "tag", nil, "filter by tag slug; multiple values require ALL tags to be present (comma-separated or repeated: --tag a,b or --tag a --tag b)")
	return cmd
}

// frontPortTemplatesCmd -------------------------------------------------------

func frontPortTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "front-port-templates", Short: "Manage front port templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("front-port-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimFrontPortTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("front-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimFrontPortTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("front-port-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableFrontPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimFrontPortTemplatesCreate(ctx).WritableFrontPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("front-port-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableFrontPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimFrontPortTemplatesPartialUpdate(ctx, id).PatchedWritableFrontPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("front-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimFrontPortTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// frontPortsCmd -------------------------------------------------------

func frontPortsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "front-ports", Short: "Manage front ports"}
	cmd.AddCommand(
		cmdutil.ListCmd("front-ports", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimFrontPortsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("front-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimFrontPortsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("front-port", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableFrontPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimFrontPortsCreate(ctx).WritableFrontPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("front-port", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableFrontPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimFrontPortsPartialUpdate(ctx, id).PatchedWritableFrontPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("front-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimFrontPortsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// interfaceTemplatesCmd -------------------------------------------------------

func interfaceTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "interface-templates", Short: "Manage interface templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("interface-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimInterfaceTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("interface-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimInterfaceTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("interface-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableInterfaceTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInterfaceTemplatesCreate(ctx).WritableInterfaceTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("interface-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableInterfaceTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInterfaceTemplatesPartialUpdate(ctx, id).PatchedWritableInterfaceTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("interface-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimInterfaceTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

func interfacesListCmd() *cobra.Command {
	var device, search string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all interfaces",
		Long:  "List device interfaces. All flags are optional; omitting them returns all records.",
		Example: `  netbox-cli dcim interfaces list --device core-sw-01
  netbox-cli dcim interfaces list --search eth0`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			req := client.DcimAPI.DcimInterfacesList(cmd.Context()).Limit(0)
			if search != "" {
				req = req.Q(search)
			}
			if device != "" {
				req = req.Device([]*string{&device})
			}
			resp, _, err := req.Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		},
	}
	cmd.Flags().StringVar(&search, "search", "", "free-text search")
	cmd.Flags().StringVar(&device, "device", "", "filter by device name")
	return cmd
}

// interfacesCmd -------------------------------------------------------

func interfacesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "interfaces", Short: "Manage device interfaces"}
	cmd.AddCommand(
		interfacesListCmd(),
		cmdutil.GetCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimInterfacesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("interface", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableInterfaceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInterfacesCreate(ctx).WritableInterfaceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableInterfaceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInterfacesPartialUpdate(ctx, id).PatchedWritableInterfaceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimInterfacesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// inventoryItemRolesCmd -------------------------------------------------------

func inventoryItemRolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "inventory-item-roles", Short: "Manage inventory item roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("inventory-item-roles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimInventoryItemRolesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("inventory-item-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimInventoryItemRolesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("inventory-item-role", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.InventoryItemRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemRolesCreate(ctx).InventoryItemRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("inventory-item-role", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedInventoryItemRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemRolesPartialUpdate(ctx, id).PatchedInventoryItemRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("inventory-item-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimInventoryItemRolesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// inventoryItemTemplatesCmd -------------------------------------------------------

func inventoryItemTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "inventory-item-templates", Short: "Manage inventory item templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("inventory-item-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimInventoryItemTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("inventory-item-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimInventoryItemTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("inventory-item-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.InventoryItemTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemTemplatesCreate(ctx).InventoryItemTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("inventory-item-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedInventoryItemTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemTemplatesPartialUpdate(ctx, id).PatchedInventoryItemTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("inventory-item-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimInventoryItemTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// inventoryItemsCmd -------------------------------------------------------

func inventoryItemsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "inventory-items", Short: "Manage inventory items"}
	cmd.AddCommand(
		cmdutil.ListCmd("inventory-items", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimInventoryItemsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("inventory-item", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimInventoryItemsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("inventory-item", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableInventoryItemRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemsCreate(ctx).WritableInventoryItemRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("inventory-item", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableInventoryItemRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemsPartialUpdate(ctx, id).PatchedWritableInventoryItemRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("inventory-item", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimInventoryItemsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// locationsCmd -------------------------------------------------------

func locationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "locations", Short: "Manage locations"}
	cmd.AddCommand(
		cmdutil.ListCmd("locations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimLocationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("location", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimLocationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("location", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableLocationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimLocationsCreate(ctx).WritableLocationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("location", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableLocationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimLocationsPartialUpdate(ctx, id).PatchedWritableLocationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("location", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimLocationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// macAddressesCmd -------------------------------------------------------

func macAddressesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "mac-addresses", Short: "Manage MAC addresses"}
	cmd.AddCommand(
		cmdutil.ListCmd("mac-addresses", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimMacAddressesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("mac-address", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimMacAddressesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("mac-address", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.MACAddressRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimMacAddressesCreate(ctx).MACAddressRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("mac-address", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedMACAddressRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimMacAddressesPartialUpdate(ctx, id).PatchedMACAddressRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("mac-address", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimMacAddressesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// manufacturersCmd -------------------------------------------------------

func manufacturersCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "manufacturers", Short: "Manage manufacturers"}
	cmd.AddCommand(
		cmdutil.ListCmd("manufacturers", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimManufacturersList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("manufacturer", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimManufacturersRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("manufacturer", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ManufacturerRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimManufacturersCreate(ctx).ManufacturerRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("manufacturer", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedManufacturerRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimManufacturersPartialUpdate(ctx, id).PatchedManufacturerRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("manufacturer", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimManufacturersDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// moduleBayTemplatesCmd -------------------------------------------------------

func moduleBayTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "module-bay-templates", Short: "Manage module bay templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("module-bay-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimModuleBayTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-bay-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimModuleBayTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("module-bay-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ModuleBayTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleBayTemplatesCreate(ctx).ModuleBayTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("module-bay-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedModuleBayTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleBayTemplatesPartialUpdate(ctx, id).PatchedModuleBayTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("module-bay-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimModuleBayTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// moduleBaysCmd -------------------------------------------------------

func moduleBaysCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "module-bays", Short: "Manage module bays"}
	cmd.AddCommand(
		cmdutil.ListCmd("module-bays", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimModuleBaysList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-bay", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimModuleBaysRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("module-bay", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ModuleBayRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleBaysCreate(ctx).ModuleBayRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("module-bay", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedModuleBayRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleBaysPartialUpdate(ctx, id).PatchedModuleBayRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("module-bay", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimModuleBaysDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// moduleTypeProfilesCmd -------------------------------------------------------

func moduleTypeProfilesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "module-type-profiles", Short: "Manage module type profiles"}
	cmd.AddCommand(
		cmdutil.ListCmd("module-type-profiles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimModuleTypeProfilesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-type-profile", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimModuleTypeProfilesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("module-type-profile", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ModuleTypeProfileRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleTypeProfilesCreate(ctx).ModuleTypeProfileRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("module-type-profile", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedModuleTypeProfileRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleTypeProfilesPartialUpdate(ctx, id).PatchedModuleTypeProfileRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("module-type-profile", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimModuleTypeProfilesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// moduleTypesCmd -------------------------------------------------------

func moduleTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "module-types", Short: "Manage module types"}
	cmd.AddCommand(
		cmdutil.ListCmd("module-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimModuleTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimModuleTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("module-type", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableModuleTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleTypesCreate(ctx).WritableModuleTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("module-type", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableModuleTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModuleTypesPartialUpdate(ctx, id).PatchedWritableModuleTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("module-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimModuleTypesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// modulesCmd -------------------------------------------------------

func modulesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "modules", Short: "Manage modules"}
	cmd.AddCommand(
		cmdutil.ListCmd("modules", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimModulesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimModulesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("module", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableModuleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModulesCreate(ctx).WritableModuleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("module", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableModuleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimModulesPartialUpdate(ctx, id).PatchedWritableModuleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("module", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimModulesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// platformsCmd -------------------------------------------------------

func platformsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "platforms", Short: "Manage platforms"}
	cmd.AddCommand(
		cmdutil.ListCmd("platforms", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPlatformsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("platform", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPlatformsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("platform", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.PlatformRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPlatformsCreate(ctx).PlatformRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("platform", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedPlatformRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPlatformsPartialUpdate(ctx, id).PatchedPlatformRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("platform", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPlatformsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// powerFeedsCmd -------------------------------------------------------

func powerFeedsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "power-feeds", Short: "Manage power feeds"}
	cmd.AddCommand(
		cmdutil.ListCmd("power-feeds", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPowerFeedsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-feed", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPowerFeedsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("power-feed", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritablePowerFeedRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerFeedsCreate(ctx).WritablePowerFeedRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("power-feed", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritablePowerFeedRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerFeedsPartialUpdate(ctx, id).PatchedWritablePowerFeedRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("power-feed", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPowerFeedsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// powerOutletTemplatesCmd -------------------------------------------------------

func powerOutletTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "power-outlet-templates", Short: "Manage power outlet templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("power-outlet-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPowerOutletTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-outlet-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPowerOutletTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("power-outlet-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritablePowerOutletTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletTemplatesCreate(ctx).WritablePowerOutletTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("power-outlet-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritablePowerOutletTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletTemplatesPartialUpdate(ctx, id).PatchedWritablePowerOutletTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("power-outlet-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPowerOutletTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// powerOutletsCmd -------------------------------------------------------

func powerOutletsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "power-outlets", Short: "Manage power outlets"}
	cmd.AddCommand(
		cmdutil.ListCmd("power-outlets", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPowerOutletsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-outlet", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPowerOutletsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("power-outlet", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritablePowerOutletRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletsCreate(ctx).WritablePowerOutletRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("power-outlet", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritablePowerOutletRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletsPartialUpdate(ctx, id).PatchedWritablePowerOutletRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("power-outlet", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPowerOutletsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// powerPanelsCmd -------------------------------------------------------

func powerPanelsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "power-panels", Short: "Manage power panels"}
	cmd.AddCommand(
		cmdutil.ListCmd("power-panels", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPowerPanelsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-panel", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPowerPanelsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("power-panel", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.PowerPanelRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerPanelsCreate(ctx).PowerPanelRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("power-panel", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedPowerPanelRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerPanelsPartialUpdate(ctx, id).PatchedPowerPanelRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("power-panel", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPowerPanelsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// powerPortTemplatesCmd -------------------------------------------------------

func powerPortTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "power-port-templates", Short: "Manage power port templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("power-port-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPowerPortTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPowerPortTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("power-port-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritablePowerPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerPortTemplatesCreate(ctx).WritablePowerPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("power-port-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritablePowerPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerPortTemplatesPartialUpdate(ctx, id).PatchedWritablePowerPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("power-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPowerPortTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// powerPortsCmd -------------------------------------------------------

func powerPortsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "power-ports", Short: "Manage power ports"}
	cmd.AddCommand(
		cmdutil.ListCmd("power-ports", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimPowerPortsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimPowerPortsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("power-port", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritablePowerPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerPortsCreate(ctx).WritablePowerPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("power-port", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritablePowerPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimPowerPortsPartialUpdate(ctx, id).PatchedWritablePowerPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("power-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimPowerPortsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// rackReservationsCmd -------------------------------------------------------

func rackReservationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rack-reservations", Short: "Manage rack reservations"}
	cmd.AddCommand(
		cmdutil.ListCmd("rack-reservations", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRackReservationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack-reservation", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRackReservationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("rack-reservation", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.RackReservationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRackReservationsCreate(ctx).RackReservationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rack-reservation", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedRackReservationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRackReservationsPartialUpdate(ctx, id).PatchedRackReservationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rack-reservation", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRackReservationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// rackRolesCmd -------------------------------------------------------

func rackRolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rack-roles", Short: "Manage rack roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("rack-roles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRackRolesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRackRolesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("rack-role", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.RackRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRackRolesCreate(ctx).RackRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rack-role", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedRackRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRackRolesPartialUpdate(ctx, id).PatchedRackRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rack-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRackRolesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// rackTypesCmd -------------------------------------------------------

func rackTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rack-types", Short: "Manage rack types"}
	cmd.AddCommand(
		cmdutil.ListCmd("rack-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRackTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRackTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("rack-type", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableRackTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRackTypesCreate(ctx).WritableRackTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rack-type", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableRackTypeRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRackTypesPartialUpdate(ctx, id).PatchedWritableRackTypeRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rack-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRackTypesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// racksCmd -------------------------------------------------------

func racksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "racks", Short: "Manage racks"}
	cmd.AddCommand(
		cmdutil.ListCmd("racks", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRacksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRacksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("rack", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableRackRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRacksCreate(ctx).WritableRackRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rack", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableRackRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRacksPartialUpdate(ctx, id).PatchedWritableRackRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rack", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRacksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// rearPortTemplatesCmd -------------------------------------------------------

func rearPortTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rear-port-templates", Short: "Manage rear port templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("rear-port-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRearPortTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rear-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRearPortTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("rear-port-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableRearPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRearPortTemplatesCreate(ctx).WritableRearPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rear-port-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableRearPortTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRearPortTemplatesPartialUpdate(ctx, id).PatchedWritableRearPortTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rear-port-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRearPortTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// rearPortsCmd -------------------------------------------------------

func rearPortsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "rear-ports", Short: "Manage rear ports"}
	cmd.AddCommand(
		cmdutil.ListCmd("rear-ports", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRearPortsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rear-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRearPortsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("rear-port", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableRearPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRearPortsCreate(ctx).WritableRearPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("rear-port", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableRearPortRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRearPortsPartialUpdate(ctx, id).PatchedWritableRearPortRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("rear-port", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRearPortsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// regionsCmd -------------------------------------------------------

func regionsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "regions", Short: "Manage regions"}
	cmd.AddCommand(
		cmdutil.ListCmd("regions", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimRegionsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("region", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimRegionsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("region", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableRegionRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRegionsCreate(ctx).WritableRegionRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("region", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableRegionRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimRegionsPartialUpdate(ctx, id).PatchedWritableRegionRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("region", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimRegionsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// siteGroupsCmd -------------------------------------------------------

func siteGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "site-groups", Short: "Manage site groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("site-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimSiteGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("site-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimSiteGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("site-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableSiteGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimSiteGroupsCreate(ctx).WritableSiteGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("site-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableSiteGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimSiteGroupsPartialUpdate(ctx, id).PatchedWritableSiteGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("site-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimSiteGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// sitesCmd -------------------------------------------------------

func sitesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "sites", Short: "Manage sites"}
	cmd.AddCommand(
		cmdutil.ListCmd("sites", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimSitesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("site", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimSitesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("site", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableSiteRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimSitesCreate(ctx).WritableSiteRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("site", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableSiteRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimSitesPartialUpdate(ctx, id).PatchedWritableSiteRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("site", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimSitesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// virtualChassisCmd -------------------------------------------------------

func virtualChassisCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-chassis", Short: "Manage virtual chassis"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-chassis", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimVirtualChassisList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-chassis", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimVirtualChassisRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("virtual-chassis", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVirtualChassisRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimVirtualChassisCreate(ctx).WritableVirtualChassisRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-chassis", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVirtualChassisRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimVirtualChassisPartialUpdate(ctx, id).PatchedWritableVirtualChassisRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-chassis", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimVirtualChassisDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// virtualDeviceContextsCmd -------------------------------------------------------

func virtualDeviceContextsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "virtual-device-contexts", Short: "Manage virtual device contexts"}
	cmd.AddCommand(
		cmdutil.ListCmd("virtual-device-contexts", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimVirtualDeviceContextsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-device-context", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimVirtualDeviceContextsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("virtual-device-context", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableVirtualDeviceContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimVirtualDeviceContextsCreate(ctx).WritableVirtualDeviceContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("virtual-device-context", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableVirtualDeviceContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.DcimAPI.DcimVirtualDeviceContextsPartialUpdate(ctx, id).PatchedWritableVirtualDeviceContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("virtual-device-context", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.DcimAPI.DcimVirtualDeviceContextsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
