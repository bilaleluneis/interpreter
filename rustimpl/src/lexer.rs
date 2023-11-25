use crate::token::Token;
use std::str;


#[derive(Debug)]
pub struct Lexer {
    input: String,
    prev_read_position: usize,
    next_read_position: usize,
    ch_under_inspection: u8,
}

impl Lexer {
    pub fn new(input: String) -> Self {
        Lexer {
            input,
            prev_read_position: 0,
            next_read_position: 0,
            ch_under_inspection: 0u8,
        }
    }

    pub fn next_token(&mut self) -> Token {
        //TODO need to handle white space
        let pos = self.next_read_position;
        let next_ch = if pos >= self.input.len() {
            self.ch_under_inspection = 0u8;
            "".to_string()
        } else {
            let ch = str::from_utf8(&[self.input.as_bytes()[pos]]).unwrap_or("").to_string();
            self.ch_under_inspection = ch.as_bytes()[0];
            ch
        };
        self.prev_read_position = self.next_read_position;
        self.next_read_position += 1;
        Token::lookup(next_ch.as_str())
    }
}

#[cfg(test)]
mod tests {
    use crate::lexer::Lexer;
    use crate::token::Token;

    #[test]
    fn basic_lexing() {
        let mut lexer = Lexer::new("=+,;(){}".to_string());
        let expected_tokens = [
            Token::Assign,
            Token::Plus,
            Token::Comma,
            Token::Semicolon,
            Token::Lparen,
            Token::Rparen,
            Token::Lbrace,
            Token::Rbrace,
            Token::Eof,
        ];
        for expected_tok in expected_tokens {
            assert_eq!(lexer.next_token(), expected_tok);
        }
    }
}
