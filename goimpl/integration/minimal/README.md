 
# Minimal integration tests — Monkey (minimal)

This folder contains a small integration test suite that verifies end-to-end support
for a minimal subset of the Monkey language: tokenization, lexing, parsing and evaluation.

Purpose
- Validate the interpreter can fully handle this sample program:

```
let five = 5;
let ten = 10;
let mul = fn(x, y) { x * y };
let add = fn(x, y) { x + y };
let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
```

Scope
- Exercises: `let` bindings, integer literals, function literals (`fn`), function calls,
	infix operators (`*`, `+`, `<`), `if ... else` expressions and block bodies.
- Intentionally limited: does not cover strings, arrays, hashes, comments or other extensions.

How to run
- From the repository root (workspace mode using `go.work`), run:

```bash
cd goimpl/integration/minimal
go test ./...
```

Expected outcome
- Tests should confirm the sample program is tokenized, lexed, parsed into an AST, and evaluated
	with the correct runtime semantics for the `result` binding.

Notes
- This integration suite imports local `goimpl/*` modules via `go.work` so it runs against
	the workspace modules. Keep tests deterministic and free of external side effects.

