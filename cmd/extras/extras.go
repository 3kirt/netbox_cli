// Package extras implements CLI subcommands for the NetBox Extras API.
package extras

import (
	"context"

	netbox "github.com/netbox-community/go-netbox/v4"
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/cmdutil"
)

// Command returns the cobra command tree for the Extras API area.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extras",
		Short: "Commands for the NetBox Extras API",
	}

	cmd.AddCommand(
		bookmarksCmd(),
		configContextsCmd(),
		configTemplatesCmd(),
		customFieldChoiceSetsCmd(),
		customFieldsCmd(),
		customLinksCmd(),
		eventRulesCmd(),
		exportTemplatesCmd(),
		imageAttachmentsCmd(),
		journalEntriesCmd(),
		notificationGroupsCmd(),
		notificationsCmd(),
		objectTypesCmd(),
		savedFiltersCmd(),
		scriptsCmd(),
		subscriptionsCmd(),
		tableConfigsCmd(),
		taggedObjectsCmd(),
		tagsCmd(),
		webhooksCmd(),
	)

	return cmd
}

// bookmarksCmd -------------------------------------------------------

func bookmarksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "bookmarks", Short: "Manage bookmarks"}
	cmd.AddCommand(
		cmdutil.ListCmd("bookmarks", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasBookmarksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("bookmark", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasBookmarksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// configContextsCmd -------------------------------------------------------

func configContextsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "config-contexts", Short: "Manage config contexts"}
	cmd.AddCommand(
		cmdutil.ListCmd("config-contexts", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasConfigContextsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("config-context", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasConfigContextsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// configTemplatesCmd -------------------------------------------------------

func configTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "config-templates", Short: "Manage config templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("config-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasConfigTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("config-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasConfigTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// customFieldChoiceSetsCmd -------------------------------------------------------

func customFieldChoiceSetsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "custom-field-choice-sets", Short: "Manage custom field choice sets"}
	cmd.AddCommand(
		cmdutil.ListCmd("custom-field-choice-sets", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("custom-field-choice-set", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// customFieldsCmd -------------------------------------------------------

func customFieldsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "custom-fields", Short: "Manage custom fields"}
	cmd.AddCommand(
		cmdutil.ListCmd("custom-fields", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("custom-field", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// customLinksCmd -------------------------------------------------------

func customLinksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "custom-links", Short: "Manage custom links"}
	cmd.AddCommand(
		cmdutil.ListCmd("custom-links", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasCustomLinksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("custom-link", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasCustomLinksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// eventRulesCmd -------------------------------------------------------

func eventRulesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "event-rules", Short: "Manage event rules"}
	cmd.AddCommand(
		cmdutil.ListCmd("event-rules", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasEventRulesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("event-rule", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasEventRulesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// exportTemplatesCmd -------------------------------------------------------

func exportTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "export-templates", Short: "Manage export templates"}
	cmd.AddCommand(
		cmdutil.ListCmd("export-templates", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasExportTemplatesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("export-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasExportTemplatesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// imageAttachmentsCmd -------------------------------------------------------

func imageAttachmentsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "image-attachments", Short: "Manage image attachments"}
	cmd.AddCommand(
		cmdutil.ListCmd("image-attachments", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasImageAttachmentsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("image-attachment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasImageAttachmentsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// journalEntriesCmd -------------------------------------------------------

func journalEntriesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "journal-entries", Short: "Manage journal entries"}
	cmd.AddCommand(
		cmdutil.ListCmd("journal-entries", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasJournalEntriesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("journal-entry", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasJournalEntriesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// notificationGroupsCmd -------------------------------------------------------

func notificationGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "notification-groups", Short: "Manage notification groups"}
	cmd.AddCommand(
		cmdutil.ListCmd("notification-groups", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasNotificationGroupsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("notification-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasNotificationGroupsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// notificationsCmd -------------------------------------------------------

func notificationsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "notifications", Short: "Manage notifications"}
	cmd.AddCommand(
		cmdutil.ListCmd("notifications", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasNotificationsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("notification", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasNotificationsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// objectTypesCmd -------------------------------------------------------

func objectTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "object-types", Short: "Manage object types"}
	cmd.AddCommand(
		cmdutil.ListCmd("object-types", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasObjectTypesList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("object-type", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasObjectTypesRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// savedFiltersCmd -------------------------------------------------------

func savedFiltersCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "saved-filters", Short: "Manage saved filters"}
	cmd.AddCommand(
		cmdutil.ListCmd("saved-filters", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasSavedFiltersList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("saved-filter", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasSavedFiltersRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// scriptsCmd -------------------------------------------------------
// Note: ExtrasScriptsRetrieve takes a string id, not int32, so only list is supported.

func scriptsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "scripts", Short: "Manage scripts"}
	cmd.AddCommand(
		cmdutil.ListCmd("scripts", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasScriptsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
	)
	return cmd
}

// subscriptionsCmd -------------------------------------------------------

func subscriptionsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "subscriptions", Short: "Manage subscriptions"}
	cmd.AddCommand(
		cmdutil.ListCmd("subscriptions", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasSubscriptionsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("subscription", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasSubscriptionsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tableConfigsCmd -------------------------------------------------------

func tableConfigsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "table-configs", Short: "Manage table configs"}
	cmd.AddCommand(
		cmdutil.ListCmd("table-configs", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasTableConfigsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("table-config", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasTableConfigsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// taggedObjectsCmd -------------------------------------------------------

func taggedObjectsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tagged-objects", Short: "Manage tagged objects"}
	cmd.AddCommand(
		cmdutil.ListCmd("tagged-objects", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasTaggedObjectsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tagged-object", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasTaggedObjectsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// tagsCmd -------------------------------------------------------

func tagsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tags", Short: "Manage tags"}
	cmd.AddCommand(
		cmdutil.ListCmd("tags", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasTagsList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tag", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasTagsRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}

// webhooksCmd -------------------------------------------------------

func webhooksCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "webhooks", Short: "Manage webhooks"}
	cmd.AddCommand(
		cmdutil.ListCmd("webhooks", func(ctx context.Context, client *netbox.APIClient) error {
			resp, _, err := client.ExtrasAPI.ExtrasWebhooksList(ctx).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("webhook", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			resp, _, err := client.ExtrasAPI.ExtrasWebhooksRetrieve(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
