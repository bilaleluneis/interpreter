from pyimpl import TokenType


def test_tokentype_creation():
    assert TokenType.ASSIGN.value == "="


def test_find_enum_value():
    assert "=" in TokenType
