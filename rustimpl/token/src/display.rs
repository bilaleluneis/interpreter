use std::fmt::{Display, Formatter, Result};

use crate::token::Token;

impl Display for Token {
    fn fmt(&self, f: &mut Formatter) -> Result {
        match self {
            Self::Illegal(s) => write!(f, "Illegal({s})"),
            Self::Ident(s) => write!(f, "Ident({s})"),
            Self::Int(i) => write!(f, "Int({i})"),
            Self::Eof => write!(f, "Eof"),
            Self::Assign => write!(f, "Assign"),
            Self::Plus => write!(f, "Plus"),
            Self::Asterisk => write!(f, "Asterisk"),
            Self::Lt => write!(f, "Lt"),
            Self::Comma => write!(f, "Comma"),
            Self::Semicolon => write!(f, "Semicolon"),
            Self::Lparen => write!(f, "Lparen"),
            Self::Rparen => write!(f, "Rparen"),
            Self::Lbrace => write!(f, "Lbrace"),
            Self::Rbrace => write!(f, "Rbrace"),
            Self::Function => write!(f, "Function"),
            Self::Let => write!(f, "Let"),
            Self::If => write!(f, "If"),
            Self::Else => write!(f, "Else"),
            Self::Boolean(b) => write!(f, "Boolean({b})"),
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::token::Token;

    #[test]
    fn token_display() {
        let tokens = [
            (Token::Illegal(String::new()), "Illegal()"),
            (Token::Ident("x".to_string()), "Ident(x)"),
            (Token::Int(5), "Int(5)"),
            (Token::Eof, "Eof"),
            (Token::Assign, "Assign"),
            (Token::Plus, "Plus"),
            (Token::Asterisk, "Asterisk"),
            (Token::Lt, "Lt"),
            (Token::Comma, "Comma"),
            (Token::Semicolon, "Semicolon"),
            (Token::Lparen, "Lparen"),
            (Token::Rparen, "Rparen"),
            (Token::Lbrace, "Lbrace"),
            (Token::Rbrace, "Rbrace"),
            (Token::Function, "Function"),
            (Token::Let, "Let"),
            (Token::If, "If"),
            (Token::Else, "Else"),
            (Token::Boolean(true), "Boolean(true)"),
            (Token::Boolean(false), "Boolean(false)"),
        ];

        for (input, expected) in tokens {
            let _ = format!("{input:?}");
            assert_eq!(input.to_string(), expected);
        }
    }
}
