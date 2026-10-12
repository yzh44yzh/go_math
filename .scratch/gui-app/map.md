# Plan the Linux multiplication GUI

Label: wayfinder:map
Status: resolved

## Destination

An agreed GUI design and implementation plan for replacing the multiplication CLI
with a Linux desktop application, including a selected Go UI library and validation strategy.

## Notes

- Planning only: implementation follows a separate handoff.
- User-confirmed scope: multiplication only; no accounts or saved history.
- Flow: settings, one question at a time with feedback, then results with
  Retry mistakes and New session.
- Editable minimum factor, maximum factor, and question count; defaults 1–9 and 20.
- Linux first. Replace the CLI rather than maintaining two interfaces.
- Consult wayfinder, grilling, and domain-modeling for decisions; research for
  library facts; prototype when screen behavior needs a concrete illustration.
- Read `CONTEXT.md` and relevant ADRs when present.
- Child decisions live in `issues/`; choose the lowest-numbered open, unblocked,
  unclaimed ticket. Claim before work. Resolve at most one non-research ticket per session.

## Decisions so far

- [Compare Go GUI libraries for Linux multiplication practice](issues/01-research-go-gui-libraries.md): cited research
  favors Fyne provisionally; Gio and Wails tradeoffs and tagged Go requirements
  are recorded; platform/build proof remains open.
- [Select the Go UI library](issues/02-select-gui-library.md): Fyne with standard
  widgets and layouts; Linux first, with Android considered for future portability.
- [Define question feedback and retry behavior](issues/03-define-practice-interactions.md):
  manual feedback advancement, validated settings, and randomized mistake-only retry rounds.
- [Review the GUI screen layout](issues/04-design-screen-layout.md): Focus layout
  selected; centered single-column screens and keyboard focus rules recorded.
- [Define Linux delivery and acceptance checks](issues/05-define-linux-delivery.md):
  local executable for Ubuntu 26.04/KDE through XWayland, with behavior and graphical checks.
- [Plan the practice model and implementation handoff](issues/06-plan-session-model-and-tests.md):
  one package with separate practice/UI files, deterministic behavior tests, and
  an ordered handoff in [the agreed specification](spec.md).

## Not yet specified

<!-- No additional uncharted decisions currently; remaining questions have tickets. -->

## Out of scope

- Application implementation during this planning effort.
- Windows and macOS delivery, mobile delivery, and browser delivery.
- Additional arithmetic operations, accounts, and persistent practice history.
- Continued CLI support.
