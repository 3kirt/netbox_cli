// Package tenancy implements CLI subcommands for the NetBox Tenancy API.
package tenancy

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the Tenancy API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tenancy",
		Short: "Commands for the NetBox Tenancy API",
	}

	cmd.AddCommand(
		contactAssignmentsCmd(),
		contactGroupsCmd(),
		contactRolesCmd(),
		contactsCmd(),
		tenantGroupsCmd(),
		tenantsCmd(),
	)

	return cmd
}

// contactAssignmentsCmd -------------------------------------------------------

func contactAssignmentsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contact-assignments", Short: "Manage contact assignments"}
	cmd.AddCommand(
		cmdutil.ListCmd("contact-assignments", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact-assignment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// contactGroupsCmd -------------------------------------------------------

func contactGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contact-groups", Short: "Manage contact groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("contact-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.TenancyAPI.TenancyContactGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// contactRolesCmd -------------------------------------------------------

func contactRolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contact-roles", Short: "Manage contact roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("contact-roles", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.TenancyAPI.TenancyContactRolesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactRolesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// contactsCmd -------------------------------------------------------

func contactsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contacts", Short: "Manage contacts"}
	cmd.AddCommand(
		cmdutil.ListCmd("contacts", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.TenancyAPI.TenancyContactsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tenantGroupsCmd -------------------------------------------------------

func tenantGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tenant-groups", Short: "Manage tenant groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("tenant-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tenant-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tenantsCmd -------------------------------------------------------

func tenantsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tenants", Short: "Manage tenants"}
	cmd.AddCommand(
		cmdutil.ListCmd("tenants", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.TenancyAPI.TenancyTenantsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tenant", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.TenancyAPI.TenancyTenantsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
