# netbox-cli — Claude Code Guide

## Project purpose

Command-line interface for the NetBox API, primarily written by Claude and
intended for use by Claude as a tool for retrieving network inventory
information as context.

## Build and run

```bash
make build          # produces ./netbox-cli binary
make docker         # builds Docker image tagged netbox-cli
make clean          # removes the binary
go vet ./...        # run static analysis
gofmt -l .          # check formatting (should produce no output)
gofmt -w .          # reformat all Go files
```

## Configuration

The CLI reads from `~/.netbox_cli.json` by default (override with `--config`):

```json
{ "url": "https://netbox.example.com", "token": "abc123" }
```

`NETBOX_TOKEN` env var takes precedence over the config file token.

## Project structure

```
main.go                              entry point — calls cmd.Execute()
cmd/
  root.go                            root cobra command, config + client init
  ipam/ipam.go                       IPAM API subcommands
  virtualization/virtualization.go   Virtualization API subcommands
internal/
  clientctx/clientctx.go            stores *netbox.APIClient in cobra context
  cmdutil/cmdutil.go                 shared helpers: OutputJSON, APIError, ListCmd, GetCmd
  config/config.go                   JSON config loading, token/URL resolution
```

## Branching

- All new features must be developed on a dedicated feature branch (`feature/<name>`).
- Branch from `main` and merge back to `main` when the work is complete.

## Adding a new API area

1. Create a feature branch: `git checkout -b feature/<area>`.
2. Create `cmd/<area>/<area>.go` — follow the pattern in `cmd/ipam/ipam.go`.
3. Register `<area>.Command()` in `cmd/root.go`.
4. Use `cmdutil.ListCmd` / `cmdutil.GetCmd` for standard list/get subcommands.
5. Retrieve the client with `clientctx.Client(cmd.Context())`.
6. Call the corresponding `client.<Area>API.<Area>XxxList(ctx).Limit(0).Execute()`.

## Code standards

- **Formatting**: `gofmt` (tabs, not spaces). `.editorconfig` has a `[*.go]` override.
- **Style**: Google Go Style Guide — lowercase error strings, no trailing punctuation,
  never silently discard errors without an explanatory comment.
- **Comments**: godoc on all exported symbols and packages.
- **Module**: `github.com/kirtis/netbox-cli`, Go 1.22.
- **go-netbox**: `github.com/netbox-community/go-netbox/v4 v4.3.0` (published module).
