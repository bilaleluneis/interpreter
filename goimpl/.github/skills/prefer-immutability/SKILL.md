---
name: prefer-immutability
description: Prefer value semantics and defensive copying; use encoding/gob for deep copies when structs cannot be made truly immutable.
---

# Prefer Immutability

Use this skill when changing code that should avoid shared mutable state.

## Core rule

- Prefer value semantics over shared mutable references.
- Do not claim that Go structs are truly immutable; for most structs, the practical goal is to avoid accidental aliasing and observable mutation.
- When a value must be reused safely, create a defensive copy instead of sharing the original reference.

## How this workspace applies immutability

- The repo already uses `common.DeepCopy` and `common.Serialize` as the main copy boundaries.
- Deep copies are implemented with `encoding/gob`, so types that participate in safe copying should be gob-serializable or provide custom `GobEncode` and `GobDecode` behavior.
- Favor immutable-by-convention runtime values, such as the value types in `eval`, over pointer-heavy shared state.
- Prefer returning values rather than exposing internal pointers when the type can remain a value.
- Keep state ownership local to the package that mutates it.

## Decision rules

- Use a value type when copying it is cheap and behavior is clearer with value semantics.
- Use a pointer only when the type has identity, mutability, or a clearly shared lifecycle.
- If a type contains slices, maps, pointers, or interfaces, assume a shallow copy is not enough.
- If a struct cannot be made truly immutable, enforce immutability by copying on input, on output, or both.
- Use `DeepCopy` or `Serialize` when a defensive deep copy is required and gob support is available.

## Gob guidance

- Prefer gob-based copying for types that need a full deep copy across nested values.
- Add custom gob encode/decode methods when a type has unexported fields or needs explicit control over its serialized representation.
- Do not rely on gob for unsupported values such as functions, channels, or types that cannot be encoded.
- If gob cannot copy a value safely, return or propagate an error instead of silently reusing the original reference.

## Code review checklist

- No function should hand out a mutable alias when a copy is expected.
- No package should mutate a value that another package still owns.
- No helper should promise true immutability for a struct that still contains shared references.
- New APIs should make the copy boundary obvious in the signature or documentation.
