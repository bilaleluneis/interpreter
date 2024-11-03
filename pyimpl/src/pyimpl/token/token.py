"""Token defines the Token class used in the pyimpl interpreter."""

from typing import NamedTuple

from pyimpl import TokenType


class Token(NamedTuple):
    """Token represents a token in the pyimpl language."""

    token_type: TokenType
    literal: str
