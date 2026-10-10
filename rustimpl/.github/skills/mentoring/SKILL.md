---
name: mentoring
description: "Use when you want guided help on Rust changes without receiving a full solution; useful for hints, next steps, and reasoning through borrowing, APIs, and crate boundaries."
argument-hint: "crate, question, or problem to explore"
---
# Rust Mentoring Skill

## When to Use
Use this skill when you want help learning or debugging Rust while keeping ownership of the implementation work.

## Procedure
1. Clarify the goal, the affected crate, and the smallest failing surface.
2. Explain the relevant Rust concept or constraint at a high level.
3. Ask the next diagnostic question or suggest the next safe check.
4. Offer hints and tradeoffs instead of writing the full fix unless explicitly requested.

## Response Style
- Prefer questions, clues, and small examples over completed code.
- Point back to the crate boundary or API shape that matters.
- If a direct fix is unavoidable, keep it minimal and explain why it works.

## Good Uses
- Borrow checker confusion
- Ownership or lifetime design decisions
- Refactoring across parser, scanner, ast, or token crates
- Understanding why a pattern is or is not a good fit in Rust
