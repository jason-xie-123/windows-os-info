# windows-os-info

[![CI](https://github.com/jason-xie-123/windows-os-info/actions/workflows/ci.yml/badge.svg)](https://github.com/jason-xie-123/windows-os-info/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/jason-xie-123/windows-os-info)](https://github.com/jason-xie-123/windows-os-info/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

A small Windows-only CLI tool that reports OS architecture, OS version, and CPU core count — useful in install scripts and CI where you need this info without pulling in a bigger dependency.

This tool only targets Windows (all source files carry a `//go:build windows` constraint) and is only published as a `windows/386` binary.

## Install

Download `windows-os-info.exe` from the [latest release](https://github.com/jason-xie-123/windows-os-info/releases/latest).

Or build from source (requires a Windows host, or cross-compile with `GOOS=windows`):

```sh
GOOS=windows go install github.com/jason-xie-123/windows-os-info/cmd/windows-os-info@latest
```

## How to Use

```
windows-os-info.exe -h
NAME:
   windows-os-info - CLI tool to echo windows os info scripts

USAGE:
   windows-os-info [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --action value  support action: os_arch / os_version / cpu_num
   --help, -h      show help
   --version, -v   print the version
```

Example:

```sh
windows-os-info.exe --action os_arch
windows-os-info.exe --action os_version
windows-os-info.exe --action cpu_num
```

## Development

This repo is Windows-only. On a non-Windows machine, compile-time checks (`go build`, `go vet`, `golangci-lint`) work fine with `GOOS=windows` set, but `go test` needs to actually execute the compiled binary and therefore requires a real Windows machine (or CI's `windows-latest` runner):

```sh
GOOS=windows go build ./...
GOOS=windows go vet ./...
GOOS=windows golangci-lint run ./...
GOOS=windows gofmt -l .

# on Windows only:
go test ./...
```

Releases are cut by pushing a `vX.Y.Z` tag — see `.github/workflows/release.yml`. Release notes live in `release_notes.md` and are drafted locally before tagging (see `AGENTS.md`).

## License

[MIT](./LICENSE)
