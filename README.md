# netbox-cli

A command-line interface for the [NetBox](https://netbox.dev) API. Designed to retrieve inventory information for use as context with Claude and other tools.

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

### Virtualization

```bash
# Cluster types
netbox-cli virtualization cluster-types list
netbox-cli virtualization cluster-types get --id 1

# Cluster groups
netbox-cli virtualization cluster-groups list
netbox-cli virtualization cluster-groups get --id 1

# Clusters
netbox-cli virtualization clusters list
netbox-cli virtualization clusters get --id 1

# Virtual machines
netbox-cli virtualization virtual-machines list
netbox-cli virtualization virtual-machines get --id 1

# Interfaces
netbox-cli virtualization interfaces list
netbox-cli virtualization interfaces get --id 1
```

All output is JSON written to stdout.

## Project Structure

```
netbox_cli/
├── main.go                          # Entry point
├── cmd/
│   ├── root.go                      # Root command, config and client setup
│   └── virtualization/
│       └── virtualization.go        # Virtualization subcommands
└── internal/
    ├── clientctx/
    │   └── clientctx.go             # Passes API client through cobra context
    └── config/
        └── config.go                # Config file loading and token resolution
```

New NetBox API areas can be added by creating a package under `cmd/` and registering it in `cmd/root.go`.
