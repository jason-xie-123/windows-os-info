## Changelog for v0.2.2

Housekeeping release, no CLI behavior changes.

- Added `.gitignore` for `.DS_Store` and removed the one already committed under `scripts/`.
- Corrected `AGENTS.md`: `internal/version/version.go` has no `//go:build windows` constraint (it's a platform-independent constant), so the previous "every `.go` file" claim was inaccurate.
- Aligned CI's `static-checks` job to `GOARCH=386`, matching the only target `release.yml` actually publishes (it previously defaulted to `amd64`, which build/vet/lint never exercised).
