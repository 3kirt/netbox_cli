# ipam

Commands for the NetBox IPAM API — prefixes, IP addresses, VLANs, VRFs, ASNs, and more.

## Actions

| Action | Flags | Description |
|---|---|---|
| `list` | | Fetch all records |
| `get` | `--id <ID>` | Fetch one record by ID |
| `create` | | Create a record; reads JSON body from stdin |
| `update` | `--id <ID>` | Partially update a record (PATCH); reads JSON body from stdin |
| `delete` | `--id <ID>` | Delete a record |

## Resource-specific flags

### `prefixes`

`list` accepts optional filter flags. Omitting all flags returns all records.

| Flag | Description |
|---|---|
| `--search <text>` | Free-text search (prefix string or description) |
| `--site <slug>` | Filter by site slug |

### `ip-addresses`

`list` accepts optional filter flags. Omitting all flags returns all records.

| Flag | Description |
|---|---|
| `--virtual-machine <name>` | Filter by virtual machine name |
| `--device <name>` | Filter by device name |
| `--address <address>` | Filter by IP address (e.g. `192.0.2.1/24`) |

## Resources

| Resource | Description | Actions |
|---|---|---|
| `aggregates` | IP aggregates | list, get, create, update, delete |
| `asn-ranges` | ASN ranges | list, get, create, update, delete |
| `asns` | Autonomous system numbers | list, get, create, update, delete |
| `fhrp-group-assignments` | FHRP group assignments | list, get, create, update, delete |
| `fhrp-groups` | First-hop redundancy protocol groups | list, get, create, update, delete |
| `ip-addresses` | IP addresses | list, get, create, update, delete |
| `ip-ranges` | IP ranges | list, get, create, update, delete |
| `prefixes` | IP prefixes | list, get, create, update, delete |
| `rirs` | Regional internet registries | list, get, create, update, delete |
| `roles` | IP/VLAN roles | list, get, create, update, delete |
| `route-targets` | VRF route targets | list, get, create, update, delete |
| `service-templates` | Service templates | list, get, create, update, delete |
| `services` | Services attached to devices or VMs | list, get, create, update, delete |
| `vlan-groups` | VLAN groups | list, get, create, update, delete |
| `vlan-translation-policies` | VLAN translation policies | list, get, create, update, delete |
| `vlan-translation-rules` | VLAN translation rules | list, get, create, update, delete |
| `vlans` | VLANs | list, get, create, update, delete |
| `vrfs` | Virtual routing and forwarding instances | list, get, create, update, delete |

## Examples

```bash
# List all prefixes
netbox-cli ipam prefixes list

# Get an IP address by ID
netbox-cli ipam ip-addresses get --id 42

# Create a prefix
echo '{"prefix": "10.0.0.0/8", "status": "active"}' \
  | netbox-cli ipam prefixes create

# Create an IP address
echo '{"address": "10.0.1.1/24", "status": "active"}' \
  | netbox-cli ipam ip-addresses create

# Assign an IP to a device interface
echo '{
  "address": "192.168.1.10/24",
  "status": "active",
  "assigned_object_type": "dcim.interface",
  "assigned_object_id": 101
}' | netbox-cli ipam ip-addresses create

# Update an IP address's DNS name
echo '{"dns_name": "gw.example.com"}' \
  | netbox-cli ipam ip-addresses update --id 42

# Delete an IP address
netbox-cli ipam ip-addresses delete --id 42

# Create a VLAN
echo '{"vid": 100, "name": "Management", "status": "active"}' \
  | netbox-cli ipam vlans create

# Create a VRF
echo '{"name": "MGMT", "rd": "65000:1"}' \
  | netbox-cli ipam vrfs create

# Create an ASN
echo '{"asn": 65001, "rir": 1}' \
  | netbox-cli ipam asns create

# List all VLANs and show just vid and name
netbox-cli ipam vlans list | jq '.[] | {id, vid, name}'

# Search prefixes by subnet string
netbox-cli ipam prefixes list --search 10.0

# List prefixes at a site
netbox-cli ipam prefixes list --site hq

# Find all active prefixes in a VRF (client-side filter still works for VRF)
netbox-cli ipam prefixes list | jq '.[] | select(.vrf.id == 3 and .status.value == "active")'

# List IP addresses assigned to a virtual machine
netbox-cli ipam ip-addresses list --virtual-machine web-01

# List IP addresses assigned to a device
netbox-cli ipam ip-addresses list --device core-sw-01

# Look up a specific IP address
netbox-cli ipam ip-addresses list --address 192.0.2.1/24
```
