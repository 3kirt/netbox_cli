# vpn

Commands for the NetBox VPN API — tunnels, IKE/IPSec policies, L2VPNs, and related resources.

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
| `ike-policies` | IKE policies | list, get, create, update, delete |
| `ike-proposals` | IKE proposals | list, get, create, update, delete |
| `ipsec-policies` | IPSec policies | list, get, create, update, delete |
| `ipsec-profiles` | IPSec profiles | list, get, create, update, delete |
| `ipsec-proposals` | IPSec proposals | list, get, create, update, delete |
| `l2vpn-terminations` | L2VPN terminations | list, get, create, update, delete |
| `l2vpns` | L2VPNs (VXLAN, MPLS, etc.) | list, get, create, update, delete |
| `tunnel-groups` | Tunnel groups | list, get, create, update, delete |
| `tunnel-terminations` | Tunnel terminations | list, get, create, update, delete |
| `tunnels` | Tunnels | list, get, create, update, delete |

## Examples

```bash
# List all tunnels
netbox-cli vpn tunnels list

# Get a tunnel by ID
netbox-cli vpn tunnels get --id 3

# Create a tunnel group
echo '{"name": "DC Interconnects", "slug": "dc-interconnects"}' \
  | netbox-cli vpn tunnel-groups create

# Create a tunnel
echo '{
  "name": "dc1-to-dc2",
  "status": "active",
  "encapsulation": "ipsec-transport",
  "tunnel_id": 1001
}' | netbox-cli vpn tunnels create

# Update a tunnel's status
echo '{"status": "planned"}' \
  | netbox-cli vpn tunnels update --id 3

# Delete a tunnel
netbox-cli vpn tunnels delete --id 3

# Create an IKE proposal
echo '{
  "name": "IKEv2-AES256",
  "authentication_algorithm": "hmac-sha256",
  "encryption_algorithm": "aes-256-cbc",
  "sa_lifetime": 86400
}' | netbox-cli vpn ike-proposals create

# Create an L2VPN
echo '{
  "name": "VXLAN-Prod",
  "slug": "vxlan-prod",
  "type": "vxlan",
  "identifier": 10001
}' | netbox-cli vpn l2vpns create

# List L2VPNs and show name and type
netbox-cli vpn l2vpns list | jq '.[] | {id, name, type: .type.value}'
```
