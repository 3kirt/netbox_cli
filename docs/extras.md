# extras

Commands for the NetBox Extras API — tags, custom fields, config contexts, webhooks, scripts, and more.

## Actions

| Action | Flags | Description |
|---|---|---|
| `list` | | Fetch all records |
| `get` | `--id <ID>` | Fetch one record by ID |
| `create` | | Create a record; reads JSON body from stdin |
| `update` | `--id <ID>` | Partially update a record (PATCH); reads JSON body from stdin |
| `delete` | `--id <ID>` | Delete a record |

## Resources

| Resource | Description | Actions |
|---|---|---|
| `bookmarks` | User bookmarks | list, get, create, update, delete |
| `config-contexts` | Config context definitions | list, get, create, update, delete |
| `config-templates` | Jinja2 config templates | list, get, create, update, delete |
| `custom-field-choice-sets` | Choice sets for custom select fields | list, get, create, update, delete |
| `custom-fields` | Custom field definitions | list, get, create, update, delete |
| `custom-links` | Custom link definitions | list, get, create, update, delete |
| `event-rules` | Event-driven rules | list, get, create, update, delete |
| `export-templates` | Export template definitions | list, get, create, update, delete |
| `image-attachments` | Image attachments on objects | list, get, create, update, delete |
| `journal-entries` | Journal entries on objects | list, get, create, update, delete |
| `notification-groups` | Notification groups | list, get, create, update, delete |
| `notifications` | Notifications | list, get, create, update, delete |
| `object-types` | NetBox object type registry | list, get |
| `saved-filters` | Saved search filters | list, get, create, update, delete |
| `scripts` | Custom scripts | list |
| `subscriptions` | User notification subscriptions | list, get, create, update, delete |
| `table-configs` | User table column configurations | list, get, create, update, delete |
| `tagged-objects` | Objects associated with a tag | list, get |
| `tags` | Tags | list, get, create, update, delete |
| `webhooks` | Webhook definitions | list, get, create, update, delete |

`object-types` and `tagged-objects` are read-only. `scripts` uses a non-integer ID and only supports `list`.

## Examples

```bash
# List all tags
netbox-cli extras tags list

# Create a tag
echo '{"name": "Production", "slug": "production", "color": "ff0000"}' \
  | netbox-cli extras tags create

# Update a tag's color
echo '{"color": "00ff00"}' \
  | netbox-cli extras tags update --id 5

# Delete a tag
netbox-cli extras tags delete --id 5

# Create a custom field
echo '{
  "object_types": ["dcim.device"],
  "type": "text",
  "name": "asset_tag",
  "label": "Asset Tag"
}' | netbox-cli extras custom-fields create

# Create a webhook
echo '{
  "name": "Deploy Hook",
  "payload_url": "https://ci.example.com/hook",
  "http_method": "POST"
}' | netbox-cli extras webhooks create

# Create a config context
echo '{
  "name": "NTP Servers",
  "weight": 1000,
  "data": {"ntp": ["10.0.0.1", "10.0.0.2"]}
}' | netbox-cli extras config-contexts create

# List all scripts
netbox-cli extras scripts list

# Browse object types
netbox-cli extras object-types list | jq '.[].model'
```
