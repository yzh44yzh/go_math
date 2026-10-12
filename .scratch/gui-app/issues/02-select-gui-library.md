# Select the Go UI library

Type: grilling
Status: resolved
Blocked by: 01

## Question

Which library should implement the Linux GUI, considering the comparison in
Compare Go GUI libraries for Linux multiplication practice, desired visual control,
dependency burden, and long-term maintenance? Confirm the user's selection.

## Comments

The user asked whether Fyne is suitable for a later Android app. Android delivery
remains outside this Linux-first effort; library selection should consider that
future possibility.

Fyne documents shared desktop/mobile application code and Android APK packaging
with `fyne package -os android`, requiring the Android SDK and NDK.
See [mobile packaging](https://docs.fyne.io/started/mobile/).
Gio also supports Android APK builds; see [Android setup](https://gioui.org/doc/install/android).

Fyne remains the provisional recommendation for this simple practice app. Future
Android work should verify small-screen layout, touch controls, the soft keyboard,
Back behavior, and lifecycle handling on a device. These have not been tested.

## Answer

The user selected **Fyne** with standard widgets and layouts for the Linux-first
GUI. Its built-in form and practice-screen controls fit the application's small
scope. Android support is a future consideration, not a delivery requirement for
this effort; favor flexible layouts and visible controls when reviewing the UI.

Use `fyne.io/fyne/v2`. Select and pin an exact version during implementation,
verifying compatibility with the repository's Go 1.24 baseline and the selected
Linux environment. The research checked Fyne v2.8.1 module metadata; no Fyne build
or runtime validation has been performed in this repository.
