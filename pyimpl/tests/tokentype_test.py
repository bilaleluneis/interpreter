from pyimpl import TokenType  # type: ignore


def test_compare_str_value_to_token() -> None:
    assert TokenType.ASSIGN.value == "="


def test_find_enum_value_of_str() -> None:
    assert "=" in TokenType
