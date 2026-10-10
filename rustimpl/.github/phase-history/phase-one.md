# Phase One History

## Goal

Implement the minimal interpreter language for the rustimpl workspace.

## Language Sample

```rust
let five = 5;
let ten = 10;
let mul = fn(x, y) { x * y };
let add = fn(x, y) { x + y };
let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
```

## Tracking Notes

- Keep all phase-one decisions that affect token shapes, parsing rules, or evaluator behavior here.
- Add concise notes for completed milestones, behavioral changes, and any deliberate limitations.
- When phase two starts, create a sibling history file instead of replacing this one.
