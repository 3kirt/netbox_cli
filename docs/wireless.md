# wireless

Commands for the NetBox Wireless API — wireless LANs, wireless LAN groups, and wireless links.

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
| `wireless-lan-groups` | Wireless LAN groups | list, get, create, update, delete |
| `wireless-lans` | Wireless LANs (SSIDs) | list, get, create, update, delete |
| `wireless-links` | Wireless links between interfaces | list, get, create, update, delete |

## Examples

```bash
# List all wireless LANs
netbox-cli wireless wireless-lans list

# Get a wireless LAN by ID
netbox-cli wireless wireless-lans get --id 5

# Create a wireless LAN group
echo '{"name": "Corporate WiFi", "slug": "corporate-wifi"}' \
  | netbox-cli wireless wireless-lan-groups create

# Create a wireless LAN (SSID)
echo '{
  "ssid": "Corp-Secure",
  "auth_type": "wpa-enterprise",
  "group": 1
}' | netbox-cli wireless wireless-lans create

# Update a wireless LAN's description
echo '{"description": "Primary corporate SSID"}' \
  | netbox-cli wireless wireless-lans update --id 5

# Delete a wireless LAN
netbox-cli wireless wireless-lans delete --id 5

# Create a wireless link between two interfaces
echo '{
  "interface_a": 10,
  "interface_b": 20,
  "status": "active",
  "ssid": "Corp-Backhaul"
}' | netbox-cli wireless wireless-links create

# List all wireless links and show interface names
netbox-cli wireless wireless-links list \
  | jq '.[] | {id, a: .interface_a.name, b: .interface_b.name, status: .status.value}'
```
