package parser

import (
	"goimpl/token"
)

type expressionTestCase struct {
	tokens             []token.Token
	expectedErrMsg     string // Non-empty if an error is expected
	expectedExpression string // Expected string representation of the parsed expression
}

var expressionTests = map[string]expressionTestCase{
	"identifier": {
		tokens: []token.Token{
			{Type: token.IDENTIFIER, Literal: "foobar"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "foobar",
	},
	"integer_literal": {
		tokens: []token.Token{
			{Type: token.INT, Literal: "5"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "5",
	},
	"boolean_true": {
		tokens: []token.Token{
			{Type: token.TRUE, Literal: "true"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "true",
	},
	"boolean_false": {
		tokens: []token.Token{
			{Type: token.FALSE, Literal: "false"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "false",
	},
	"prefix_bang": {
		tokens: []token.Token{
			{Type: token.BANG, Literal: "!"},
			{Type: token.INT, Literal: "5"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "(!5)",
	},
	"prefix_minus": {
		tokens: []token.Token{
			{Type: token.MINUS, Literal: "-"},
			{Type: token.INT, Literal: "15"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "(-15)",
	},
	"infix_add": {
		tokens: []token.Token{
			{Type: token.INT, Literal: "5"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.INT, Literal: "10"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "(5 + 10)",
	},
	"infix_subtract": {
		tokens: []token.Token{
			{Type: token.INT, Literal: "7"},
			{Type: token.MINUS, Literal: "-"},
			{Type: token.INT, Literal: "3"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "(7 - 3)",
	},
	"infix_multiply": {
		tokens: []token.Token{
			{Type: token.INT, Literal: "2"},
			{Type: token.ASTER, Literal: "*"},
			{Type: token.INT, Literal: "8"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "(2 * 8)",
	},
	"infix_divide": {
		tokens: []token.Token{
			{Type: token.INT, Literal: "10"},
			{Type: token.SLASH, Literal: "/"},
			{Type: token.INT, Literal: "2"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "(10 / 2)",
	},
	"grouped_expression": {
		tokens: []token.Token{
			{Type: token.INT, Literal: "5"},
			{Type: token.ASTER, Literal: "*"},
			{Type: token.INT, Literal: "5"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.INT, Literal: "10"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "((5 * 5) + 10)",
	},
	"function_literal": { // TODO: move function parsing to its own tests
		tokens: []token.Token{
			{Type: token.FUNCTION, Literal: "fn"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENTIFIER, Literal: "y"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RETURN, Literal: "return"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.IDENTIFIER, Literal: "y"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedExpression: "fn(x, y) {return (x + y);}",
	},
}
