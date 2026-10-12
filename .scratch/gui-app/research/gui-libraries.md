# Go GUI library research

Research date: 2026-10-12. Scope: Linux desktop replacement for the multiplication
CLI; settings, one question with feedback, results, and retry. This records facts
and a provisional recommendation, not the library-selection decision.

## Verified facts

### Fyne

- Provides labels, buttons, entries, forms, progress indicators, and box/grid/border
  layouts. Entries support validation and an `OnSubmitted` callback for Enter in
  single-line fields. [Widget documentation](https://docs.fyne.io/),
  [Entry API](https://docs.fyne.io/api/v2/widget/entry/).
- Linux development needs Go, a C compiler, and graphics headers/drivers. The
  Debian/Ubuntu guide lists GCC, Mesa/OpenGL, Xorg, Wayland, and xkbcommon development
  packages. Fyne uses native graphics integration through CGo, so cross-compilation
  also needs a target C toolchain. These are development prerequisites; verify the
  resulting binary on the target desktop rather than assuming a static executable.
  [Quick start](https://docs.fyne.io/started/quick/),
  [cross-compilation](https://docs.fyne.io/started/cross-compiling/).
- `fyne package -os linux -icon <icon.png>` creates a `.tar.gz` containing a Linux
  installation layout. A Wayland build tag is documented separately from the
  default desktop build. [Packaging](https://docs.fyne.io/started/packaging/),
  [compile options](https://docs.fyne.io/explore/compiling/).
- License: BSD-3-Clause. [License](https://github.com/fyne-io/fyne/blob/master/LICENSE).

### Gio

- Uses immediate-mode UI: receive events, update application state, then lay out
  and draw. Provides editors/clickable widgets, Material styling, and Flex, Stack,
  and List layouts. Keyboard and focus are explicit parts of its input system.
  [Architecture](https://gioui.org/doc/architecture),
  [widgets](https://pkg.go.dev/gioui.org/widget),
  [Material widgets](https://pkg.go.dev/gioui.org/widget/material),
  [layouts](https://gioui.org/doc/architecture/layout),
  [input](https://gioui.org/doc/architecture/input).
- Linux development requires Wayland, X11/xkbcommon, GLES/EGL, and Xcursor packages,
  plus GCC and pkg-config in the documented distro commands. Supports `nox11` and
  `nowayland` build tags. Vulkan availability may affect performance.
  [Linux installation](https://gioui.org/doc/install/linux).
- Desktop applications can be built with the Go tool directly; `gogio` supports
  mobile and additional desktop features such as icons. Linux installation/package
  conventions still need to be chosen for this repository.
  [Build tooling](https://gioui.org/doc/install).
- License: MIT or Unlicense, at the user's choice.
  [License](https://github.com/gioui/gio/blob/main/LICENSE).

### Wails v2: relevant web UI alternative

- Combines Go with a web frontend; templates include vanilla JavaScript, React,
  Vue, and Svelte. Development includes frontend tooling such as npm.
  [Project structure](https://wails.io/docs/gettingstarted/firstproject/),
  [installation](https://v2.wails.io/docs/gettingstarted/installation/).
- Linux builds need GCC, GTK3, and WebKitGTK. End users need GTK3/WebKitGTK runtime
  libraries. ABI 4.0/4.1 and package names vary by distribution, so packaging must
  declare the matching dependencies. `wails build` places the binary in `build/bin`.
  [Linux dependencies/packaging](https://wails.io/docs/guides/linux-distro-support/),
  [build command](https://wails.io/docs/gettingstarted/building/).
- Official project FAQ identifies v2 as stable and v3 as beta. License: MIT.
  [Status](https://v3.wails.io/faq/),
  [license](https://github.com/wailsapp/wails/blob/master/LICENSE).

## Go 1.24 compatibility

Checked tagged module metadata, not successful compilation or the complete
transitive dependency graph:

| Version checked | Module Go directive | Consequence |
| --- | --- | --- |
| Fyne v2.8.1 | `go 1.22.0` | Its declared minimum fits Go 1.24. |
| Gio v0.10.2 | `go 1.24.0` | Its declared minimum matches this repository. |
| Wails v2.15.0 | `go 1.25.0` | Requires a newer toolchain than the repository baseline. |
| Wails v2.10.2 | `go 1.22.0` | Older candidate with a declared minimum that fits Go 1.24. |

Sources: [Fyne v2.8.1 go.mod](https://raw.githubusercontent.com/fyne-io/fyne/v2.8.1/go.mod),
[Gio v0.10.2 go.mod](https://raw.githubusercontent.com/gioui/gio/v0.10.2/go.mod),
[Wails v2.15.0 go.mod](https://raw.githubusercontent.com/wailsapp/wails/v2.15.0/v2/go.mod),
[Wails v2.10.2 go.mod](https://raw.githubusercontent.com/wailsapp/wails/v2.10.2/v2/go.mod).
Gio's September newsletter reports v0.10.3, but its tagged module could not be
retrieved through the research tool; its exact baseline remains unverified.
[September newsletter](https://gioui.org/news/2026-09).
Pin the chosen version and prove a build with Go 1.24 before treating compatibility
as confirmed. Tagged module metadata takes precedence over generic installation
prose, particularly for Wails.

## Accessibility and validation limits

Keyboard submission is documented for Fyne; Gio has keyboard/editor facilities.
Current Fyne exposes accessibility labels/roles since 2.8, and Gio exposes semantic
operations intended for external tools such as screen readers. These interfaces do
not by themselves establish functioning Linux screen-reader integration.
[Fyne Accessible](https://docs.fyne.io/api/v2/fyne/accessible/),
[Gio semantics](https://pkg.go.dev/gioui.org/io/semantic).
No GUI was built or launched. Keyboard focus, resizing, X11/Wayland behavior,
runtime distribution, and Orca support remain platform checks if required.

## Recommendation, separate from facts

Offer **Fyne** as the default candidate: the app's modest forms and practice flow
map directly to its ready-made widgets and layouts, and its packaging tool gives a
concrete Linux starting point. This suitability judgment is an inference.

Offer **Gio** if learning immediate-mode UI or substantial custom drawing is a
goal. Offer **Wails** if HTML/CSS/frontend development is preferred enough to
justify another toolchain and Linux webview dependencies. Do not select a library
until the user weighs those priorities.
