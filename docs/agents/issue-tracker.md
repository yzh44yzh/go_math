# Issue Tracker: Local Markdown

Issues and specs live as Markdown files in `.scratch/`.

## Conventions

- One feature per directory: `.scratch/<feature-slug>/`.
- Specs live at `.scratch/<feature-slug>/spec.md`.
- Store each implementation ticket separately at
  `.scratch/<feature-slug>/issues/<NN>-<slug>.md`, numbered from `01`.
- Record triage state in a `Status:` line near the top of each issue.
  Use the strings defined in `triage-labels.md`.
- Append conversation history under `## Comments`.

## Skill Operations

To publish a spec or ticket, create its file using these paths.
To fetch a ticket, read the referenced file. Resolve bare ticket numbers
within the relevant feature directory.

## Wayfinding Operations

- Map: `.scratch/<effort>/map.md`, with Notes, Decisions-so-far, and Fog.
- Child tickets: `.scratch/<effort>/issues/NN-<slug>.md`.
- Record `Type: research`, `prototype`, `grilling`, or `task`.
- Record dependencies near the top as `Blocked by: NN, NN`.
- Select the lowest-numbered open, unblocked, unclaimed ticket.
  Dependencies are satisfied when all referenced tickets are resolved.
- Claim a ticket by saving `Status: claimed` before starting work.
- Resolve it by appending `## Answer`, saving `Status: resolved`,
  and adding a gist and link to the map's Decisions-so-far section.
