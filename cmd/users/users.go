// Package users implements CLI subcommands for the NetBox Users API.
package users

import (
	"context"

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
	)
	return cmd
}
