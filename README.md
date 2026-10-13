# Go Math

A Fyne app for practicing multiplication. Supported targets: Ubuntu 26.04,
x86-64, KDE on Wayland with XWayland installed, and Android (`arm64-v8a`,
sideloaded APK).

## Build and run

Use Go 1.24 or newer. Fyne 2.8.1 is pinned in `go.mod`; native builds require CGO,
a C compiler, and OpenGL/X11 development headers. On Ubuntu:

```bash
sudo apt install build-essential pkg-config libgl1-mesa-dev xorg-dev
make build
./go_math
```

Or use `make run` to build and launch directly. Both commands use `-tags x11`,
matching the supported XWayland target and avoiding native Wayland build headers.
The executable needs an active graphical session and working OpenGL/X11 runtime
libraries; it is not a fully static or universally portable Linux binary.

## Practice

1. Use the sliders to choose minimum and maximum factors from **1–9** and **1–100** questions.
   Defaults are **1–9** and **20** questions.
   Labels show the selected values. Moving minimum above maximum raises maximum;
   lowering maximum below minimum lowers minimum, keeping the range valid.
2. Start practice with exactly the selected number of questions. Each batch
   shuffles all available ordered pairs and uses each once before starting the
   next batch. Repeats are allowed, including at batch boundaries; `3 × 4` and
   `4 × 3` are separate questions.
3. Build the answer using digit buttons **0–9**, then press **Enter** or
   **Check answer**. The display accepts up to two digits; entering another digit
   after an initial zero replaces that zero. **Clear** removes the answer and
   **⌫** removes its last digit. An empty answer leaves the question active.
4. Read the feedback and press **Enter** or **Next** to continue. The keypad stays
   visible but disabled; **Check answer**, **Next**, and **See results** share one position.
5. Review round totals. **Retry mistakes** reshuffles only mistakes from the
   preceding round, including every missed occurrence of repeated questions.
   Retry counts depend on mistakes, rather than the settings count.
   **New Session** returns to settings and retains their values.

All screens use larger text and buttons: body and button text is 20 logical units,
the question is 36, and buttons are at least 56 high. Progress uses orange separately
from the primary button color. Smaller windows and Android landscape allow scrolling.

Use **Tab** to move between controls and arrow keys to adjust sliders by 1.
On settings, **Enter** starts only when **Start practice** is focused.
During answer entry, hardware digit keys also work; **Backspace** removes the last
digit and **Delete** clears the answer. **Enter** submits from the keypad or Check
answer button; use **Space** to activate a focused digit or editing button.
**New Session** asks before discarding
an unfinished round; closing the window exits. Settings and results live only in
memory. The GUI replaces the previous terminal interface and command-line flags.

## Android

Android packaging uses the fyne CLI (matching the pinned library) and a local
Android SDK with NDK:

```bash
go install fyne.io/fyne/v2/cmd/fyne@v2.8.1
make android          # builds Go_Math.apk
make android-install  # also installs over adb
```

`ANDROID_HOME` defaults to `~/dev/android-sdk`; override it, `ADB`, `FYNE`, or
`APK` on the make command line. The device needs USB debugging enabled and, on
some devices, the package verifier relaxed for `adb install` — see
[`.scratch/android/spec.md`](.scratch/android/spec.md) for details and known
limitations. On Android the app is fullscreen, all entry is by touch, and the
soft keyboard is intentionally never shown.

## Development checks

```bash
make fmt        # Standard Go formatting
make vet        # Static checks for the native X11 build
make test       # Model and Fyne UI tests using the headless ci driver
make build      # Format, vet, and build the native executable
```

`go test -tags ci ./...` does not require a display or native graphics headers.
Use the `x11` tag for ordinary native Go commands, such as
`go test -tags x11 ./...` and `go vet -tags x11 ./...`.

Source is one `main` package: `go_math.go` starts the app, `practice.go` owns practice
state, and `ui.go` implements the centered Focus layout. Tests use seeded random
sources and cover state transitions, validation, pair selection, retries, and UI
interaction. No coverage threshold is configured.

For graphical verification, launch on the supported desktop and check Enter/Tab
focus, slider bounds, keypad/keyboard editing, empty answers, exact question counts, feedback, repeated retries,
discard confirmation, retained settings, resizing, and closing. Headless tests
cannot certify native focus, graphics, or screen-reader behavior.

The approved specification and decision history live in
[`.scratch/gui-app/spec.md`](.scratch/gui-app/spec.md); the Android target in
[`.scratch/android/spec.md`](.scratch/android/spec.md).
Saved history and distributable packages are future work.
