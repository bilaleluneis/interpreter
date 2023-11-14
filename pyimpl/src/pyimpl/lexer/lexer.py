from pyimpl import Token, TokenType  # type: ignore


class Lexer:
    def __init__(self, input_str: str) -> None:
        if len(input_str) == 0:
            raise Exception("empty input provided to Lexer")

        self.__input: str = input_str
        self.__prev_read_index: int = 0
        self.__next_read_index: int = 0
        self.__char_under_inspection: bytes = b"0"

    def next_token(self) -> Token:
        self.__read_char()
        while str(self.__char_under_inspection).isspace():
            self.__read_char()

        if self.__char_under_inspection == b"0":
            return Token(TokenType.EOF, "")

        tok_as_str = self.__char_under_inspection.decode("utf-8")
        match tok_as_str:
            case "=":
                return Token(TokenType.ASSIGN, tok_as_str)
            case "+":
                return Token(TokenType.PLUS, tok_as_str)
            case "(":
                return Token(TokenType.LEFT_PAREN, tok_as_str)
            case ")":
                return Token(TokenType.RIGHT_PAREN, tok_as_str)
            case "{":
                return Token(TokenType.LEFT_BRACE, tok_as_str)
            case "}":
                return Token(TokenType.RIGHT_BRACE, tok_as_str)
            case ",":
                return Token(TokenType.COMMA, tok_as_str)
            case ";":
                return Token(TokenType.SEMICOLON, tok_as_str)
            case _ if tok_as_str.isalpha():
                return self.__spit_token_from_literal()
            case _:
                return Token(TokenType.ILLIGAL, tok_as_str)

    def __read_char(self) -> None:
        if self.__next_read_index >= len(self.__input):
            self.__char_under_inspection = b"0"
        else:
            self.__char_under_inspection = self.__input[self.__next_read_index].encode("utf-8")
        self.__prev_read_index = self.__next_read_index
        self.__next_read_index += 1

    def __spit_token_from_literal(self) -> Token:
        literal: str = self.__char_under_inspection.decode("utf-8")
        self.__read_char()
        while (ch := self.__char_under_inspection.decode("utf-8")) and not (ch.isspace() or ch == '0'):
            literal += ch
            self.__read_char()
        if literal in TokenType:
            return Token(TokenType(literal), literal)
        return Token(TokenType.IDENTIFIER, literal)
