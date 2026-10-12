# Repository Guidelines

## Project Structure & Module Organization

Go Math is a Fyne multiplication practice app with one root-level `main` package.
`go_math.go` starts the window; `practice.go` owns validation, question selection,
answers, and retries; `ui.go` owns the Focus layout, navigation, and keyboard focus.
`practice_test.go` and `ui_test.go` exercise behavior and UI actions. `go.mod` pins
Fyne; `go.sum` records dependency checksums. `README.md` documents setup and use.

Read `.scratch/gui-app/spec.md` and its linked decisions when changing GUI behavior.
The reviewed browser prototype is preserved on `codex/prototype/gui-layouts` and
is a layout reference, not production source.

## Build, Test, and Development Commands

Use Go 1.24 or newer. Native Ubuntu builds need CGO, a C compiler, and OpenGL/X11
headers; see README setup. The first target is Ubuntu/KDE through XWayland.

- `make build`: format, vet, and build `go_math` with the `x11` tag.
- `make run`: launch the native GUI.
- `make fmt`: format all packages.
- `make vet`: vet the native X11 build.
- `make test`: run headless model and UI tests with Fyne's `ci` tag.
- `go test -tags x11 ./...`: test with native build dependencies available.

## Coding Style & Naming Conventions

Use standard Go formatting with tabs, mixedCaps for unexported identifiers, and
MixedCaps for exported identifiers. Keep practice logic independent of Fyne
widgets. Explicit question, feedback, and results states enforce legal actions;
UI callbacks render results and manage focus. Keep this small app on the UI event
loop rather than adding background workers.

## Testing Guidelines

Use Go's `testing` package in `*_test.go` files with `TestXxx` functions. Use table-driven
validation cases and seeded randomness. Verify pair/permutation invariants rather
than expecting every shuffle to change order. Fyne test helpers cover meaningful
UI actions; keep native keyboard, resizing, and launch checks separate. There is
no coverage threshold. Run `make test`, `make vet`, and `make build` after changes.

## Commit & Pull Request Guidelines

Use short, lowercase commit descriptions, following history such as `colored output`.
Keep commits focused. Describe behavior changes and verification results in PRs;
link relevant issues and include screenshots for UI changes. Update README and
this guide when commands, dependencies, or source organization change.

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
