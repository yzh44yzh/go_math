# Define Linux delivery and acceptance checks

Type: grilling
Status: resolved
Blocked by: 02

## Question

Which Linux distributions and display environments should the first version support,
and how should it be installed and launched? Agree on packaging, native dependency
expectations, and smoke checks on the intended environment. Decide whether screen-reader
support is an acceptance requirement and how to verify it if so.

## Answer

The user accepted the following first-release delivery and acceptance requirements.

### Supported environment

- Ubuntu 26.04 on x86-64, KDE with a Wayland session.
- Use Fyne's default X11 build through XWayland, which is available on the current
  machine. Native Wayland delivery is deferred.
- Other Linux distributions and desktop environments are not release acceptance
  targets for this first version.

### Delivery

- Build a local executable; document development/runtime prerequisites and launch
  commands in `README.md`.
- Verify the pinned Fyne version builds with the repository's Go 1.24 baseline.
- Identify actual native dependencies during implementation and verify the binary
  starts on the supported environment; do not assume a static standalone binary.
- An installer or distributable package is deferred until needed.

Fyne documents a separate `wayland` build tag in
[compile options](https://docs.fyne.io/explore/compiling/) and Linux `.tar.gz`
distribution in [packaging](https://docs.fyne.io/started/packaging/).

### Acceptance checks

- Invalid settings are rejected; valid settings start practice with the correct
  factor range, requested count, and any explained cap on available pairs.
- Correct and incorrect answers produce the approved feedback. Blank/non-integer
  input keeps the question active without recording a mistake.
- Feedback advances only on Next or Enter, including the transition to results.
- Results reflect the completed round. Retry mistakes includes only the preceding
  round's mistakes and shuffles them anew for each retry round.
- Enter and Tab support the full practice flow, with appropriate focus after each
  transition and visible controls for mouse/touch-style interaction.
- Back to settings confirms discarding practice progress; cancelling preserves
  progress. New session retains settings. Closing the window exits.
- Resizing preserves readable text and reachable controls without clipping.
- Build and automated behavior checks pass; perform a graphical smoke check on
  the supported Linux environment.

Readable text and keyboard operation are required. Screen-reader certification is
deferred; no working Linux screen-reader integration is claimed.

These are agreed requirements, not completed validation: no Fyne application has
yet been built or launched in this repository.
