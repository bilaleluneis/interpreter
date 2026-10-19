use token::token::Token;

use crate::lazy::LazyLexer;

pub trait Lexer: Clone + Copy {
    fn next_token(&mut self) -> Token;
}

// the warning is not correct, as lifetime is needed
#[allow(clippy::needless_lifetimes)]
#[must_use]
pub fn new_lexer(input: &str) -> impl Lexer + '_ {
    LazyLexer::new(input)
}
