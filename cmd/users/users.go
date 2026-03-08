// Package users implements CLI subcommands for the NetBox Users API.
package users

import (
	"context"
	"encoding/json"
	"fmt"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the Users API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Commands for the NetBox Users API",
	}

	cmd.AddCommand(
		groupsCmd(),
		permissionsCmd(),
		tokensCmd(),
		usersCmd(),
	)

	return cmd
}

// groupsCmd -------------------------------------------------------

func groupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "groups", Short: "Manage groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.UsersAPI.UsersGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.UsersAPI.UsersGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.GroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersGroupsCreate(ctx).GroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersGroupsPartialUpdate(ctx, id).PatchedGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.UsersAPI.UsersGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// permissionsCmd -------------------------------------------------------

func permissionsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "permissions", Short: "Manage permissions"}
	cmd.AddCommand(
		cmdutil.ListCmd("permissions", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.UsersAPI.UsersPermissionsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("permission", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.UsersAPI.UsersPermissionsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("permission", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ObjectPermissionRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersPermissionsCreate(ctx).ObjectPermissionRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("permission", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedObjectPermissionRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersPermissionsPartialUpdate(ctx, id).PatchedObjectPermissionRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("permission", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.UsersAPI.UsersPermissionsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// tokensCmd -------------------------------------------------------

func tokensCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tokens", Short: "Manage tokens"}
	cmd.AddCommand(
		cmdutil.ListCmd("tokens", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.UsersAPI.UsersTokensList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("token", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.UsersAPI.UsersTokensRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("token", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.TokenRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersTokensCreate(ctx).TokenRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("token", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedTokenRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersTokensPartialUpdate(ctx, id).PatchedTokenRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("token", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.UsersAPI.UsersTokensDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// usersCmd -------------------------------------------------------

func usersCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "users", Short: "Manage users"}
	cmd.AddCommand(
		cmdutil.ListCmd("users", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.UsersAPI.UsersUsersList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("user", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.UsersAPI.UsersUsersRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("user", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.UserRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersUsersCreate(ctx).UserRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("user", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedUserRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.UsersAPI.UsersUsersPartialUpdate(ctx, id).PatchedUserRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("user", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.UsersAPI.UsersUsersDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
