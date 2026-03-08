# circuits

Commands for the NetBox Circuits API — providers, circuits, terminations, and virtual circuits.

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
| `circuit-group-assignments` | Assign a circuit to a circuit group | list, get, create, update, delete |
| `circuit-groups` | Circuit groups | list, get, create, update, delete |
| `circuit-terminations` | Circuit terminations (A-side / Z-side) | list, get, create, update, delete |
| `circuit-types` | Circuit types | list, get, create, update, delete |
| `circuits` | Circuits | list, get, create, update, delete |
| `provider-accounts` | Provider accounts | list, get, create, update, delete |
| `provider-networks` | Provider networks | list, get, create, update, delete |
| `providers` | Circuit providers | list, get, create, update, delete |
| `virtual-circuit-terminations` | Virtual circuit terminations | list, get, create, update, delete |
| `virtual-circuit-types` | Virtual circuit types | list, get, create, update, delete |
| `virtual-circuits` | Virtual circuits | list, get, create, update, delete |

## Examples

```bash
# List all providers
netbox-cli circuits providers list

# Get a single circuit by ID
netbox-cli circuits circuits get --id 7

# Create a provider
echo '{"name": "Acme ISP", "slug": "acme-isp"}' \
  | netbox-cli circuits providers create

# Create a circuit type
echo '{"name": "Dark Fibre", "slug": "dark-fibre"}' \
  | netbox-cli circuits circuit-types create

# Create a circuit (requires provider ID and circuit type ID)
echo '{"cid": "CKT-001", "provider": 1, "type": 2}' \
  | netbox-cli circuits circuits create

# Update a circuit's status
echo '{"status": "active"}' \
  | netbox-cli circuits circuits update --id 7

# Delete a circuit
netbox-cli circuits circuits delete --id 7

# Create a provider network
echo '{"name": "Backbone-West", "provider": 1}' \
  | netbox-cli circuits provider-networks create

# List virtual circuits and filter with jq
netbox-cli circuits virtual-circuits list | jq '.[] | {id, cid, status}'
```
