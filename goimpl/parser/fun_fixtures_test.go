package parser

import (
	"fmt"
	"goimpl/parser/internal"
	"goimpl/token"
)

type funTestCase struct {
	tokens                []token.Token
	expectedErrMsg        string // Non-empty if an error is expected
	expectedFunExpression string // Expected string representation of the parsed Function
}

var funTests = map[string]funTestCase{
	"simple_function": {
		tokens: []token.Token{
			{Type: token.FUNCTION, Literal: "fn"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedFunExpression: normalize(`fn() {}`),
	},
	"function_with_parameters": {
		tokens: []token.Token{
			{Type: token.FUNCTION, Literal: "fn"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENTIFIER, Literal: "y"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedFunExpression: normalize(`fn(x, y) {}`),
	},
	"function_with_body": {
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
		expectedFunExpression: normalize(`fn(x, y) {
		return (x + y);
		}`),
	},
	// FIXME: fix test bellow
	"function_missing_rparen": {
		tokens: []token.Token{
			{Type: token.FUNCTION, Literal: "fn"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.COMMA, Literal: ","},
			{Type: token.IDENTIFIER, Literal: "y"},
			// Missing RPRAN here
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.ErrExpectedClosePren, token.LBRACE),
	},
}
