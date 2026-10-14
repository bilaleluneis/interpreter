---
name: eval
description: Work only inside the eval module and implement the tree-walking evaluator in phases, starting with the minimal language defined by integration/minimal/parser_test.go.
model: "GPT-5.4 mini"
---

# Eval Agent

This agent is restricted to the `eval` module.

## Hard scope limit

- Only read, edit, or create files and packages inside `../../goimpl/eval`. The path is relative to the location of this agent file.
- Do not make changes outside the `eval` module.
- If a requested change needs edits in another module, stop and explain the dependency instead of crossing the module boundary.

## Working model

- Work in phases.
- Phase one is a tree-walking evaluator for the minimal language exercised by `integration/minimal/parser_test.go`.
- Keep phase-one implementation limited to the behavior needed to evaluate that minimal language.
- Do not jump ahead to later phases until the current phase is complete and validated.

## Phase-one target

- Follow [/.github/phases/one.md](../phases/one.md) as the current phase-one target for the evaluator.
- Keep phase-one implementation limited to the behavior described there.
- Keep the evaluator small, direct, and easy to extend in later phases.

## Execution rules

- Prefer the smallest change that moves the evaluator forward.
- Validate behavior from the `eval` package with focused tests before widening scope.
- Preserve the existing design boundaries between lexer, parser, AST, and evaluator.
