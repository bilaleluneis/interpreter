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
        while self.peek().is_ascii_whitespace() {
            self.advance()
        }

        match self.next_char().as_str() {
            ch if ch.is_empty() => Token::Eof,
            ch if ch.chars().all(char::is_alphabetic) => self.literal_token(),
            ch if ch.chars().all(char::is_numeric) => self.numeric_token(),
            ch => Token::lookup(ch),
        }
    }

    fn literal_token(&mut self) -> Token {
        let mut ch = self.byte_to_str();
        while !self.peek().is_ascii_whitespace() && self.peek().is_ascii_alphabetic() {
            ch += self.next_char().as_str();
        }
        let tok = Token::lookup(ch.as_str());
        match tok {
            Token::Illegal(t) => Token::Ident(t),
            _ => tok
        }
    }

    fn numeric_token(&mut self) -> Token {
        let mut ch = self.byte_to_str();
        while !self.peek().is_ascii_whitespace() && self.peek().is_ascii_digit() {
            ch += self.next_char().as_str();
        }
        Token::lookup(ch.as_str()) // shoult return Int or Illigal token
    }

    // TODO:can you return char instead of string as utf8?
    fn next_char(&mut self) -> String {
        self.ch_under_inspection = self.peek();
        self.advance();
        self.byte_to_str()
    }

    fn peek(&self) -> u8 {
        let pos = self.next_read_position;
        let ch = if pos >= self.input.len() {
            0u8
        } else {
            self.input.as_bytes()[pos]
        };
        ch
    }

    //TODO: come up with better name
    fn byte_to_str(&self) -> String {
        match self.ch_under_inspection {
            0u8 => "".to_string(),
            b => str::from_utf8(&[b]).unwrap_or("").to_string()
        }
    }

    fn advance(&mut self) {
        self.prev_read_position = self.next_read_position;
        self.next_read_position += 1;
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

    #[test]
    fn min_lang_construct() {
        let mut lexer = Lexer::new(r#"
            let five = 5;
            let ten = 10;
            let add = fn(x, y) { x + y};
            let result = add(five, ten);
        "#.to_string());

        let expected_tokens = [
            // let five = 5;
            Token::Let,
            Token::Ident("five".to_string()),
            Token::Assign,
            Token::Int(5),
            Token::Semicolon,

            //let ten = 10;
            Token::Let,
            Token::Ident("ten".to_string()),
            Token::Assign,
            Token::Int(10),
            Token::Semicolon,

            //let add = fn(x, y) {x, y};
            Token::Let,
            Token::Ident("add".to_string()),
            Token::Assign,
            Token::Function,
            Token::Lparen,
            Token::Ident("x".to_string()),
            Token::Comma,
            Token::Ident("y".to_string()),
            Token::Rparen,
            Token::Lbrace,
            Token::Ident("x".to_string()),
            Token::Plus,
            Token::Ident("y".to_string()),
            Token::Rbrace,
            Token::Semicolon,
        ];
        for expected_tok in expected_tokens {
            assert_eq!(lexer.next_token(), expected_tok);
        }
    }
}
