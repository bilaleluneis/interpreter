from pyimpl import TokenType  # type: ignore


def test_tokentype_creation() -> None:
    assert TokenType.ASSIGN.value == "="


def test_find_enum_value() -> None:
    assert "=" in TokenType
