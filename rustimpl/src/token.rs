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
            t if t.parse::<u64>().is_ok() => Token::Int(t.parse::<u64>().unwrap()),
            t if t.chars().all(|c| c.is_alphabetic()) => Token::Ident(t.to_string()),
            _ => Token::Illegal(possible_tok.to_string()),
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
        ];

        for (input, expected) in tokens {
            assert_eq!(Token::lookup(input), expected);
        }
    }
}
