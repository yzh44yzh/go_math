# Plan the practice model and implementation handoff

Type: grilling
Status: resolved
Blocked by: 02, 03, 04, 05

## Question

What separation between practice state and GUI event handling will keep the application
simple and testable? Agree on source organization, deterministic test seams, behavior
tests, UI verification, and ordered implementation steps. Capture the agreed design
and acceptance criteria in `.scratch/gui-app/spec.md` for the implementation handoff.

## Answer

The user accepted the proposed organization, state ownership, tests, and sequence.

- Keep one root-level `main` package. `go_math.go` starts Fyne and creates the
  window; `practice.go` owns settings validation and practice behavior; `ui.go`
  owns the Focus layout, navigation, and keyboard focus.
- Keep practice logic independent of Fyne widgets. Explicit question, feedback,
  and results states prevent duplicate submissions and skipped questions.
- Supply seeded randomness in tests. Verify validation, ordered-pair uniqueness
  within each batch, exact question counts, answer handling, round totals, and
  mistake-only retries preserving repeated occurrences (2026-10-13 revision).
- Add focused Fyne UI interaction checks and perform the agreed Linux graphical
  smoke check. Fyne documents [testing helpers](https://docs.fyne.io/started/testing/).
- Implement in order: prove a pinned Fyne build, extract/test the practice model,
  build screens, connect navigation/focus, update commands/docs, then validate.

The implementation handoff is [the agreed GUI specification](../spec.md). Earlier
decision tickets remain the detailed sources for behavior, layout, and delivery.
Application implementation is a separate effort.
