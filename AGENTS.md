# AGENTS.md

Guidance for AI coding assistants (Claude Code, Codex, etc.) working in this repository.

## This repo is Windows-only

Every `.go` file in this repository carries a `//go:build windows` (and `// +build windows`) constraint. Any new source file you add MUST carry the same constraint unless you have a specific reason not to.

Implications for building/testing on a non-Windows machine:

- `go build`, `go vet`, `gofmt`, `golangci-lint` all work fine with `GOOS=windows` set — these are compile-time/static-analysis only and don't execute the compiled binary.
- `go test` does NOT work cross-platform here: it compiles the test binary for `GOOS=windows` and then tries to *run* it, which fails with `exec format error` on macOS/Linux. Only run `go test` on an actual Windows machine, or rely on CI's `windows-latest` job (`.github/workflows/ci.yml`) for real test execution.

## Project layout

- `cmd/windows-os-info/main.go` — CLI entrypoint
- `cmd/windows-os-info/util.go` — OS arch / version / CPU count detection logic
- `cmd/windows-os-info/util_test.go` — tests (run on `windows-latest` in CI)
- `internal/version/version.go` — single `Version` constant, bumped manually before each release

## Build, test, lint

```sh
GOOS=windows go build ./...
GOOS=windows go vet ./...
GOOS=windows golangci-lint run ./...
GOOS=windows gofmt -l .   # must produce no output
```

Run `go test ./...` on Windows (or trust the `windows-latest` CI job) before considering any change complete.

## Release scope

This tool is only ever published as a single `windows/386` binary — do not add `windows/amd64` or `windows/arm64` targets casually; that's a deliberate scope decision, not an oversight. If broader platform support is genuinely needed, raise it as its own change rather than folding it into an unrelated commit.

## Commit messages

Write commit messages in English. Keep them short and describe the actual change — avoid placeholder messages like `init` or `update`.

## Release process

Releases are tag-triggered, not push-triggered:

1. Draft `release_notes.md` locally by reading the diff since the last tag (`git diff <last-tag>..HEAD`) — an AI assistant can draft this, but a human must review it before tagging.
2. Bump the `Version` constant in `internal/version/version.go` to match the new tag.
3. `git tag vX.Y.Z && git push origin vX.Y.Z` — this triggers `.github/workflows/release.yml`, which cross-compiles the `windows/386` binary and creates the GitHub Release using the committed `release_notes.md`.

Do not call any LLM API from within CI to generate release notes — that step happens locally, before tagging, to avoid paying per-run API costs in the pipeline.
