# Introduction
API for initiating, maintaining and authenticating [ScanBadge](https://scanbadge.xyz/discover).

This API is written in Go.

## Setup

1. `go get github.com/scanbadge/api` (and its dependencies)
- `go install github.com/scanbadge/api`
- Edit the `config.example.json` file as following:
  - Add the missing MySQL details (`Username`,`Password`,`DatabaseName`)
  - Edit applicable values to match your current MySQL configuration
  - Rename `config.example.json` to `config.json`
  - Move `config.json` to the directory of your current ScanBadge application (e.g. `$GOPATH/bin/`).
- Use `cd $GOPATH/bin && ./api` to run ScanBadge API.
- If applicable, Gorp will automatically create empty tables in the selected database.
- ???
- Profit! :)

## Links
- [Project website](https://scanbadge.xyz/)
- [Full documentation](https://scanbadge.xyz/documentation)
- [Coding style rules](https://golang.org/doc/effective_go.html#formatting)
