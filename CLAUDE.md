# netbox-cli — Claude Code Guide

## Project purpose

Command-line interface for the NetBox API, primarily written by Claude and
intended for use by Claude as a tool for retrieving network inventory
information as context.

## Build and run

```bash
make build          # produces ./netbox-cli binary
make lint           # run golangci-lint (includes go vet, staticcheck, and more)
make docker         # builds Docker image tagged netbox-cli
make clean          # removes the binary
gofmt -l .          # check formatting (should produce no output)
gofmt -w .          # reformat all Go files
```

## Configuration

The CLI reads from `~/.netbox_cli.json` by default (override with `--config`):

```json
{ "url": "https://netbox.example.com", "token": "abc123" }
```

Token resolution priority:

1. `NETBOX_TOKEN` environment variable
2. `token` key in the JSON config file

The client is initialised with `netbox.NewAPIClientFor(url, token)`, which sets
`Authorization: Token <token>` and `Accept-Language: en-US` on every request.

## Project structure

```
main.go                              entry point — calls cmd.Execute()
cmd/
  root.go                            root cobra command, config + client init
  circuits/circuits.go               Circuits API subcommands
  dcim/dcim.go                       DCIM API subcommands
  extras/extras.go                   Extras API subcommands
  ipam/ipam.go                       IPAM API subcommands
  tenancy/tenancy.go                 Tenancy API subcommands
  users/users.go                     Users API subcommands
  virtualization/virtualization.go   Virtualization API subcommands
  vpn/vpn.go                         VPN API subcommands
  wireless/wireless.go               Wireless API subcommands
internal/
  clientctx/clientctx.go            stores *netbox.APIClient in cobra context
  cmdutil/cmdutil.go                 shared helpers: OutputJSON, APIError, ListCmd, GetCmd, CreateCmd, UpdateCmd, DeleteCmd
  config/config.go                   JSON config loading, token/URL resolution
```

## Git workflow

1. Create a feature branch: `git checkout -b feature/<name>`.
2. Make changes, then verify before committing:
   ```bash
   gofmt -w .
   go build ./...
   make lint
   ```
3. Fix any issues, then commit with a conventional commit message.
4. Merge to `main` and push: `git checkout main && git merge feature/<name> && git push`.

## Branching

- All new features must be developed on a dedicated feature branch (`feature/<name>`).
- Branch from `main` and merge back to `main` when the work is complete.

## Adding a new API area

1. Create a feature branch: `git checkout -b feature/<area>`.
2. Create `cmd/<area>/<area>.go` — follow the pattern in `cmd/ipam/ipam.go`.
3. Register `<area>.Command()` in `cmd/root.go`.
4. Use `cmdutil.ListCmd` / `cmdutil.GetCmd` for list/get subcommands.
5. Use `cmdutil.CreateCmd` / `cmdutil.UpdateCmd` / `cmdutil.DeleteCmd` for write subcommands. Create and update callbacks receive `[]byte` (raw stdin JSON) and unmarshal into the appropriate `netbox.XxxRequest` / `netbox.PatchedXxxRequest` type. Both commands read JSON from stdin — pipe it in: `echo '{"name":"x"}' | netbox-cli <area> <resource>s create` or `echo '{"status":"staged"}' | netbox-cli <area> <resource>s update --id 42`.
6. Callbacks receive `ctx context.Context` and `client *netbox.APIClient` directly — no manual client retrieval needed.
7. Call the corresponding `client.<Area>API.<Area>XxxList(ctx).Limit(0).Execute()`.

## Code standards

- **Formatting**: `gofmt` (tabs, not spaces). `.editorconfig` has a `[*.go]` override.
- **Style**: See [style.md](style.md) for a full summary. Primary sources: [Effective Go](https://go.dev/doc/effective_go), [Google Go Best Practices](https://google.github.io/styleguide/go/best-practices), [Google Go Decisions](https://google.github.io/styleguide/go/decisions).
- **Linting**: `make lint` runs `golangci-lint` (config in `.golangci.yml`). All issues must be resolved before committing.
- **Comments**: godoc on all exported symbols and packages.
- **Module**: `github.com/kirtis/netbox-cli`, Go 1.22.
- **go-netbox**: `github.com/netbox-community/go-netbox/v4 v4.3.0` (published module).

## Documentation

- Write for human readability — avoid unnecessary jargon.
- Keep docs consolidated; do not create separate planning files when content can live in existing docs.
- When updating the README, ensure it accurately reflects current project capabilities.

## API coverage

All ten NetBox API areas are fully implemented with list, get, create, update,
and delete. Read-only resources (e.g. `core/data-files`, `extras/object-types`)
expose only list and get.

| Area | Status |
|---|---|
| circuits | complete |
| core | complete |
| dcim | complete |
| extras | complete |
| ipam | complete |
| tenancy | complete |
| users | complete |
| virtualization | complete |
| vpn | complete |
| wireless | complete |
