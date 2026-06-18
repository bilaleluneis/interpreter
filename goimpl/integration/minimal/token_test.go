package minimal

import (
	"goimpl/token"
	"testing"
)

// Test tokenizing:
// let five = 5;
// let ten = 10;
// let mul = fn(x, y) { x * y };
// let add = fn(x, y) { x + y };
// let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };

func TestToken(t *testing.T) {
	fixture := []struct {
		lookup string
		want   token.TokenType
	}{
		// let five = 5;
		{"let", token.LET},
		{"five", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"5", token.INT},
		{";", token.SEMICOLON},

		// let ten = 10;
		{"let", token.LET},
		{"ten", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"10", token.INT},
		{";", token.SEMICOLON},

		// let mul = fn(x, y) { x * y };
		{"let", token.LET},
		{"mul", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"fn", token.FUNCTION},
		{"(", token.LPRAN},
		{"x", token.IDENTIFIER},
		{",", token.COMMA},
		{"y", token.IDENTIFIER},
		{")", token.RPRAN},
		{"{", token.LBRACE},
		{"x", token.IDENTIFIER},
		{"*", token.ASTER},
		{"y", token.IDENTIFIER},
		{"}", token.RBRACE},
		{";", token.SEMICOLON},

		// let add = fn(x, y) { x + y };
		{"let", token.LET},
		{"add", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"fn", token.FUNCTION},
		{"(", token.LPRAN},
		{"x", token.IDENTIFIER},
		{",", token.COMMA},
		{"y", token.IDENTIFIER},
		{")", token.RPRAN},
		{"{", token.LBRACE},
		{"x", token.IDENTIFIER},
		{"+", token.PLUS},
		{"y", token.IDENTIFIER},
		{"}", token.RBRACE},
		{";", token.SEMICOLON},

		// let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
		{"let", token.LET},
		{"result", token.IDENTIFIER},
		{"=", token.ASSIGN},
		{"if", token.IF},
		{"(", token.LPRAN},
		{"five", token.IDENTIFIER},
		{"<", token.LT},
		{"ten", token.IDENTIFIER},
		{")", token.RPRAN},
		{"{", token.LBRACE},
		{"mul", token.IDENTIFIER},
		{"(", token.LPRAN},
		{"five", token.IDENTIFIER},
		{",", token.COMMA},
		{"ten", token.IDENTIFIER},
		{")", token.RPRAN},
		{"}", token.RBRACE},
		{"else", token.ELSE},
		{"{", token.LBRACE},
		{"add", token.IDENTIFIER},
		{"(", token.LPRAN},
		{"five", token.IDENTIFIER},
		{",", token.COMMA},
		{"ten", token.IDENTIFIER},
		{")", token.RPRAN},
		{"}", token.RBRACE},
		{";", token.SEMICOLON},
	}

	for i, fx := range fixture {
		t.Run(fx.lookup, func(t *testing.T) {
			tokType := token.Lookup(fx.lookup).Type
			if tokType != fx.want {
				t.Fatalf("token %d: type mismatch: got=%v want=%v", i, tokType, fx.want)
			}
		})
	}
}
