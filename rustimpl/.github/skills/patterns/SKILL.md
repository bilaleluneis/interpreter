---
name: patterns
description: "Use when discussing or applying Rust design and testing patterns such as TDD, immutability, SOLID, and small-crate architecture in this workspace."
argument-hint: "pattern, crate, or design question"
---
# Rust Patterns Skill

## When to Use
Use this skill when you want guidance on Rust coding patterns, especially for interpreter-style code and crate boundaries.

## Procedure
1. Identify which pattern is relevant to the current problem.
2. Map the pattern to this workspace's crates and ownership model.
3. Prefer the smallest change that improves clarity or testability.
4. Validate the pattern with an example from the actual code path.

## Focus Areas
- TDD: write or update the smallest failing test before changing behavior.
- Immutability: prefer shared references and read-only views over copies when data is not mutated.
- SOLID: keep responsibilities narrow and avoid mixing scanning, parsing, AST, and token concerns.
- Rust fit: choose patterns that work with ownership and borrowing, not against them.

## Guidance Rules
- Do not force a pattern where Rust already gives a clearer primitive.
- Avoid unnecessary abstraction layers when a small function or enum is enough.
- Keep examples grounded in the existing workspace crates.
