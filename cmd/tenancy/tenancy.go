// Package tenancy implements CLI subcommands for the NetBox Tenancy API.
package tenancy

import (
	"context"
	"encoding/json"
	"fmt"

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
		cmdutil.ListCmd("contact-assignments", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("contact-assignment", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableContactAssignmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsCreate(ctx).WritableContactAssignmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("contact-assignment", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableContactAssignmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactAssignmentsPartialUpdate(ctx, id).PatchedWritableContactAssignmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("contact-assignment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.TenancyAPI.TenancyContactAssignmentsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// contactGroupsCmd -------------------------------------------------------

func contactGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contact-groups", Short: "Manage contact groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("contact-groups", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactGroupsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("contact-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableContactGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactGroupsCreate(ctx).WritableContactGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("contact-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableContactGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactGroupsPartialUpdate(ctx, id).PatchedWritableContactGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("contact-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.TenancyAPI.TenancyContactGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// contactRolesCmd -------------------------------------------------------

func contactRolesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contact-roles", Short: "Manage contact roles"}
	cmd.AddCommand(
		cmdutil.ListCmd("contact-roles", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactRolesList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("contact-role", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ContactRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactRolesCreate(ctx).ContactRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("contact-role", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedContactRoleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactRolesPartialUpdate(ctx, id).PatchedContactRoleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("contact-role", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.TenancyAPI.TenancyContactRolesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// contactsCmd -------------------------------------------------------

func contactsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "contacts", Short: "Manage contacts"}
	cmd.AddCommand(
		cmdutil.ListCmd("contacts", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.TenancyAPI.TenancyContactsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("contact", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ContactRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactsCreate(ctx).ContactRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("contact", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedContactRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyContactsPartialUpdate(ctx, id).PatchedContactRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("contact", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.TenancyAPI.TenancyContactsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// tenantGroupsCmd -------------------------------------------------------

func tenantGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tenant-groups", Short: "Manage tenant groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("tenant-groups", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("tenant-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableTenantGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsCreate(ctx).WritableTenantGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("tenant-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableTenantGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyTenantGroupsPartialUpdate(ctx, id).PatchedWritableTenantGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("tenant-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.TenancyAPI.TenancyTenantGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// tenantsCmd -------------------------------------------------------

func tenantsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tenants", Short: "Manage tenants"}
	cmd.AddCommand(
		cmdutil.ListCmd("tenants", func(ctx context.Context, client *netbox.APIClient, limit int32) error {
			resp, _, err := client.TenancyAPI.TenancyTenantsList(ctx).Limit(limit).Execute()
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
		cmdutil.CreateCmd("tenant", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.TenantRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyTenantsCreate(ctx).TenantRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("tenant", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedTenantRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.TenancyAPI.TenancyTenantsPartialUpdate(ctx, id).PatchedTenantRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("tenant", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.TenancyAPI.TenancyTenantsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
