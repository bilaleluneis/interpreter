---
name: design-principles
description: "Use when designing or implementing Scala code that should stay immutable, SOLID-aligned, serialization-friendly, and test-driven."
argument-hint: "What feature, module, or refactor should be built with immutable, SOLID, and TDD design constraints?"
model: "GPT-5.4 mini"
---

# Design Principles

Use this skill when the work should emphasize immutable data modeling, SOLID design, serialization-friendly types, and test-driven implementation.

## Core Stance

- Create or adjust tests early so the implementation can be driven by failing examples.
- Prefer designs that keep behavior explicit and state transitions easy to reason about.
- Avoid hidden mutation or cross-module coupling that will complicate later refactors.

## Workflow

1. Start from tests.
   - Write or update the smallest failing test that captures the intended behavior.
   - Prefer behavior-level tests over implementation details.
   - Keep the test readable and focused on one rule at a time.

2. Model the data immutably.
   - Prefer `case class`, `enum`, and other value-oriented types.
   - Avoid mutable fields, shared mutable collections, and hidden state.
   - Make state transitions explicit by returning new values rather than mutating in place.

3. Keep types serialization-friendly.
   - Use plain, composable data shapes that are easy to deep copy.
   - Avoid embedding runtime resources, open handles, or environment-bound state in core domain types.
   - Prefer data that can be safely copied, persisted, or moved across process boundaries later.

4. Apply SOLID deliberately.
   - Give each type one clear reason to change.
   - Separate parsing, domain modeling, evaluation, and infrastructure concerns.
   - Depend on small abstractions where it reduces coupling, but avoid unnecessary indirection.
   - Prefer composition over inheritance unless inheritance is the clearest fit.

5. Implement the minimum code to pass.
   - Make the smallest change that satisfies the failing test.
   - Refactor only after the behavior is green.
   - Keep public APIs narrow and stable.
   - If several implementations are possible, choose the one with the clearest correctness story and the smallest blast radius.

6. Verify the change.
   - Run the targeted test first.
   - If the test passes, check for any nearby design issues such as accidental mutability or tight coupling.
   - Add regression tests when the change reveals a new edge case.

## Decision Rules

- If a design requires mutation, first ask whether the mutation can be replaced with a pure return value.
- If a type is hard to serialize or deep copy, simplify the shape before adding behavior.
- If a class has multiple responsibilities, split it before adding more logic.
- If a refactor changes behavior, codify that behavior in tests before broadening the change.

## Completion Check

A change is complete only when:

- The tests describe the intended behavior.
- Core domain types remain immutable.
- Domain types stay serialization-friendly and deep-copy-safe where practical.
- Responsibilities are split cleanly enough that future changes should not force cross-module churn.

## Good Prompts

- "Implement this feature with immutable data and tests first."
- "Refactor this module to remove mutable state and keep it SOLID."
- "Make these domain types safe to serialize and deep copy."
- "Add the smallest TDD change that satisfies this new behavior."
