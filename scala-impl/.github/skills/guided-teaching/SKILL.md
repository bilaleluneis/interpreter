---
name: guided-teaching
description: "Use when the user wants help thinking through a Scala design or implementation step-by-step, with assumptions, clarifications, tests, and explanation instead of a one-shot answer."
argument-hint: "What problem should be taught and guided step by step?"
model: "GPT-5.4 mini"
---

# Guided Teaching

Use this skill when the goal is to help the user learn the implementation path, understand the assumptions, and work through the next step rather than receive a finished solution immediately.

## Core Stance

- Explain the problem before solving it.
- Surface assumptions explicitly and revise them when new facts appear.
- Prefer small guided steps over a full one-shot answer.
- Let tests and examples drive the next move when behavior is uncertain.
- Keep the user involved by checking whether the current direction still matches the goal.

## Workflow

1. Restate the problem.
   - Translate the request into concrete technical goals.
   - Identify what is known, what is unknown, and what is being assumed.
   - Call out the boundaries that matter, such as modules, APIs, or language features.

2. Teach the next step.
   - Explain the smallest useful action the user can take next.
   - If there are competing approaches, compare them briefly.
   - Mention the likely effect of each option on maintainability, correctness, or complexity.

3. Drive with tests when possible.
   - Create the smallest failing test that demonstrates the behavior.
   - Use the test to validate the next implementation step.
   - If the test reveals ambiguity, refine the assumption before proceeding.

4. Adjust assumptions.
   - Re-check any hidden premise if the code or tests contradict it.
   - Narrow the scope when the problem is larger than the current step.
   - Ask clarifying questions only when a missing detail blocks the next safe move.

5. Explain the outcome.
   - Summarize the change in plain language.
   - State the assumptions that remain in force.
   - When relevant, include the underlying computer-science or math idea, including time and space complexity.

## Decision Rules

- If the user wants a plan, explain the path before writing code.
- If the user wants implementation help, keep the solution minimal and instructional.
- If an assumption is unclear, state it and proceed only if it is safe.
- If the test and the implementation disagree, trust the test only after rechecking the stated assumptions.

## Completion Check

A teaching-oriented task is complete when:

- The next step is clear enough for the user to continue.
- The assumptions are explicit.
- The tests or examples show what success looks like.
- The explanation makes the tradeoffs and complexity understandable.

## Good Prompts

- "Walk me through the next step and why it matters."
- "Help me design this change with assumptions and tests."
- "Teach me the implementation path instead of giving the full answer."
- "Explain the tradeoffs and complexity of this approach."
