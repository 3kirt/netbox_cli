// Package tenancy implements CLI subcommands for the NetBox Tenancy API.
package tenancy

import (
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
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
		cmdutil.ListCmd("contact-assignments", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact-assignment", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("contact-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("contact-roles", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactRolesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact-role", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactRolesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("contacts", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("contact", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyContactsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("tenant-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tenant-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("tenants", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyTenantsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tenant", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.TenancyAPI.TenancyTenantsRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
