# Compare Go GUI libraries for Linux multiplication practice

Type: research
Status: resolved

## Question

What do current official sources establish about Fyne, Gio, and relevant alternatives
for this application's settings, question, feedback, and results screens?

Compare widgets/layout, keyboard interaction, Linux build/runtime dependencies,
packaging, licensing, and compatibility with the project's Go 1.24 baseline.
Distinguish verified facts from recommendations and untested platform behavior.
Capture cited findings as an asset on a throwaway research branch and link them here.

## Answer

See [the cited research artifact](../research/gui-libraries.md). Preserved on local
branch `codex/research/gui-libraries`, commit `79763d1`, artifact path
`.scratch/gui-app/research/gui-libraries.md`; separate worktree:
[`/tmp/go-math-gui-library-research`](/tmp/go-math-gui-library-research/.scratch/gui-app/research/gui-libraries.md).

Fyne supplies the required form/practice widgets and Linux packaging tooling;
Gio offers immediate-mode layout and event management; Wails adds a web frontend
and GTK/WebKit runtime requirements. Fyne is the provisional recommendation,
pending the user's decision in [Select the Go UI library](02-select-gui-library.md).

Tagged module minimums: Fyne v2.8.1 declares Go 1.22, Gio v0.10.2 declares Go 1.24,
and Wails v2.15.0 declares Go 1.25. Earlier Wails v2.10.2 declares Go 1.22.
Gio v0.10.3 metadata was not retrievable. No complete dependency/build verification
or graphical/platform testing was performed. Accessibility APIs are present but
Linux screen-reader integration remains untested.
