from pyimpl import TokenType  # type: ignore


class Token:
    """Represents a token in a simple programming language."""

    def __init__(self, token_type: TokenType, literal: str) -> None:
        """Initializes a new instance of the Token class."""
        self.__token_type = token_type
        self.__literal = literal

    def __eq__(self, other: object) -> bool:
        """Checks if the current token is equal to the other token."""
        if not isinstance(other, Token):
            return False

        return self.__token_type == other.__token_type and self.__literal == other.__literal

    def __repr__(self) -> str:
        """Returns the string representation of the token."""
        return f"Token({self.__token_type}, {self.__literal})"

    @property
    def token_type(self) -> TokenType:
        """Gets the token type."""
        return self.__token_type
