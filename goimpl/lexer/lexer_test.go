package lexer

import (
	"goimpl/token"
	"testing"
)

func TestSimpleTokens(t *testing.T) {

	input := "=+,;(){}"

	tests := []expectedToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LPRAN, "("},
		{token.RPRAN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	testLexer(input, tests, t)

}

func TestMinLanguageConstruct(t *testing.T) {

	input := `	let five = 5;
				let ten = 10;
				let add = fn(x, y) { x + y };
				let result = add(five, ten);
			`

	tests := []expectedToken{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// let ten = 10;
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// let add = fn(x, y) { x + y };
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPRAN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPRAN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPRAN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPRAN, ")"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	testLexer(input, tests, t)

}

type expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testLexer(input string, tests []expectedToken, t *testing.T) {
	l := NewLazyLexer(input)
	for i, tokTest := range tests {
		tok := l.NextToken()

		if tok.Type != tokTest.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tokTest.expectedType, tok.Type)
		}

		if tok.Literal != tokTest.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tokTest.expectedLiteral, tok.Literal)
		}
	}
}
