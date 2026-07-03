## Changelog for v0.2.0

This release is a repository relaunch — no changes to the CLI's flags, output, or behavior. It focuses on making the project properly usable and maintainable by others:

- **Licensing**: added an MIT `LICENSE` (previously the repo had none).
- **CI/CD migrated from Azure DevOps to GitHub Actions**: this tool is Windows-only (every `.go` file carries `//go:build windows`), so CI uses a mixed runner strategy — `go build`/`go vet`/`golangci-lint` run on `ubuntu-latest` with `GOOS=windows` (compile-time checks only), while `go test` runs on `windows-latest` since executing a cross-compiled Windows test binary isn't possible on Linux. Releases are now triggered by pushing a `vX.Y.Z` tag instead of every push to `main`; the published binary is still `windows/386` only, no new platform targets were added.
- **Project layout**: moved to the standard `cmd/windows-os-info/` + `internal/version/` Go layout.
- **Code quality**: fixed a `golangci-lint` finding (simplified an if/else chain into a tagged switch) and pinned the lint ruleset in `.golangci.yml`.
- **Docs**: expanded `README.md` with install/usage/dev instructions, added `AGENTS.md` for AI coding assistants (including an explicit note about the Windows-only build constraint).
- Translated the remaining Chinese comments in `util.go` to English.
