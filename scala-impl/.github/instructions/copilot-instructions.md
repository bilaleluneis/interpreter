# MonkeyLang Scala Interpreter Guidance

Treat this repository as a Scala 3.9.0 multi-module Maven project.

- Prefer Scala 3.9 features first when proposing solutions; only fall back to older Scala 3.x features if 3.9 cannot express the design cleanly.
- Use MUnit for tests.
- Keep module boundaries tight to reduce recompilation and avoid cyclic dependencies.

Current module direction:

- `token` is the lowest-level module and owns lexical token definitions and keyword lookup.
- `ast` owns syntax tree and language model types.
- `lexer` may depend on `token`, but not on parser or evaluator internals.
- `parser` may depend on `token` and `ast`, but should not reach back into lexer implementation details.
- Future evaluator/runtime code should depend on `ast` and shared low-level types only; do not create upward dependencies from foundational modules to runtime code.

Phase one is an MVP evaluator/interpreter for the Monkey/Miniml subset. Keep changes focused on the minimum language needed for expressions like:

```text
let five = 5;
let ten = 10;
let mul = fn(x, y) { x * y };
let add = fn(x, y) { x + y };
let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
```

- Preserve a clean path for later cluster/distributed execution by keeping evaluator semantics and module APIs as pure and serializable as practical.
- When adding or reorganizing modules, update the root `pom.xml` first and keep dependencies flowing downward through the layer order above.
- Prefer the smallest module that owns the behavior being changed.
