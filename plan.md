# Implementation Plan: netbox-cli

## Overview

A Go command-line tool for querying the NetBox API, designed to provide inventory information as context for Claude. Built on the `go-netbox` library (OpenAPI-generated client), organized for incremental expansion across all NetBox API areas.

## Standards

- UTF-8 encoding, LF line endings, 2-space indentation, final newline — per `.editorconfig`

## Project Structure

```
netbox_cli/
├── specification.md
├── plan.md
├── .editorconfig
├── go.mod
├── go.sum
├── main.go                        # Entry point: calls cmd.Execute()
├── cmd/
│   ├── root.go                    # Root cobra command, config loading, client init
│   └── virtualization/
│       └── virtualization.go      # Virtualization subcommands
└── internal/
    └── config/
        └── config.go              # Config struct and JSON loading
```

The `cmd/` directory is the primary extension point. Adding a new NetBox API area (e.g., dcim, ipam) means adding a new subdirectory under `cmd/` and registering it in `cmd/root.go`.

## Dependencies

- `github.com/spf13/cobra` — CLI framework
- `github.com/netbox-community/go-netbox/v4` — NetBox API client (local checkout at `/Users/kirtis/source/repos/go-netbox`)

## Configuration

### Config File

Default location: `~/.netbox_cli.json`. Overridable via `--config FILE`.

```json
{
  "url": "https://netbox.example.com",
  "token": "optional-fallback-token"
}
```

### Token Resolution

Resolved in this priority order:

1. `NETBOX_TOKEN` environment variable
2. `token` key in the JSON config file

If no token is found, the tool exits with a clear error message on stderr.

## Client Initialization

The go-netbox library provides a convenience constructor:

```go
client := netbox.NewAPIClientFor(url, token)
```

This sets the `Authorization: Token <token>` header and `Accept-Language: en-US` on all requests. The root command initializes the client once and passes it to subcommands via cobra's context.

## CLI Design

```
netbox-cli [--config FILE] <app> <endpoint> <action> [FLAGS]
```

### Global Flags

| Flag | Description |
|------|-------------|
| `--config FILE` | Path to JSON config file (default: `~/.netbox_cli.json`) |

### Actions

Each endpoint exposes two subcommands:

| Action | Description |
|--------|-------------|
| `list` | Retrieve all records. Accepts `--filter key=value` flags (repeatable). |
| `get` | Retrieve a single record by `--id ID`. |

### Virtualization Commands (Initial Scope)

```
netbox-cli virtualization cluster-types list [--filter key=value ...]
netbox-cli virtualization cluster-types get --id ID

netbox-cli virtualization cluster-groups list [--filter key=value ...]
netbox-cli virtualization cluster-groups get --id ID

netbox-cli virtualization clusters list [--filter key=value ...]
netbox-cli virtualization clusters get --id ID

netbox-cli virtualization virtual-machines list [--filter key=value ...]
netbox-cli virtualization virtual-machines get --id ID

netbox-cli virtualization interfaces list [--filter key=value ...]
netbox-cli virtualization interfaces get --id ID
```

go-netbox API calls follow the pattern:

```go
// list
result, _, err := client.VirtualizationAPI.VirtualizationVirtualMachinesList(ctx).Execute()

// get by ID
result, _, err := client.VirtualizationAPI.VirtualizationVirtualMachinesRetrieve(ctx, int32(id)).Execute()
```

## Output

All output is JSON written to stdout. Errors are written to stderr with a non-zero exit code. The go-netbox response structs are marshalled directly to JSON.

## Expansion Path

When adding a new NetBox API area:

1. Create `cmd/<area>/<area>.go` with a cobra command group and endpoint subcommands.
2. Register the group in `cmd/root.go` with `rootCmd.AddCommand(...)`.

Config loading, client initialization, and output formatting are shared via the root command — no changes needed to those layers.
