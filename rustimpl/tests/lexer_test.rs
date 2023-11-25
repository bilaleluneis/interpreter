use rustimpl::lexer::Lexer;
use rustimpl::token::Token;

#[test]
fn lexer_integration() {
    let mut lex = Lexer::new("".to_string());
    let tok = lex.next_token();
    assert_eq!(tok, Token::Eof);
}
