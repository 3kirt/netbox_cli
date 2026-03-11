# netbox-cli

A command-line tool for reading data from a [NetBox](https://netbox.dev) instance via its REST API. It outputs JSON, making it easy to pipe results into other tools.

This project is primarily written by [Claude](https://claude.ai) and intended for use by Claude as a tool for retrieving network inventory information as context.

## Installation

**From source** (requires Go 1.22+):

```bash
git clone https://github.com/3kirt/netbox_cli.git
cd netbox_cli
make build
```

This produces a `netbox-cli` binary in the current directory.

**Docker:**

```bash
make docker
docker run --rm -e NETBOX_TOKEN=your-token netbox-cli --config /path/to/config.json ipam ip-addresses list
```

## Configuration

Create `~/.netbox_cli.json` with your NetBox URL:

```json
{
  "url": "https://netbox.example.com"
}
```

Set your API token as an environment variable:

```bash
export NETBOX_TOKEN=your-token-here
```

The token can also be stored in the config file as `"token"`, but the environment variable takes priority. You can point to a different config file with the `--config` flag.

## Usage

Every command follows the same pattern:

```
netbox-cli [--config FILE] <app> <resource> <action> [flags]
```

- `<app>` is a NetBox application (`circuits`, `core`, `dcim`, `extras`, `ipam`, `tenancy`, `users`, `virtualization`, `vpn`, `wireless`)
- `<resource>` is the object type (e.g. `ip-addresses`, `devices`, `tunnels`)
- `<action>` is one of:

| Action | Flags | Description |
|---|---|---|
| `list` | `--limit <N>` | Fetch records (default 0: all) |
| `get` | `--id <ID>` | Fetch one record by ID |
| `create` | | Create a record; reads JSON body from stdin |
| `update` | `--id <ID>` | Partially update a record (PATCH); reads JSON from stdin |
| `delete` | `--id <ID>` | Delete a record |

Some resources accept additional filter flags on `list` and `get`. See the per-area docs linked below for details.

**Examples:**

```bash
# List all IP addresses
netbox-cli ipam ip-addresses list

# Get a single prefix by ID
netbox-cli ipam prefixes get --id 42

# Create an IP address
echo '{"address": "10.0.1.1/24", "status": "active"}' \
  | netbox-cli ipam ip-addresses create

# Update a device's serial number
echo '{"serial": "FDO2142A0BC"}' \
  | netbox-cli dcim devices update --id 100

# Delete a circuit
netbox-cli circuits circuits delete --id 7

# List all virtual machines
netbox-cli virtualization virtual-machines list

# List virtual machines at a site, filtered by status
netbox-cli virtualization virtual-machines list --site lon01 --status active

# List IP addresses assigned to a VM
netbox-cli ipam ip-addresses list --virtual-machine web-01
```

Output is JSON. Use [`jq`](https://jqlang.org) to filter or format it:

```bash
netbox-cli ipam ip-addresses list | jq '.[].address'
```

## Available resources

Full command references with examples are in the [docs/](docs/) folder.

### `circuits` — [docs/circuits.md](docs/circuits.md)

| Resource | Description |
|---|---|
| `circuit-group-assignments` | Circuit group assignments |
| `circuit-groups` | Circuit groups |
| `circuit-terminations` | Circuit terminations |
| `circuit-types` | Circuit types |
| `circuits` | Circuits |
| `provider-accounts` | Provider accounts |
| `provider-networks` | Provider networks |
| `providers` | Providers |
| `virtual-circuit-terminations` | Virtual circuit terminations |
| `virtual-circuit-types` | Virtual circuit types |
| `virtual-circuits` | Virtual circuits |

### `core` — [docs/core.md](docs/core.md)

| Resource | Description |
|---|---|
| `data-files` | Data files |
| `data-sources` | Data sources |
| `jobs` | Background jobs |
| `object-changes` | Object change log |

### `dcim` — [docs/dcim.md](docs/dcim.md)

| Resource | Description |
|---|---|
| `cable-terminations` | Cable terminations |
| `cables` | Cables |
| `console-port-templates` | Console port templates |
| `console-ports` | Console ports |
| `console-server-port-templates` | Console server port templates |
| `console-server-ports` | Console server ports |
| `device-bay-templates` | Device bay templates |
| `device-bays` | Device bays |
| `device-roles` | Device roles |
| `device-types` | Device types |
| `devices` | Devices |
| `front-port-templates` | Front port templates |
| `front-ports` | Front ports |
| `interface-templates` | Interface templates |
| `interfaces` | Device interfaces |
| `inventory-item-roles` | Inventory item roles |
| `inventory-item-templates` | Inventory item templates |
| `inventory-items` | Inventory items |
| `locations` | Locations |
| `mac-addresses` | MAC addresses |
| `manufacturers` | Manufacturers |
| `module-bay-templates` | Module bay templates |
| `module-bays` | Module bays |
| `module-type-profiles` | Module type profiles |
| `module-types` | Module types |
| `modules` | Modules |
| `platforms` | Platforms |
| `power-feeds` | Power feeds |
| `power-outlet-templates` | Power outlet templates |
| `power-outlets` | Power outlets |
| `power-panels` | Power panels |
| `power-port-templates` | Power port templates |
| `power-ports` | Power ports |
| `rack-reservations` | Rack reservations |
| `rack-roles` | Rack roles |
| `rack-types` | Rack types |
| `racks` | Racks |
| `rear-port-templates` | Rear port templates |
| `rear-ports` | Rear ports |
| `regions` | Regions |
| `site-groups` | Site groups |
| `sites` | Sites |
| `virtual-chassis` | Virtual chassis |
| `virtual-device-contexts` | Virtual device contexts |

### `extras` — [docs/extras.md](docs/extras.md)

| Resource | Description |
|---|---|
| `bookmarks` | Bookmarks |
| `config-contexts` | Config contexts |
| `config-templates` | Config templates |
| `custom-field-choice-sets` | Custom field choice sets |
| `custom-fields` | Custom fields |
| `custom-links` | Custom links |
| `event-rules` | Event rules |
| `export-templates` | Export templates |
| `image-attachments` | Image attachments |
| `journal-entries` | Journal entries |
| `notification-groups` | Notification groups |
| `notifications` | Notifications |
| `object-types` | Object types |
| `saved-filters` | Saved filters |
| `scripts` | Scripts (list only) |
| `subscriptions` | Subscriptions |
| `table-configs` | Table configs |
| `tagged-objects` | Tagged objects |
| `tags` | Tags |
| `webhooks` | Webhooks |

### `ipam` — [docs/ipam.md](docs/ipam.md)

| Resource | Description |
|---|---|
| `aggregates` | IP aggregates |
| `asns` | Autonomous system numbers |
| `asn-ranges` | ASN ranges |
| `fhrp-groups` | First-hop redundancy protocol groups |
| `fhrp-group-assignments` | FHRP group assignments |
| `ip-addresses` | IP addresses |
| `ip-ranges` | IP ranges |
| `prefixes` | IP prefixes |
| `rirs` | Regional internet registries |
| `roles` | IP/VLAN roles |
| `route-targets` | VRF route targets |
| `service-templates` | Service templates |
| `services` | Services |
| `vlan-groups` | VLAN groups |
| `vlan-translation-policies` | VLAN translation policies |
| `vlan-translation-rules` | VLAN translation rules |
| `vlans` | VLANs |
| `vrfs` | Virtual routing and forwarding instances |

### `tenancy` — [docs/tenancy.md](docs/tenancy.md)

| Resource | Description |
|---|---|
| `contact-assignments` | Contact assignments |
| `contact-groups` | Contact groups |
| `contact-roles` | Contact roles |
| `contacts` | Contacts |
| `tenant-groups` | Tenant groups |
| `tenants` | Tenants |

### `users` — [docs/users.md](docs/users.md)

| Resource | Description |
|---|---|
| `groups` | Groups |
| `permissions` | Permissions |
| `tokens` | API tokens |
| `users` | Users |

### `virtualization` — [docs/virtualization.md](docs/virtualization.md)

| Resource | Description |
|---|---|
| `cluster-groups` | Cluster groups |
| `cluster-types` | Cluster types |
| `clusters` | Clusters |
| `interfaces` | Virtual machine interfaces |
| `virtual-disks` | Virtual disks |
| `virtual-machines` | Virtual machines |

### `vpn` — [docs/vpn.md](docs/vpn.md)

| Resource | Description |
|---|---|
| `ike-policies` | IKE policies |
| `ike-proposals` | IKE proposals |
| `ipsec-policies` | IPSec policies |
| `ipsec-profiles` | IPSec profiles |
| `ipsec-proposals` | IPSec proposals |
| `l2vpn-terminations` | L2VPN terminations |
| `l2vpns` | L2VPNs |
| `tunnel-groups` | Tunnel groups |
| `tunnel-terminations` | Tunnel terminations |
| `tunnels` | Tunnels |

### `wireless` — [docs/wireless.md](docs/wireless.md)

| Resource | Description |
|---|---|
| `wireless-lan-groups` | Wireless LAN groups |
| `wireless-lans` | Wireless LANs |
| `wireless-links` | Wireless links |
