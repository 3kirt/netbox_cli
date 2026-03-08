# core

Commands for the NetBox Core API — data sources, data files, jobs, and the object change log.

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
| `data-files` | Files synced from a data source | list, get |
| `data-sources` | External data source definitions | list, get, create, update, delete |
| `jobs` | Background job records | list, get |
| `object-changes` | Audit log of object changes | list, get |

`data-files`, `jobs`, and `object-changes` are read-only endpoints managed by NetBox itself.

## Examples

```bash
# List all data sources
netbox-cli core data-sources list

# Get a single data source by ID
netbox-cli core data-sources get --id 3

# Create a data source (Git backend)
echo '{
  "name": "Network Configs",
  "type": "git",
  "source_url": "https://github.com/example/netconfigs.git"
}' | netbox-cli core data-sources create

# Update a data source's enabled state
echo '{"enabled": false}' \
  | netbox-cli core data-sources update --id 3

# Delete a data source
netbox-cli core data-sources delete --id 3

# Browse the change log
netbox-cli core object-changes list | jq '.[] | {id, time, action, object_type}'

# List background jobs
netbox-cli core jobs list | jq '.[] | {id, name, status}'
```
