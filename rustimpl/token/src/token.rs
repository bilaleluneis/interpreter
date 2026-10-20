#[derive(PartialEq, Eq)]
/// Lexical token produced by scanning Monkey source text.
pub enum Token {
    /// A token that cannot be recognized by the scanner.
    Illegal(String),
    /// An identifier name.
    Ident(String),
    /// An unsigned integer literal.
    Int(u64),
    /// End-of-input marker.
    Eof,
    /// Assignment operator (`=`).
    Assign,
    /// Addition operator (`+`).
    Plus,
    /// Multiplication operator (`*`).
    Asterisk,
    /// Less-than operator (`<`).
    Lt,
    /// Comma delimiter.
    Comma,
    /// Semicolon delimiter.
    Semicolon,
    /// Left parenthesis (`(`).
    Lparen,
    /// Right parenthesis (`)`).
    Rparen,
    /// Left brace (`{`).
    Lbrace,
    /// Right brace (`}`).
    Rbrace,
    /// `fn` keyword.
    Function,
    /// `let` keyword.
    Let,
    /// `if` keyword.
    If,
    /// `else` keyword.
    Else,
    /// Boolean literal.
    Boolean(bool),
}
