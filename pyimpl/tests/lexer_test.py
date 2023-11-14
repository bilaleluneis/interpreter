import pytest

from pyimpl import Token, TokenType, Lexer  # type: ignore


@pytest.mark.parametrize("literal, expected", [
    ("=", TokenType.ASSIGN),
    ("+", TokenType.PLUS),
    ("(", TokenType.LEFT_PAREN),
    (")", TokenType.RIGHT_PAREN),
    ("{", TokenType.LEFT_BRACE),
    ("}", TokenType.RIGHT_BRACE),
    (",", TokenType.COMMA),
    (";", TokenType.SEMICOLON),
    ("fn", TokenType.FUNCTION),
    ("let", TokenType.LET),
    ("x", TokenType.IDENTIFIER),
    ("int", TokenType.INT),
])
def test_token_emitted(literal: str, expected: TokenType) -> None:
    lexer = Lexer(literal)
    token = lexer.next_token()
    assert token.token_type == expected


def test_min_language_construct() -> None:
    input_tokens = ("   let five = 5;\n"
                    "   let ten = 10;\n"
                    "	let add = fn(x, y) { x + y };\n"
                    "	let result = add(five, ten);    ")

    expected_tokens = [
        # let five = 5;
        Token(TokenType.LET, "let"),
        Token(TokenType.IDENTIFIER, "five"),
        Token(TokenType.ASSIGN, "="),
        Token(TokenType.INT, "5"),
        Token(TokenType.SEMICOLON, ";"),

        # let ten = 10;
        Token(TokenType.LET, "let"),
        Token(TokenType.IDENTIFIER, "ten"),
        Token(TokenType.ASSIGN, "="),
        Token(TokenType.INT, "10"),
        Token(TokenType.SEMICOLON, ";"),

        # let add = fn(x, y) { x + y };
        Token(TokenType.LET, "let"),
        Token(TokenType.IDENTIFIER, "add"),
        Token(TokenType.ASSIGN, "="),
        Token(TokenType.FUNCTION, "fn"),
        Token(TokenType.LEFT_PAREN, "("),
        Token(TokenType.IDENTIFIER, "x"),
        Token(TokenType.COMMA, ","),
        Token(TokenType.IDENTIFIER, "y"),
        Token(TokenType.RIGHT_PAREN, ")"),
        Token(TokenType.LEFT_BRACE, "{"),
        Token(TokenType.IDENTIFIER, "x"),
        Token(TokenType.PLUS, "+"),
        Token(TokenType.IDENTIFIER, "y"),
        Token(TokenType.RIGHT_BRACE, "}"),
        Token(TokenType.SEMICOLON, ";"),

        # let result = add(five, ten);
        Token(TokenType.LET, "let"),
        Token(TokenType.IDENTIFIER, "result"),
        Token(TokenType.ASSIGN, "="),
        Token(TokenType.IDENTIFIER, "add"),
        Token(TokenType.LEFT_PAREN, "("),
        Token(TokenType.IDENTIFIER, "five"),
        Token(TokenType.COMMA, ","),
        Token(TokenType.IDENTIFIER, "ten"),
        Token(TokenType.RIGHT_PAREN, ")"),
        Token(TokenType.SEMICOLON, ";"),

        Token(TokenType.EOF, ""),
    ]

    assert lexable(input_tokens, expected_tokens)


def lexable(input_str: str, expected_tok: list[Token]) -> bool:
    lexer = Lexer(input_str)
    for index, expected_token in enumerate(expected_tok):
        token = lexer.next_token()
        if token != expected_token:
            print(f"\nexpected token: {expected_token} at index: {index}")
            print(f"actual token: {token}")
            return False
    return True
