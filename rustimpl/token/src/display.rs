use std::fmt::{Display, Formatter, Result};

use crate::token::Token;

impl Display for Token {
    fn fmt(&self, f: &mut Formatter) -> Result {
        match self {
            Token::Illegal(s) => write!(f, "Duh..Illegal({})", s),
            Token::Ident(s) => write!(f, "Ident({})", s),
            Token::Int(i) => write!(f, "Int({})", i),
            Token::Eof => write!(f, "Eof"),
            Token::Assign => write!(f, "Assign"),
            Token::Plus => write!(f, "Plus"),
            Token::Comma => write!(f, "Comma"),
            Token::Semicolon => write!(f, "Semicolon"),
            Token::Lparen => write!(f, "Lparen"),
            Token::Rparen => write!(f, "Rparen"),
            Token::Lbrace => write!(f, "Lbrace"),
            Token::Rbrace => write!(f, "Rbrace"),
            Token::Function => write!(f, "Function"),
            Token::Let => write!(f, "Let"),
            Token::Boolean(b) => write!(f, "Boolean({})", b),
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::token::Token;

    #[test]
    fn token_display() {
        let tokens = [
            (Token::Illegal("".to_string()), "Illegal()"),
            (Token::Ident("x".to_string()), "Ident(x)"),
            (Token::Int(5), "Int(5)"),
            (Token::Eof, "Eof"),
            (Token::Assign, "Assign"),
            (Token::Plus, "Plus"),
            (Token::Comma, "Comma"),
            (Token::Semicolon, "Semicolon"),
            (Token::Lparen, "Lparen"),
            (Token::Rparen, "Rparen"),
            (Token::Lbrace, "Lbrace"),
            (Token::Rbrace, "Rbrace"),
            (Token::Function, "Function"),
            (Token::Let, "Let"),
            (Token::Boolean(true), "Boolean(true)"),
            (Token::Boolean(false), "Boolean(false)"),
        ];

        for (input, expected) in tokens {
            let _ = format!("{:?}", input);
            assert_eq!(input.to_string(), expected);
        }
    }
}
