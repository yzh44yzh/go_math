# Repository Guidelines

## Project Structure & Module Organization

This repository contains a small Go CLI for practicing multiplication. `go_math.go` holds the `main` package, flag parsing, question generation, terminal interaction, and retry loop. `go.mod` declares module `go_math` and Go 1.24.0. `Makefile` provides formatting, vetting, and build targets; `README.md` documents flags and an example session. There are currently no separate test or asset directories and no external dependencies.

Keep related changes in the existing package. If the source grows, split focused responsibilities into additional root-level `.go` files.

## Build, Test, and Development Commands

Use a Go toolchain compatible with `go.mod`.

- `make build`: runs formatting and vetting, then builds the `go_math` executable.
- `make fmt`: formats all Go packages with `go fmt ./...`.
- `make vet`: runs static checks with `go vet ./...`.
- `go test ./...`: runs all tests; currently there are no test files.
- `go run . --from=5 --to=8 --limit=8`: starts an interactive practice session without a separate build step.
- `./go_math --help`: lists flags after building.

## Coding Style & Naming Conventions

Use Go's standard formatting and tab indentation; run `make fmt` before submitting changes. Use mixedCaps for unexported functions and local variables, such as `makePairs` and `incorrectAnswers`, and MixedCaps for exported identifiers, such as `Pair`. Use descriptive names and keep functions focused. Preserve the long and shorthand CLI flags together when changing options.

## Testing Guidelines

Add tests using Go's standard `testing` package in root-level `*_test.go` files, with functions named `TestXxx`. Prefer table-driven cases for question generation and input parsing. Keep tests deterministic despite shuffled question order. No coverage threshold is configured. For terminal changes, manually check correct and incorrect answers, whitespace handling, and the repeat prompt.

## Commit & Pull Request Guidelines

Git history uses short, lowercase descriptions such as `colored output` and `better input handling`; follow that concise style. Keep commits focused. Pull requests should explain the behavior change and list validation commands and results. Link relevant issues when available, include terminal transcripts for output changes, and update `README.md` when flags or usage change.

## Agent skills

### Issue tracker

Track issues and specs as local Markdown under `.scratch/<feature>/`.
See `docs/agents/issue-tracker.md`.

### Triage labels

Use the five default triage roles as issue status strings.
See `docs/agents/triage-labels.md`.

### Domain docs

Use a single-context layout: root `CONTEXT.md` and `docs/adr/`.
See `docs/agents/domain.md`.
