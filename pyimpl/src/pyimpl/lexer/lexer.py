from pyimpl import Token, TokenType


class Lexer:
    def __init__(self, input_str: str) -> None:
        if len(input_str) == 0:
            raise Exception("empty input provided to Lexer")

        self.input: str = input_str
        self.prev_read_index: int = 0
        self.next_read_index: int = 0
        self.char_under_inspection: bytes = b"0"

    def next_token(self) -> Token:
        self.__read_char()
        while str(self.char_under_inspection).isspace():
            self.__read_char()

        if self.char_under_inspection == b"0":
            return Token(TokenType.EOF, "")

        tok_as_str = str(self.char_under_inspection)
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
            case _:
                return Token(TokenType.ILLIGAL, tok_as_str)

    def __read_char(self):
        if self.next_read_index >= len(self.input):
            self.char_under_inspection = b"0"
        else:
            self.char_under_inspection = self.input[self.next_read_index]
        self.prev_read_index = self.next_read_index
        self.next_read_index += 1

    def __read_literal(self):
        pass
