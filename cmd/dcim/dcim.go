// Package dcim implements CLI subcommands for the NetBox DCIM API.
package dcim

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

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
	)
	return cmd
}

// devicesCmd -------------------------------------------------------

func devicesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "devices", Short: "Manage devices"}
	cmd.AddCommand(
		cmdutil.ListCmd("devices", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimDevicesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimDevicesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
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
	)
	return cmd
}

// interfacesCmd -------------------------------------------------------

func interfacesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "interfaces", Short: "Manage device interfaces"}
	cmd.AddCommand(
		cmdutil.ListCmd("interfaces", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.DcimAPI.DcimInterfacesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("interface", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.DcimAPI.DcimInterfacesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
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
	)
	return cmd
}
