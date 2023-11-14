from enum import UNIQUE, verify, StrEnum, EnumMeta
from typing import Any


class __MetaEnum(EnumMeta):
    def __contains__(cls, item: Any) -> bool:
        """
        Checks if the given item is a member of the enum.
        the use of Any for item as type is to allow the use of
        this meta class in other enums.
        """
        try:
            cls(item)
        except ValueError:
            return False
        return True


@verify(UNIQUE)
class TokenType(StrEnum, metaclass=__MetaEnum):
    """Enum representing the possible token types."""

    # Single character tokens.
    LEFT_PAREN = "("
    RIGHT_PAREN = ")"
    LEFT_BRACE = "{"
    RIGHT_BRACE = "}"
    COMMA = ","
    SEMICOLON = ";"

    # Operators
    ASSIGN = "="
    PLUS = "+"

    # Literals.
    IDENTIFIER = "identifier"
    INT = "int"

    # Keywords.
    FUNCTION = "fn"
    LET = "let"

    ILLIGAL = "illigal"
    EOF = "eof"
