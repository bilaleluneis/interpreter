from pyimpl import TokenType  # type: ignore
from typing import NamedTuple


class Token(NamedTuple):
    token_type: TokenType
    literal: str
