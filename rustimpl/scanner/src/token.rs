use std::fmt;
use std::fmt::{Display, Formatter};

#[derive(Debug, PartialEq)]
pub enum Token {
    Illegal(String),
    Ident(String),
    Int(u64),
    Eof,
    Assign,
    Plus,
    Comma,
    Semicolon,
    Lparen,
    Rparen,
    Lbrace,
    Rbrace,
    Function,
    Let,
    Boolean(bool),
}

impl Token {
    pub fn lookup(possible_tok: &str) -> Self {
        match possible_tok {
            "" => Token::Eof,
            "fn" => Token::Function,
            "let" => Token::Let,
            "(" => Token::Lparen,
            ")" => Token::Rparen,
            "{" => Token::Lbrace,
            "}" => Token::Rbrace,
            "," => Token::Comma,
            ";" => Token::Semicolon,
            "+" => Token::Plus,
            "=" => Token::Assign,
            t if t == "true" || t == "false" => Token::Boolean(t == "true"),
            t if t.parse::<u64>().is_ok() => Token::Int(t.parse::<u64>().unwrap()),
            t if t.chars().all(|c| c.is_alphabetic()) => Token::Ident(t.to_string()),
            _ => Token::Illegal(possible_tok.to_string()),
        }
    }
}

impl Display for Token {
    fn fmt(&self, f: &mut Formatter) -> fmt::Result {
        match self {
            Token::Illegal(s) => write!(f, "Illegal({})", s),
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
    fn token_lookup() {
        let tokens = [
            ("", Token::Eof),
            ("fn", Token::Function),
            ("let", Token::Let),
            ("(", Token::Lparen),
            (")", Token::Rparen),
            ("{", Token::Lbrace),
            ("}", Token::Rbrace),
            (",", Token::Comma),
            (";", Token::Semicolon),
            ("+", Token::Plus),
            ("=", Token::Assign),
            ("5", Token::Int(5)),
            ("x", Token::Ident("x".to_string())),
            ("@#", Token::Illegal("@#".to_string())),
            ("true", Token::Boolean(true)),
            ("false", Token::Boolean(false)),
        ];

        for (input, expected) in tokens {
            assert_eq!(Token::lookup(input), expected);
        }
    }
    
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
            assert_eq!(input.to_string(), expected);
        }
    }
    
}
