package minimal

import (
	"goimpl/lexer"
	"goimpl/token"
	"testing"
)

// Test lexing:
// let five = 5;
// let ten = 10;
// let mul = fn(x, y) { x * y };
// let add = fn(x, y) { x + y };
// let result = if(five < ten) { mul(five, ten) } else { add(five, ten };

func TestLexing(t *testing.T) {

	input := `	let five = 5;
				let ten = 10;
				let mul = fn(x, y) { x * y };
				let add = fn(x, y) { x + y };
				let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
			`

	fixture := []struct {
		tokType         token.TokenType
		expectedLiteral string
	}{
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

	l := lexer.NewLazyLexer(input)

	for i, tt := range fixture {
		t.Run(tt.expectedLiteral, func(t *testing.T) {
			tok := l.NextToken()
			expect := tt.tokType
			got := tok.Type

			if got != expect {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, expect, got)
			}

			expectLit := tt.expectedLiteral
			gotLit := tok.Literal

			if gotLit != expectLit {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, expectLit, gotLit)
			}
		})
	}

}
