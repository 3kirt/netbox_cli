// Package extras implements CLI subcommands for the NetBox Extras API.
package extras

import (
	"context"
	"encoding/json"
	"fmt"

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
		cmdutil.CreateCmd("bookmark", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.BookmarkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasBookmarksCreate(ctx).BookmarkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("bookmark", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedBookmarkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasBookmarksPartialUpdate(ctx, id).PatchedBookmarkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("bookmark", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasBookmarksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("config-context", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ConfigContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigContextsCreate(ctx).ConfigContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("config-context", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedConfigContextRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigContextsPartialUpdate(ctx, id).PatchedConfigContextRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("config-context", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasConfigContextsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("config-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ConfigTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigTemplatesCreate(ctx).ConfigTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("config-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedConfigTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasConfigTemplatesPartialUpdate(ctx, id).PatchedConfigTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("config-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasConfigTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("custom-field-choice-set", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableCustomFieldChoiceSetRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsCreate(ctx).WritableCustomFieldChoiceSetRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("custom-field-choice-set", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableCustomFieldChoiceSetRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsPartialUpdate(ctx, id).PatchedWritableCustomFieldChoiceSetRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("custom-field-choice-set", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasCustomFieldChoiceSetsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("custom-field", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableCustomFieldRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldsCreate(ctx).WritableCustomFieldRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("custom-field", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableCustomFieldRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomFieldsPartialUpdate(ctx, id).PatchedWritableCustomFieldRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("custom-field", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasCustomFieldsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("custom-link", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.CustomLinkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomLinksCreate(ctx).CustomLinkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("custom-link", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedCustomLinkRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasCustomLinksPartialUpdate(ctx, id).PatchedCustomLinkRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("custom-link", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasCustomLinksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("event-rule", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableEventRuleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasEventRulesCreate(ctx).WritableEventRuleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("event-rule", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableEventRuleRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasEventRulesPartialUpdate(ctx, id).PatchedWritableEventRuleRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("event-rule", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasEventRulesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("export-template", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ExportTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasExportTemplatesCreate(ctx).ExportTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("export-template", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedExportTemplateRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasExportTemplatesPartialUpdate(ctx, id).PatchedExportTemplateRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("export-template", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasExportTemplatesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("image-attachment", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.ImageAttachmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasImageAttachmentsCreate(ctx).ImageAttachmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("image-attachment", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedImageAttachmentRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasImageAttachmentsPartialUpdate(ctx, id).PatchedImageAttachmentRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("image-attachment", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasImageAttachmentsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("journal-entry", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WritableJournalEntryRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasJournalEntriesCreate(ctx).WritableJournalEntryRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("journal-entry", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWritableJournalEntryRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasJournalEntriesPartialUpdate(ctx, id).PatchedWritableJournalEntryRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("journal-entry", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasJournalEntriesDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("notification-group", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.NotificationGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationGroupsCreate(ctx).NotificationGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("notification-group", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedNotificationGroupRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationGroupsPartialUpdate(ctx, id).PatchedNotificationGroupRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("notification-group", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasNotificationGroupsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("notification", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.NotificationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationsCreate(ctx).NotificationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("notification", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedNotificationRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasNotificationsPartialUpdate(ctx, id).PatchedNotificationRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("notification", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasNotificationsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// objectTypesCmd -------------------------------------------------------
// Note: read-only endpoint — no create, update, or delete.

func objectTypesCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "object-types", Short: "Browse object types"}
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
		cmdutil.CreateCmd("saved-filter", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.SavedFilterRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasSavedFiltersCreate(ctx).SavedFilterRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("saved-filter", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedSavedFilterRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasSavedFiltersPartialUpdate(ctx, id).PatchedSavedFilterRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("saved-filter", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasSavedFiltersDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("subscription", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.SubscriptionRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasSubscriptionsCreate(ctx).SubscriptionRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("subscription", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedSubscriptionRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasSubscriptionsPartialUpdate(ctx, id).PatchedSubscriptionRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("subscription", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasSubscriptionsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("table-config", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.TableConfigRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasTableConfigsCreate(ctx).TableConfigRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("table-config", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedTableConfigRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasTableConfigsPartialUpdate(ctx, id).PatchedTableConfigRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("table-config", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasTableConfigsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}

// taggedObjectsCmd -------------------------------------------------------
// Note: read-only endpoint — no create, update, or delete.

func taggedObjectsCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "tagged-objects", Short: "Browse tagged objects"}
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
		cmdutil.CreateCmd("tag", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.TagRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasTagsCreate(ctx).TagRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("tag", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedTagRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasTagsPartialUpdate(ctx, id).PatchedTagRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("tag", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasTagsDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
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
		cmdutil.CreateCmd("webhook", func(ctx context.Context, client *netbox.APIClient, data []byte) error {
			var body netbox.WebhookRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasWebhooksCreate(ctx).WebhookRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.UpdateCmd("webhook", func(ctx context.Context, client *netbox.APIClient, id int32, data []byte) error {
			var body netbox.PatchedWebhookRequest
			if err := json.Unmarshal(data, &body); err != nil {
				return fmt.Errorf("invalid JSON: %w", err)
			}
			resp, _, err := client.ExtrasAPI.ExtrasWebhooksPartialUpdate(ctx, id).PatchedWebhookRequest(body).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return cmdutil.OutputJSON(resp)
		}),
		cmdutil.DeleteCmd("webhook", func(ctx context.Context, client *netbox.APIClient, id int32) error {
			_, err := client.ExtrasAPI.ExtrasWebhooksDestroy(ctx, id).Execute()
			if err != nil {
				return cmdutil.APIError(err)
			}
			return nil
		}),
	)
	return cmd
}
