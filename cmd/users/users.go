// Package users implements CLI subcommands for the NetBox Users API.
package users

import (
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
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
		cmdutil.ListCmd("groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("permissions", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersPermissionsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("permission", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersPermissionsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("tokens", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersTokensList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("token", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersTokensRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("users", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersUsersList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("user", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.UsersAPI.UsersUsersRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
