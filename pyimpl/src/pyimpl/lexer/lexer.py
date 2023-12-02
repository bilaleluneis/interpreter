from pyimpl import Token, TokenType  # type: ignore


class Lexer:
    def __init__(self, input_str: str) -> None:
        if len(input_str) == 0:
            raise Exception("empty input provided to Lexer")

        self.__input: str = input_str
        self.__prev_read_index: int = 0
        self.__next_read_index: int = 0
        self.__char_under_inspection: bytes = b''

    def next_token(self) -> Token:
        """Returns the next token in the input string."""
        self.__skip_whitespace()
        self.__read_char()
        tok_as_str = self.str_under_inspection
        match tok_as_str:
            case "":
                return Token(TokenType.EOF, tok_as_str)
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
            case tok if tok.isalpha():
                return self.__spit_token_from_literal()
            case tok if tok.isdigit():
                return Token(TokenType.INT, self.__read_digits())
            case _:
                return Token(TokenType.ILLIGAL, tok_as_str)

    def __read_char(self) -> None:
        if self.__next_read_index >= len(self.__input):
            self.__char_under_inspection = b''
        else:
            self.__char_under_inspection = self.__input[self.__next_read_index].encode("utf-8")
        self.__prev_read_index = self.__next_read_index
        self.__next_read_index += 1

    def __read_digits(self) -> str:
        digits = self.str_under_inspection
        while (ch := self.peek) and ch.isdigit():
            self.__read_char()
            digits += self.str_under_inspection
        return digits

    def __spit_token_from_literal(self) -> Token:
        literal = self.str_under_inspection
        while (ch := self.peek) and ch.isalpha():
            self.__read_char()
            literal += self.str_under_inspection

        if literal in TokenType:
            return Token(TokenType(literal), literal)
        return Token(TokenType.IDENTIFIER, literal)

    def __skip_whitespace(self) -> None:
        while (ch := self.peek) and ch.isspace():
            self.__read_char()

    @property
    def str_under_inspection(self) -> str:
        """Returns the string representation of the character under inspection."""
        if self.__char_under_inspection == b'':
            return ""
        return self.__char_under_inspection.decode("utf-8")

    @property
    def peek(self) -> str:
        """Returns the string representation of the next character to be read."""
        if self.__next_read_index >= len(self.__input):
            return ""
        return self.__input[self.__next_read_index].encode("utf-8").decode("utf-8")
