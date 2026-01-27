#[derive(PartialEq)]
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
