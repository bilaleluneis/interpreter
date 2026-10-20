use crate::token::Token;

impl Token {
    /// Maps a source lexeme into its corresponding [`Token`].
    pub fn lookup(possible_tok: &str) -> Self {
        match possible_tok {
            "" => Self::Eof,
            "fn" => Self::Function,
            "let" => Self::Let,
            "(" => Self::Lparen,
            ")" => Self::Rparen,
            "{" => Self::Lbrace,
            "}" => Self::Rbrace,
            "," => Self::Comma,
            ";" => Self::Semicolon,
            "+" => Self::Plus,
            "*" => Self::Asterisk,
            "<" => Self::Lt,
            "if" => Self::If,
            "else" => Self::Else,
            "=" => Self::Assign,
            t if t == "true" || t == "false" => Self::Boolean(t == "true"),
            t if t.chars().all(char::is_alphabetic) => Self::Ident(t.to_string()),
            _ => possible_tok
                .parse::<u64>()
                .map_or_else(|_| Self::Illegal(possible_tok.to_string()), Self::Int),
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
            ("*", Token::Asterisk),
            ("<", Token::Lt),
            ("if", Token::If),
            ("else", Token::Else),
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
}
