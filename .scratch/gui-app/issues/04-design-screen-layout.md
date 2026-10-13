# Review the GUI screen layout

Type: prototype
Status: resolved
Blocked by: 02, 03

## Question

What concrete settings, practice, feedback, and results layout should be built?
Use a throwaway prototype to agree on keyboard focus, readable question text,
feedback presentation, button placement, and window resizing behavior.

## Prototype asset

[Throwaway interactive layout sketches](/tmp/go-math-gui-layout-prototype/.scratch/gui-app/prototypes/gui-layouts.prototype.html)
compare **Focus** (centered single task), **Sidebar** (persistent session context),
and **Worksheet** (wide settings and split question/answer areas).

These are browser sketches for discussing layout, not exact Fyne widget renderings.
Preview settings, questions, correct/incorrect feedback, and results; local practice
controls model the accepted submission, validation, and retry rules.

Proposed focus behavior: answer entry on each question, Next after feedback, and
the primary action on results. Enter submits an answer or activates the focused
Next action. Layouts reflow to a single column on narrow screens.

Preserved on branch `codex/prototype/gui-layouts`, commit `3f929c3`, at
`.scratch/gui-app/prototypes/gui-layouts.prototype.html`. The branch contains all
three alternatives as the primary source for this review. The main checkout keeps
the decision, not prototype code. Application implementation remains outside this
planning effort.

## Answer

The user selected **Focus** without additional changes.

- Use one resizable window, a small Go Math header, and a centered content column.
- Settings show vertically stacked Minimum factor, Maximum factor, and Questions
  inputs, followed by range guidance and Start practice. The 2026-10-13 revision
  removes the question-count cap notice.
- Practice shows the current question count and progress, a large centered
  multiplication question, then answer entry and Check answer.
- Focus the answer field on each question. Enter checks the answer.
- Feedback replaces answer entry with the submitted answer and a text result:
  Correct, or the correct multiplication result. Focus Next (See results after
  the final question); Enter activates it. Feedback does not advance automatically.
- Results show this round's correct and incorrect counts. Retry mistakes is the
  primary action when mistakes remain; otherwise New session is primary.
- Back to settings remains visible during practice, with the agreed discard
  confirmation. Narrow windows retain the single-column layout and allow controls
  to wrap without clipping.
- Implement this hierarchy using standard Fyne widgets and layouts. The sketch's
  browser styling is illustrative; native Fyne appearance and focus/resizing
  behavior require verification during implementation.
