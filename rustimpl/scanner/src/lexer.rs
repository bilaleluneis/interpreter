use token::token::Token;

use crate::lazy::LazyLexer;

/// A token source that yields the next token on demand.
pub trait Lexer: Clone + Copy {
    /// Returns the next token from the underlying input.
    fn next_token(&mut self) -> Token;
}

// the warning is not correct, as lifetime is needed
#[allow(clippy::needless_lifetimes)]
#[must_use]
/// Creates a lexer for the given input text.
pub fn new_lexer(input: &str) -> impl Lexer + '_ {
    LazyLexer::new(input)
}
