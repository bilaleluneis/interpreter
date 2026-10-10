---
name: rustimpl-workspace
description: "Use when working on the rustimpl interpreter workspace, especially for parser, scanner, AST, token, Cargo workspace, ownership, borrowing, and borrow-checker-friendly refactors."
argument-hint: "crate name, bug, or refactor scope"
---
# Rustimpl Workspace Skill

## When to Use
Use this skill when you need to inspect, change, or validate the Rust interpreter workspace. It is especially useful for:
- Borrow checker errors
- Ownership or lifetime refactors
- Parser, scanner, AST, or token changes
- Workspace or Cargo manifest edits
- Reducing unnecessary cloning by using shared references

## Procedure
1. Identify the owning crate and the smallest boundary where data changes hands.
2. Prefer shared references, slices, and iterators over new owned copies.
3. Keep AST, token, scanner, and parser responsibilities in their own crates.
4. Update crate-local tests and any direct downstream call sites together.
5. Validate with the narrowest useful `cargo fmt`, `cargo check`, or `cargo test` command.

## Rust-Specific Checks
- Confirm the API expresses ownership clearly.
- Avoid moving data if a read-only reference is enough.
- Verify any borrowed data outlives the use site.
- Re-run tests that cover the exact crate boundary you changed.

## Notes
The current workspace members are `ast`, `parser`, `scanner`, and `token`. The `interpeter` crate exists in the repository and should be added to the workspace only when that is an intentional change.