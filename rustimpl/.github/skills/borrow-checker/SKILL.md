---
name: borrow-checker
description: "Use when a Rust edit triggers borrow checker errors, lifetime mismatches, or cloning that should be replaced with references."
argument-hint: "crate, error, or signature to inspect"
---
# Rust Borrow Checker Skill

## When to Use
Use this skill when code changes fail because of lifetimes, ownership, mutable aliasing, or avoidable cloning.

## Procedure
1. Locate the exact signature or expression that owns the data.
2. Determine whether the data can be borrowed instead of copied.
3. Tighten the lifetime or reference boundary only as far as the current use requires.
4. Update the smallest set of downstream callers and tests.
5. Validate with the narrowest useful `cargo check` or `cargo test` command.

## Rust Checks
- Prefer `&str`, `&T`, `&[T]`, and iterators where the code only reads.
- Use cloning only when the value must outlive the source or be independently mutated.
- Confirm the resulting API still expresses who owns the data.
