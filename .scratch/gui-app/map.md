# Plan the Linux multiplication GUI

Label: wayfinder:map

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
  are recorded, while library selection and platform/build proof remain open.

## Not yet specified

- Screen layout and visual style need refinement after library selection.
- Implementation sequencing depends on the session model and packaging decisions.

## Out of scope

- Application implementation during this planning effort.
- Windows and macOS delivery, mobile delivery, and browser delivery.
- Additional arithmetic operations, accounts, and persistent practice history.
- Continued CLI support.
