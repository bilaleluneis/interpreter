# rustimpl Workspace Instructions

This repository is a phased Rust workspace for learning interpreters and compilers with MonkeyLang as the target language.

## Workspace Shape

- Treat the repository as a Rust workspace first.
- The current core crates are `token`, `scanner`, `parser`, `ast`, and `eval`.
- The `interpeter` crate exists in the repository and should only be brought into workspace membership or changed intentionally.
- Keep crate responsibilities narrow and avoid moving logic across crate boundaries unless the change truly belongs there.

## Phase Workflow

- Work in explicit phases.
- Use phase-specific instructions, agents, and history notes instead of mixing future work into the current phase.
- Preserve phase history in `.github/phase-history/` so later phases can build on earlier decisions without rewriting them.
- Phase one is the minimal MonkeyLang interpreter path documented in `.github/phase-history/phase-one.md`.

## Rust Decision Rules

- Prefer Rust 1.99.0-era capabilities and idioms before falling back to older approaches.
- Prefer immutable data flow, shared references, slices, iterators, and pattern matching over unnecessary cloning or owned copies.
- Make ownership boundaries explicit in public APIs.
- Use `Result` and `Option` for fallible and missing-data cases instead of sentinel values.
- When borrow checking gets difficult, first ask whether the API boundary can be tightened before introducing extra copies or shared ownership.

## Design and Learning Rules

- Follow TDD: write or update the smallest failing test before changing behavior whenever practical.
- Follow SOLID where it improves clarity, but do not force abstraction when a small Rust type or function is clearer.
- Keep scanning, parsing, AST, token, and evaluation concerns separated.
- When a change spans multiple crates, update downstream call sites and tests in the same change so the impact stays local and understandable.

## Analysis and Explanations

- When evaluating an approach, include the relevant computer science or math background when it helps explain the tradeoffs.
- Explain the pros and cons of competing approaches before committing to one when the decision is not obvious.
- Favor explanations grounded in how Rust ownership and borrowing affect the design.

## Agent Boundaries

- Custom agents should be crate-boundary aware.
- An agent may inspect other crates for context, but it should only create or modify files inside its assigned crate.
- If a change requires edits outside the agent's crate, describe the downstream impact instead of expanding the edit scope.

## Validation

- Keep validation narrow and relevant to the touched crates.
- Prefer the smallest useful `cargo fmt`, `cargo check`, or `cargo test` command.
- Update or add tests with behavior changes, especially when a new phase feature is introduced.
