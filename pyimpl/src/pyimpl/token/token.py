from pyimpl import TokenType


class Token:
    """Represents a token in a simple programming language."""

    def __init__(self, token_type: TokenType, literal: str) -> None:
        """Initializes a new instance of the Token class."""
        self.token_type = token_type
        self.literal = literal

    def __eq__(self, other: object) -> bool:
        """Checks if the current token is equal to the other token."""
        if not isinstance(other, Token):
            return False

        return self.token_type == other.token_type and self.literal == other.literal

    def __repr__(self) -> str:
        """Returns the string representation of the token."""
        return f"Token({self.token_type}, {self.literal})"
