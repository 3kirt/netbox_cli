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

- `<app>` is a NetBox application (`ipam`, `virtualization`)
- `<resource>` is the object type (e.g. `ip-addresses`, `clusters`)
- `<action>` is either `list` (all records) or `get --id <ID>` (one record)

**Examples:**

```bash
# List all IP addresses
netbox-cli ipam ip-addresses list

# Get a single prefix by ID
netbox-cli ipam prefixes get --id 42

# List all virtual machines
netbox-cli virtualization virtual-machines list
```

Output is JSON. Use [`jq`](https://jqlang.org) to filter or format it:

```bash
netbox-cli ipam ip-addresses list | jq '.[].address'
```

## Available resources

### `ipam`

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

### `virtualization`

| Resource | Description |
|---|---|
| `cluster-groups` | Cluster groups |
| `cluster-types` | Cluster types |
| `clusters` | Clusters |
| `interfaces` | Virtual machine interfaces |
| `virtual-disks` | Virtual disks |
| `virtual-machines` | Virtual machines |
