# Triage Labels

Use these canonical roles as local issue `Status:` values.

| Role | Status string | Meaning |
| --- | --- | --- |
| needs-triage | needs-triage | Maintainer evaluation needed |
| needs-info | needs-info | More information needed |
| ready-for-agent | ready-for-agent | Fully specified for agent implementation |
| ready-for-human | ready-for-human | Requires human implementation |
| wontfix | wontfix | Will not be actioned |

When a skill asks to apply a triage label, update the issue's `Status:`
line. Wayfinding tickets use `claimed` and `resolved` for their lifecycle,
as described in `issue-tracker.md`.
