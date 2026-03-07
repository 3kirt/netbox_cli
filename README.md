# netbox-cli

A command-line interface for the [NetBox](https://netbox.dev) API, primarily written by [Claude](https://claude.ai) and intended for use by Claude as a tool for retrieving network inventory information as context.

## Requirements

- Go 1.22+
- A NetBox instance with API access

## Installation

```bash
git clone https://github.com/3kirt/netbox_cli.git
cd netbox_cli
make build
```

This produces a `netbox-cli` binary in the current directory.

## Configuration

Create a config file at `~/.netbox_cli.json`:

```json
{
  "url": "https://netbox.example.com"
}
```

Set your API token in the environment:

```bash
export NETBOX_TOKEN=your-token-here
```

The token can also be stored in the config file as `"token"`, but the environment variable takes priority.

A different config file can be specified with the `--config` flag:

```bash
netbox-cli --config /path/to/config.json virtualization clusters list
```

## Usage

```
netbox-cli [--config FILE] <app> <endpoint> <action> [flags]
```

All output is JSON written to stdout.

### Virtualization

```bash
netbox-cli virtualization cluster-groups list
netbox-cli virtualization cluster-groups get --id 1

netbox-cli virtualization cluster-types list
netbox-cli virtualization cluster-types get --id 1

netbox-cli virtualization clusters list
netbox-cli virtualization clusters get --id 1

netbox-cli virtualization interfaces list
netbox-cli virtualization interfaces get --id 1

netbox-cli virtualization virtual-disks list
netbox-cli virtualization virtual-disks get --id 1

netbox-cli virtualization virtual-machines list
netbox-cli virtualization virtual-machines get --id 1
```

### IPAM

```bash
netbox-cli ipam aggregates list
netbox-cli ipam aggregates get --id 1

netbox-cli ipam asns list
netbox-cli ipam asns get --id 1

netbox-cli ipam asn-ranges list
netbox-cli ipam asn-ranges get --id 1

netbox-cli ipam fhrp-groups list
netbox-cli ipam fhrp-groups get --id 1

netbox-cli ipam fhrp-group-assignments list
netbox-cli ipam fhrp-group-assignments get --id 1

netbox-cli ipam ip-addresses list
netbox-cli ipam ip-addresses get --id 1

netbox-cli ipam ip-ranges list
netbox-cli ipam ip-ranges get --id 1

netbox-cli ipam prefixes list
netbox-cli ipam prefixes get --id 1

netbox-cli ipam rirs list
netbox-cli ipam rirs get --id 1

netbox-cli ipam roles list
netbox-cli ipam roles get --id 1

netbox-cli ipam route-targets list
netbox-cli ipam route-targets get --id 1

netbox-cli ipam service-templates list
netbox-cli ipam service-templates get --id 1

netbox-cli ipam services list
netbox-cli ipam services get --id 1

netbox-cli ipam vlan-groups list
netbox-cli ipam vlan-groups get --id 1

netbox-cli ipam vlan-translation-policies list
netbox-cli ipam vlan-translation-policies get --id 1

netbox-cli ipam vlan-translation-rules list
netbox-cli ipam vlan-translation-rules get --id 1

netbox-cli ipam vlans list
netbox-cli ipam vlans get --id 1

netbox-cli ipam vrfs list
netbox-cli ipam vrfs get --id 1
```

## Project Structure

```
netbox_cli/
├── main.go                              # Entry point
├── cmd/
│   ├── root.go                          # Root command, config and client setup
│   ├── ipam/
│   │   └── ipam.go                      # IPAM subcommands
│   └── virtualization/
│       └── virtualization.go            # Virtualization subcommands
└── internal/
    ├── clientctx/
    │   └── clientctx.go                 # Passes API client through cobra context
    ├── cmdutil/
    │   └── cmdutil.go                   # Shared helpers: OutputJSON, APIError, ListCmd, GetCmd
    └── config/
        └── config.go                    # Config file loading and token resolution
```

New NetBox API areas can be added by creating a package under `cmd/` and registering it in `cmd/root.go`.
