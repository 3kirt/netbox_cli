// Package core implements CLI subcommands for the NetBox Core API.
package core

import (
	"context"
	"encoding/json"
	"fmt"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the Core API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "core",
		Short: "Commands for the NetBox Core API",
	}

	cmd.AddCommand(
		dataFilesCmd(),
		dataSourcesCmd(),
		jobsCmd(),
		objectChangesCmd(),
	)

	return cmd
}

// dataFilesCmd -------------------------------------------------------

func dataFilesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "data-files", Short: "Manage data files"}
	cmd.AddCommand(
		cmdutil.ListCmd("data-files", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CoreAPI.CoreDataFilesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("data-file", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CoreAPI.CoreDataFilesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// dataSourcesCmd -------------------------------------------------------

func dataSourcesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "data-sources", Short: "Manage data sources"}
	cmd.AddCommand(
		cmdutil.ListCmd("data-sources", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CoreAPI.CoreDataSourcesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("data-source", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CoreAPI.CoreDataSourcesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.CreateCmd("data-source", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableDataSourceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CoreAPI.CoreDataSourcesCreate(ctx).WritableDataSourceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("data-source", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableDataSourceRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.CoreAPI.CoreDataSourcesPartialUpdate(ctx, id).PatchedWritableDataSourceRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("data-source", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.CoreAPI.CoreDataSourcesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// jobsCmd -------------------------------------------------------

func jobsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "jobs", Short: "Manage jobs"}
	cmd.AddCommand(
		cmdutil.ListCmd("jobs", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CoreAPI.CoreJobsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("job", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CoreAPI.CoreJobsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// objectChangesCmd -------------------------------------------------------

func objectChangesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "object-changes", Short: "Manage object changes"}
	cmd.AddCommand(
		cmdutil.ListCmd("object-changes", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.CoreAPI.CoreObjectChangesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("object-change", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.CoreAPI.CoreObjectChangesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
