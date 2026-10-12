# Domain Docs

## Read Before Exploring

Read root `CONTEXT.md` when it exists, then read ADRs in `docs/adr/`
that concern the area being explored.

If these files are absent, proceed silently. Create domain documentation
lazily through the domain-modeling skill when terminology or decisions
are resolved.

## Layout

This repository uses a single-context layout:

- `CONTEXT.md`: domain terminology and glossary.
- `docs/adr/NNNN-<decision-slug>.md`: architecture decisions.
- Root-level `.go` files: application source.

## Vocabulary and Decisions

Use glossary terms in issue titles, proposals, hypotheses, and tests.
If a needed concept is missing, reconsider the terminology or note the
gap for domain-modeling.

Explicitly identify any proposal that conflicts with an existing ADR,
including the ADR number and the reason to revisit its decision.
