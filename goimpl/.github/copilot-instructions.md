## Go Impl Repository Instructions

Apply these instructions repository-wide whenever this workspace is active, including in chat, agents, code suggestions, and when making code changes.

These instructions are the default baseline for the workspace. Agents and skills should follow them unless a more specific instruction file, skill, or user request overrides a rule.

### Tooling and language level

- Target Go 1.27 for all code in this workspace.
- The repository is organized as a Go workspace with separate modules under `ast/`, `common/`, `eval/`, `lexer/`, `parser/`, `token/`, and `integration/`.
- Prefer the standard library. Do not add third-party dependencies unless the user explicitly asks for one.

### Code style and design

- Keep changes small and local to the package that owns the behavior.
- Follow the existing package-by-package design: AST types live in `ast`, token definitions in `token`, parsing logic in `parser`, lexing in `lexer`, and evaluation logic in `eval`.
- Preserve the current use of interfaces and marker methods for AST nodes, and use pointer receivers for stateful types while keeping simple value types as values.
- Match the repository's generics usage where it already exists, especially in `common`, but avoid introducing generics where a simple concrete type is clearer.
- Prefer explicit error returns and existing AST error types over panics for recoverable parsing and evaluation failures.
- Keep code within 90 columns when practical, including statements and comments.
- Keep exported comments short and descriptive when adding new public identifiers.

### Testing and validation

- Use table-driven tests when adding or updating behavior.
- Keep tests close to the package being changed, following the current `_test.go` layout.
- Run `gofmt` on any Go files you edit.
- Validate changes with `go test ./...` when practical, or at least with the narrowest package-level test command that covers the touched code.

### Working conventions

- Do not leave placeholder implementations, `TODO`s, or stubbed return values unless the user asked for a scaffold.
- Prefer the existing naming style, even where spelling is inconsistent, unless a change is clearly part of the requested fix.
- Keep changes compatible with the current public APIs unless the user explicitly requests a breaking change.
- When answering questions or providing code suggestions about Go language features that may be newer than the model's built-in knowledge, or newer than the current supported baseline, consult `.github/references/features.md` before responding.
- Treat `.github/references/features.md` as the source of truth for project-specific Go 1.26 and Go 1.27 feature guidance.
- When suggesting or explaining Go language features, consult `.github/references/features.md` and use it eagerly when it helps provide a better solution or example.
- When a solution pattern matches one covered in `.github/references/best-practices.md`, consult that file and prefer the shaped solution it recommends when it improves clarity, composability, or maintainability.
