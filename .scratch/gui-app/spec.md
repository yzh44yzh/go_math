# Linux Multiplication GUI Specification

Status: approved

## Feedback Placement Revision

Requested on 2026-10-15: place Correct and Correct answer feedback directly below
Your answer, above the keypad. Reserve that message space during answer entry
as well, keeping the keypad and primary action stationary between states.

## Settings Limits Revision

Requested on 2026-10-15: both factor sliders allow 1–12 and Questions allows
1–50. Defaults remain factors 1–9 and 20 questions. Answer entry accepts three
digits so products through 12 × 12 = 144 can be entered. This revision supersedes
the earlier factor/count bounds and two-digit answer limit in linked decisions.

## Exact Question Count Revision

Approved on 2026-10-13: the initial round must contain exactly the configured
count (1–50). Shuffle all available ordered pairs, use each once, and start
another shuffled batch as needed. Repeats are allowed, including at batch
boundaries. Remove the count-cap notice from settings. Retry rounds contain
every missed occurrence from the immediately preceding round; their count is
the mistake count rather than the configured initial count.

This revision supersedes earlier selection-without-replacement and count capping.

## Readability and Stable Actions Revision

Approved on 2026-10-13 for all screens, confirmation dialogs, desktop, and Android:

- Body and button fonts: 20 logical units; question: 36; buttons: minimum height 56.
- Preserve these sizes on small screens and allow vertical scrolling.
- Keep the keypad visible but disabled during feedback. Reserve a message slot
  for empty-answer validation and correct/incorrect feedback, so Check answer,
  Next, and See results have the same position and size.
- Use orange progress fill distinct from the primary button color, supporting
  light and dark themes.
- Use exact wording New Session for header navigation, discard confirmation,
  and results. Keep discard confirmation and preserve selected settings.

This revision supersedes earlier feedback layout and navigation wording.

## Answer Control Revision

On 2026-10-12 the user approved replacing the answer text input with digit buttons
and a read-only answer display. Arrange rows as `1 2 3`, `4 5 6`, `7 8 9`, then
`Clear 0 ⌫`. Digit presses append without submitting, up to three digits; an initial
zero is replaced by the next digit. Clear removes everything; ⌫ removes the last
digit. Empty submission keeps the current question active without a mistake.

Hardware digit keys, Backspace, and Delete perform the same entry/editing actions.
Enter submits while a keypad or Check answer control is focused. Tab traverses
controls; Space activates a focused digit or editing button. Focus Check answer
on each new question to support direct hardware digit entry. Answer entry is
inactive during feedback, results, settings, and navigation confirmation.
Back to settings retains its navigation action when explicitly activated.
This revision supersedes references to answer text inputs and their focus rules
in earlier decisions.

## Settings Control Revision

The user approved replacing the three settings inputs with integer sliders on
2026-10-12. Each slider displays its current value in its label. Factors remain
1–12 and question count 1–50, with existing defaults.
Raising minimum past maximum raises maximum; lowering maximum past minimum
lowers minimum. Arrow keys adjust by 1, Tab traverses the three sliders and Start
practice, and Enter starts only when Start practice is focused. Model validation
remains; sliders prevent malformed and out-of-range settings entry in the GUI.
This revision supersedes references to settings text inputs in earlier decisions.

## Goal and Scope

Replace the multiplication CLI with a Fyne desktop application using the approved
Focus layout. First delivery targets Ubuntu 26.04, x86-64, KDE on Wayland through
XWayland. Deliver a locally built executable with documented prerequisites.

This is an implementation handoff, not an implemented or tested application.
Android is a future portability consideration. Accounts, saved history, other
arithmetic operations, installers, native Wayland delivery, other platform releases,
and continued CLI support are outside this first version.

## Approved Decisions

- [Library research](issues/01-research-go-gui-libraries.md)
- [Fyne selection](issues/02-select-gui-library.md)
- [Practice interactions](issues/03-define-practice-interactions.md)
- [Focus layout](issues/04-design-screen-layout.md)
- [Linux delivery and acceptance](issues/05-define-linux-delivery.md)
- [Practice model and tests](issues/06-plan-session-model-and-tests.md)

The linked decisions hold the detailed requirements. Read them alongside this
handoff before implementation. Domain vocabulary is in root `CONTEXT.md`.

## Application Flow

Settings → question → feedback → next question or round results.
From results, retry mistakes starts another round; New session returns to settings.

| Area | Required behavior |
| --- | --- |
| Settings | Integer sliders for minimum and maximum factors from 1–12, automatically keeping minimum ≤ maximum; question count 1–50, default 20; default factor range 1–9. |
| Selection | Honor the exact configured count using shuffled batches of all ordered pairs. Use every pair once per batch before repeating. Repeats at batch boundaries are allowed. `3 × 4` and `4 × 3` are separate questions. |
| Submission | Buttons/keyboard compose up to three digits in a read-only display. Enter or Check answer submits; an empty answer shows Enter a whole number without recording a mistake or advancing. |
| Feedback | Correct, or the correct multiplication result. Wait for Next/Enter; the final feedback advances to results. |
| Results | Correct and incorrect counts for the completed round; offer Retry mistakes only when mistakes remain. |
| Retry | Every missed occurrence from the immediately preceding round, including repeated questions, reshuffled for every round. A random shuffle may reproduce the previous order. |
| Navigation | Back to settings confirms discarding progress; cancellation preserves it. New session retains settings. Closing exits. |

## UI and Focus

Use standard Fyne widgets and layouts in one resizable window. Keep a small Go Math
header and a centered content column. Settings sliders stack vertically. Practice
shows question progress, a large centered multiplication question, then a read-only
answer display, reserved feedback space, and keypad. Feedback shows the submitted
answer and result above the disabled keypad. Results
display round totals and the available actions.

Focus Check answer on each question, Next/See results after feedback, and the
primary results action. Support Tab navigation and Enter activation. Keep controls
visible and layouts usable when resized; use text as well as any feedback color.

Prototype source: branch `codex/prototype/gui-layouts`, commit `3f929c3`, file
`.scratch/gui-app/prototypes/gui-layouts.prototype.html`. It is a layout reference;
implement Fyne screens afresh rather than promoting the browser prototype.

## Source Organization and State

Keep the root-level `main` package:

- `go_math.go`: application startup and window creation.
- `practice.go`: validation, ordered-pair generation/selection, answer checking,
  progress, round results, and retry state.
- `ui.go`: Fyne screen construction, event handling, navigation, and focus.
- `practice_test.go` and `ui_test.go`: behavior tests and focused UI checks.

The practice model owns state independently of widgets. Use explicit question,
feedback, and results states. Actions must not record a submitted answer twice or
skip an unanswered question. Keep settings values available when starting a new
session. Supply a seeded random source for tests; normal sessions use randomness.

## Implementation Sequence

1. Pin a Fyne v2 release compatible with the Go 1.24 baseline. Verify dependencies
   and a build on the supported environment before committing to that version.
   Research checked Fyne v2.8.1 module metadata, not a full dependency build.
2. Extract practice logic from terminal input/output and add deterministic tests
   for the approved behavior. Remove the terminal practice flow as the GUI replaces it.
3. Implement settings, question, feedback, and results using the Focus hierarchy.
4. Connect model actions, retry behavior, discard confirmation, and keyboard focus.
   Keep work on the UI event loop; this small app requires no background workers.
5. Update Makefile with usable build/test/run commands and update README and
   AGENTS.md for the final GUI structure, dependencies, and launch instructions.
6. Run automated checks, build the executable, and perform the graphical smoke check.

## Validation and Completion

Use Go's standard testing package. Table-driven tests should cover valid/invalid
settings, inclusive range boundaries, ordered-pair uniqueness within each batch, exact counts, input
handling, correct/incorrect answers, manual advancement, round totals, and repeated
mistake-only retries. Verify selection/permutation invariants rather than assuming
every random shuffle changes the order.

Use Fyne's testing helpers for a few meaningful checks across the UI/model
connection: starting practice, checking an answer, feedback advancement, and retry.
Document any headless test configuration actually needed by the selected release.
Headless checks do not replace graphical platform verification.

Run formatting, `go vet ./...`, `go test ./...`, and a native GUI build with the
chosen dependency/toolchain setup. Manually verify keyboard focus, Enter/Tab,
validation messages, exact question counts, correct/incorrect feedback, repeated retries, navigation
confirmation, retained settings, resizing, and closing on Ubuntu/KDE through
XWayland. Screen-reader certification is deferred.

Implementation is complete when the agreed behavior is delivered, these checks
pass, the executable launches in the supported environment, and documentation
matches actual commands and prerequisites. Report any unavailable graphical check
as a remaining validation gap rather than a completed check.
