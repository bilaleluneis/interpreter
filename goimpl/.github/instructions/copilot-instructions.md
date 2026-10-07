## Go Impl Workspace Instructions

Apply these instructions when editing this workspace.

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
