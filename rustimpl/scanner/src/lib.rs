//! Scanner crate for turning source text into a stream of tokens.

#![allow(dead_code)]
extern crate core;

/// Lazy, byte-wise lexer implementation.
pub mod lazy;
/// Lexer trait and constructor helpers.
pub mod lexer;
