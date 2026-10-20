//! Token definitions and keyword lookup for the Monkey language frontend.

mod debug;
mod display;
/// Keyword and literal lookup helpers for tokens.
pub mod lookup;
/// Core token enum used by scanner and parser.
pub mod token;
mod tokens;
