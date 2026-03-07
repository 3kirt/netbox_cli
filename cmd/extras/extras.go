// Package extras implements CLI subcommands for the NetBox Extras API.
package extras

import (
	"github.com/spf13/cobra"

	"github.com/kirtis/netbox-cli/internal/clientctx"
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
		cmdutil.ListCmd("bookmarks", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasBookmarksList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("bookmark", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasBookmarksRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("config-contexts", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigContextsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("config-context", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigContextsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("config-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("config-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("custom-field-choice-sets", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("custom-field-choice-set", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("custom-fields", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("custom-field", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("custom-links", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomLinksList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("custom-link", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomLinksRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("event-rules", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasEventRulesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("event-rule", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasEventRulesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("export-templates", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasExportTemplatesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("export-template", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasExportTemplatesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("image-attachments", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasImageAttachmentsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("image-attachment", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasImageAttachmentsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("journal-entries", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasJournalEntriesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("journal-entry", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasJournalEntriesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("notification-groups", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationGroupsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("notification-group", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationGroupsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("notifications", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("notification", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("object-types", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasObjectTypesList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("object-type", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasObjectTypesRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("saved-filters", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasSavedFiltersList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("saved-filter", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasSavedFiltersRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("scripts", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasScriptsList(cmd.Context()).Limit(0).Execute()
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
		cmdutil.ListCmd("subscriptions", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasSubscriptionsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("subscription", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasSubscriptionsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("table-configs", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasTableConfigsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("table-config", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasTableConfigsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("tagged-objects", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasTaggedObjectsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tagged-object", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasTaggedObjectsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("tags", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasTagsList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("tag", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasTagsRetrieve(cmd.Context(), id).Execute()
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
		cmdutil.ListCmd("webhooks", func(cmd *cobra.Command) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasWebhooksList(cmd.Context()).Limit(0).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp.GetResults())
		}),
		cmdutil.GetCmd("webhook", func(cmd *cobra.Command, id int32) error {
			client, err := clientctx.Client(cmd.Context())
			if err != nil {
				return err
			}
			resp, _, err := client.ExtrasAPI.ExtrasWebhooksRetrieve(cmd.Context(), id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
	)
	return cmd
}
