---
name: enforce-patterns
description: Enforce the workspace's existing Go patterns, architecture boundaries, and SOLID-aligned design decisions.
---

# Enforce Patterns

Use this skill when you need to keep changes consistent with the current `goimpl` workspace design.

## Workspace patterns to preserve

- Target Go 1.27.
- Keep changes local to the owning package.
- Respect the workspace split: `ast` defines syntax tree types, `token` defines lexical tokens, `lexer` produces tokens, `parser` builds ASTs, `eval` evaluates ASTs, and `common` holds shared helpers and generic utilities.
- Prefer the standard library and avoid new dependencies unless the user explicitly asks for one.
- Follow the existing style of small, explicit types with short exported comments.
- Keep tests table-driven and colocated with the package they exercise.
- Run `gofmt` on any edited Go file.

## Current codebase conventions

- AST nodes use marker methods like `statementNode()` and `expressionNode()` to control which types satisfy the interfaces.
- State-bearing types typically use pointer receivers, while simple value types remain values.
- Parsing code favors explicit AST error nodes and clear failure paths over panics.
- Shared generic helpers live in `common`, but generics are used only where they improve clarity or reduce duplication.
- Existing naming and structure should be preserved unless the user explicitly requests a refactor.

## SOLID guidance for this workspace

- Single Responsibility Principle: keep each package, type, and function focused on one job; do not mix lexing, parsing, and evaluation logic.
- Open/Closed Principle: extend behavior by adding new AST nodes, parser branches, or evaluator cases instead of rewriting stable shared logic.
- Liskov Substitution Principle: any type that satisfies an interface should behave predictably in existing call sites; avoid interface contracts that require special-case handling.
- Interface Segregation Principle: keep interfaces small and purpose-built; prefer a narrow parser, lexer, or AST interface over a large combined abstraction.
- Dependency Inversion Principle: depend on abstractions at package boundaries, but do not introduce abstractions that obscure a concrete, simple dependency.

## Decision rules

- Prefer composition and small helper functions over large monolithic types.
- Introduce an interface only when it removes a real coupling or enables a clear test boundary.
- Avoid over-abstracting one-off code paths; simple concrete code is preferred when it is already stable and readable.
- Keep public APIs backward compatible unless the requested change explicitly requires a break.
- When behavior changes, add or update the nearest package test that proves the new rule.

## Red flags

- Cross-package shortcuts that bypass the intended layer boundaries.
- Shared helpers that know too much about parser or evaluator internals.
- Large interfaces that force unrelated methods onto callers.
- Premature generalization that makes the code harder to read without solving a real duplication problem.
