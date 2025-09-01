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
				let mul = fn(x, y) { x * y };
				let add = fn(x, y) { x + y };
				let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
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

		// let mul = fn(x, y) { x * y };
		{token.LET, "let"},
		{token.IDENTIFIER, "mul"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPRAN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPRAN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.ASTER, "*"},
		{token.IDENTIFIER, "y"},
		{token.RBRACE, "}"},
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

		// let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IF, "if"},
		{token.LPRAN, "("},
		{token.IDENTIFIER, "five"},
		{token.LT, "<"},
		{token.IDENTIFIER, "ten"},
		{token.RPRAN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "mul"},
		{token.LPRAN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPRAN, ")"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "add"},
		{token.LPRAN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPRAN, ")"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// EOF
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
			got := tok.Type
			expected := tokTest.expectedType
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expected, got)
		}

		if tok.Literal != tokTest.expectedLiteral {
			got := tok.Literal
			expected := tokTest.expectedLiteral
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expected, got)
		}
	}
}
