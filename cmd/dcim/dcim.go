// Package dcim implements CLI subcommands for the NetBox DCIM API.
package dcim

import (
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
		cmdutil.ListCmd("cable-terminations", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimCableTerminationsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cable-termination", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimCableTerminationsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("cables", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimCablesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("cable", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimCablesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("console-port-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsolePortTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-port-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsolePortTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("console-ports", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsolePortsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-port", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsolePortsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("console-server-port-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-server-port-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("console-server-ports", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("console-server-port", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimConsoleServerPortsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("device-bay-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceBayTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-bay-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceBayTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("device-bays", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceBaysList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-bay", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceBaysRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("device-roles", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceRolesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-role", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceRolesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("device-types", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceTypesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device-type", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDeviceTypesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("devices", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDevicesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("device", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimDevicesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("front-port-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimFrontPortTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("front-port-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimFrontPortTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("front-ports", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimFrontPortsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("front-port", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimFrontPortsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("interface-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInterfaceTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("interface-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInterfaceTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("interfaces", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInterfacesList(cmd.Context()).Limit(0).Execute()
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
			resp, _, err := client.DcimAPI.DcimInterfacesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("inventory-item-roles", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemRolesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("inventory-item-role", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemRolesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("inventory-item-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("inventory-item-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("inventory-items", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("inventory-item", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimInventoryItemsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("locations", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimLocationsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("location", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimLocationsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("mac-addresses", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimMacAddressesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("mac-address", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimMacAddressesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("manufacturers", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimManufacturersList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("manufacturer", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimManufacturersRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("module-bay-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleBayTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-bay-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleBayTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("module-bays", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleBaysList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-bay", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleBaysRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("module-type-profiles", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleTypeProfilesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-type-profile", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleTypeProfilesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("module-types", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleTypesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module-type", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModuleTypesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("modules", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModulesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("module", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimModulesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("platforms", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPlatformsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("platform", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPlatformsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("power-feeds", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerFeedsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-feed", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerFeedsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("power-outlet-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-outlet-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("power-outlets", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-outlet", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerOutletsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("power-panels", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerPanelsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-panel", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerPanelsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("power-port-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerPortTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-port-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerPortTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("power-ports", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerPortsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("power-port", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimPowerPortsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("rack-reservations", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRackReservationsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack-reservation", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRackReservationsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("rack-roles", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRackRolesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack-role", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRackRolesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("rack-types", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRackTypesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack-type", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRackTypesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("racks", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRacksList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rack", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRacksRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("rear-port-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRearPortTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rear-port-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRearPortTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("rear-ports", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRearPortsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("rear-port", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRearPortsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("regions", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRegionsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("region", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimRegionsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("site-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimSiteGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("site-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimSiteGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("sites", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimSitesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("site", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimSitesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("virtual-chassis", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimVirtualChassisList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-chassis", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimVirtualChassisRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("virtual-device-contexts", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimVirtualDeviceContextsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("virtual-device-context", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.DcimAPI.DcimVirtualDeviceContextsRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
