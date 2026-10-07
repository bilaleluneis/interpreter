---
name: enforce-tdd
description: Enforce test-driven development during implementation by following a strict red-green-refactor loop.
---

# Enforce TDD

Use this skill when development should follow a test-first workflow.

## TDD loop to follow

1. Start from the smallest observable behavior change.
2. Write or update the nearest test first, in the package that owns the behavior.
3. Run the narrowest test that proves the new expectation fails for the right reason.
4. Implement the smallest code change that makes that test pass.
5. Re-run the same focused test.
6. Refactor only after the test is green.
7. Expand to the next closest test only when the current slice is stable.

## How to enforce TDD in this workspace

- Prefer table-driven tests because the repo already uses them heavily.
- Keep tests close to the code under change, such as `lexer/*_test.go`, `parser/*_test.go`, `eval/*_test.go`, and `common/*_test.go`.
- Match the existing test style of explicit assertions and direct error messages.
- When a behavior is missing, add a failing test case before touching the implementation.
- If an implementation already exists but is incorrect, write the regression test first and then fix the code.
- Use the smallest test command that exercises the change; widen to `go test ./...` only after the local slice passes.
- Run `gofmt` on edited Go files before the final validation pass.

## Development workflow rules

- Do not implement extra behavior beyond the active failing test.
- Do not refactor unrelated code while a test is still failing.
- Keep one test failure at a time so the cause of failure stays clear.
- If a test becomes hard to express, stop and simplify the design instead of bypassing the test.
- Prefer small, verifiable increments over large multi-package edits.
- When changing public behavior, add a regression test that will fail if the bug returns.

## Validation ladder

- First choice: a single package test or single test function that exercises the changed slice.
- Second choice: the package test suite for the touched package.
- Final check: `go test ./...` once the local change is stable.

## Signals that TDD is being followed correctly

- The test fails before implementation changes are made.
- The implementation change is small and directly addresses the failing assertion.
- The same test passes after the code change.
- Any later refactor keeps the test passing without changing the observable behavior.
