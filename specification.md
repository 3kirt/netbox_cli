# Introduction

I would like to work on a command line interface for netbox.  This tool will primarily be used to get inventory information as context for Claude.

# Implementation

- The command line tool must be written in go
- The command line tool must use the go-netbox library.  I have checked out the source code for go-netbox in /Users/kirtis/source/repos/go-netbox

- The command line tool will use a netbox token for API access.  It should be able to read the token from an environment variable.
- The command line tool should be able to read it's configuration from a JSON configuration file.
- To limit the scope of the initial implementation, the tool should start by providing read access to the virtualization and ipam APIs.
- The tool will be expanded over time to support all of the netbox API, so you should organize the code to support expansion.

# API Coverage

| Area | Status |
|---|---|
| circuits | not started |
| dcim | complete |
| extras | complete |
| ipam | complete |
| tenancy | not started |
| users | not started |
| virtualization | complete |
| vpn | not started |
| wireless | not started |

# Standards

## Text encoding

- The project should follow the text encoding standards in .editorconfig

## Style

- The project should follow the go style guide: https://google.github.io/styleguide/go/guide
