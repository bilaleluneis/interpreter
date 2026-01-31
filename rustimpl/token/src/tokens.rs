use crate::token::Token;

#[allow(dead_code)]
trait TokenMatcher: AsRef<[Token]> {
    fn contains(&self, token: &Token) -> bool {
        self.as_ref().iter().any(|t| t == token)
    }
}

impl<T: AsRef<[Token]>> TokenMatcher for T {}

#[cfg(test)]
mod tests {
    use crate::{token::Token, tokens::TokenMatcher};

    fn sample_tokens() -> Vec<Token> {
        vec![
            Token::Illegal("".to_string()),
            Token::Ident("x".to_string()),
            Token::Int(42),
            Token::Eof,
            Token::Assign,
            Token::Plus,
            Token::Comma,
            Token::Semicolon,
            Token::Lparen,
            Token::Rparen,
            Token::Lbrace,
            Token::Rbrace,
            Token::Function,
            Token::Let,
            Token::Boolean(true),
            Token::Boolean(false),
        ]
    }

    #[test]
    fn finds_match() {
        assert!(sample_tokens().contains(&Token::Ident("x".to_string())));
        assert!(sample_tokens().contains(&Token::Int(42)));
        assert!(sample_tokens().contains(&Token::Eof));
    }

    #[test]
    fn no_match() {
        assert!(!sample_tokens().contains(&Token::Ident("y".to_string())));
        assert!(!sample_tokens().contains(&Token::Int(100)));
    }
}
