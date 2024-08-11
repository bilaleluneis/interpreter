use crate::lazy::LazyLexer;
use crate::token::Token;

pub trait Lexer: Clone + Copy {
    fn next_token(&mut self) -> Token;
}

// the warning is not correct, as lifetime is needed
#[allow(clippy::needless_lifetimes)]
pub fn new_lexer<'a>(input: &'a str) -> impl Lexer + 'a {
    LazyLexer::new(input)
}




