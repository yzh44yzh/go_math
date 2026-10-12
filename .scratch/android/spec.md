# Android Target Specification

Status: implemented and verified on device (2026-10-12)

## Goal and Scope

Add Android as a second target of the existing Fyne application with minimal
changes: package the app, fix only what device testing proves broken.
Distribution is a personal sideloaded APK over `adb`; no Play Store, no signing
infrastructure beyond the debug key the fyne CLI uses. One build architecture:
`arm64-v8a`. No persistence — a killed app loses the round, matching the
existing spec.

## Decisions

Settled in a grilling session on 2026-10-12:

- **Scope**: Fyne mobile packaging, fix-if-broken. No mobile redesign,
  no native rewrite.
- **Distribution**: personal sideload; `adb install`.
- **Architecture**: `arm64-v8a` only.
- **Verification**: physical device (Samsung Galaxy A34, SM-A346E, Android 16)
  driven over `adb`; headless tests unchanged.
- **Breakpoint policy**: test-first for the hardcoded window resize, hardware
  back button, and rotation; fix only proven breakage.
- **Identity**: app ID `io.github.yzh44yzh.go_math`, name "Go Math",
  version 1.0, generated placeholder `Icon.png` (replaceable).
- **Toolchain**: existing SDK at `~/dev/android-sdk` (platform android-34,
  build-tools 29.0.3, NDK r22.1 in legacy `ndk-bundle`). fyne CLI v2.8.1
  installed via `go install fyne.io/fyne/v2/cmd/fyne@v2.8.1`, matching the
  pinned library. SDK wired through Makefile-local `ANDROID_HOME` only;
  no shell-rc changes.
- **Persistence**: none on Android either.

## Implementation

- `FyneApp.toml`: packaging metadata (icon, name, ID, version).
- `Icon.png`: generated placeholder, white × over blue rounded square.
- `Makefile`: `make android` packages `Go_Math.apk`; `make android-install`
  additionally installs over `adb`. `ANDROID_HOME`, `ADB`, `FYNE`, `APK` are
  overridable variables.
- `ui.go`: all focus grabs go through `practiceUI.focus`, which no-ops when
  `fyne.CurrentDevice().IsMobile()`. On mobile, focusing any widget makes the
  Fyne driver pop the soft keyboard; this app has no text fields, so focus
  exists only for desktop hardware-key entry.

## Verification on device

- Installed and launched; settings, practice, feedback, and confirmation
  dialog all render and respond to touch (keypad entry, sliders, buttons).
- Full flow exercised over `adb`: digit entry → Check answer → Correct
  feedback → progress advance; Back to settings → discard confirmation →
  settings.
- Rotation: layout relayouts correctly; the content column scrolls in
  landscape. No orientation lock added.
- Hardware back button: harmless no-op (it neither exits nor navigates).
- Soft keyboard: was popping on every screen before the fix (Fyne mobile
  shows it for any focused widget); gone after the focus fix on all screens.
- Desktop regression checks after the change: `make test`, `make vet`,
  `make build` all pass.

## Known limitations

- The packaged manifest carries fyne's defaults: `minSdkVersion 15`,
  `targetSdkVersion 29`. The CLI exposes no flag for either; the intended
  floor of 23 was not enforceable and is moot for the target device.
- Android shows a one-time launch warning: the app is debuggable (expected
  for sideloads) and the native libraries lack 16 KB page-size ELF alignment
  (NDK r22 toolchain). The app runs regardless; the dialog was dismissed with
  "don't show again". Revisit if a future device refuses to run it.
- The device required `settings put global verifier_verify_adb_installs 0`
  and `package_verifier_enable 0` to permit `adb install`; both were left
  disabled so `make android-install` keeps working. Restore with value `1`.
- A Bluetooth keyboard on Android gets no focus-based shortcuts (Tab/Enter
  traversal); the on-screen keypad covers all entry.
- The fyne CLI at v2.8.1 is deprecated upstream in favor of
  `fyne.io/tools/cmd/fyne`; keep the version matched to the pinned library.
