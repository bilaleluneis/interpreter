package parser

import (
	"fmt"
	"goimpl/parser/internal"
	"goimpl/token"
)

type letTestCase struct {
	tokens            []token.Token
	expectedErrMsg    string // Non-empty if an error is expected
	expectedStatement string // Non-empty if a valid statement is expected
}

var letStatementTests = map[string]letTestCase{
	"simple_let": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedStatement: "let x = 5;",
	},
	"missing_identifier": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.LetErrExpectedIdentifier, token.ASSIGN),
	},
	"missing_assign": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.LetErrExpectedAssign, token.INT),
	},
	"missing_semicolon": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.INT, Literal: "5"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.LetErrExpectedSemicolon, token.EOF),
	},
	"missing_value_expression": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.LetErrExpectedExpression, token.SEMICOLON),
	},
	"boolean_literal_true": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.TRUE, Literal: "true"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedStatement: "let x = true;",
	},
	"boolean_literal_false": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.FALSE, Literal: "false"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedStatement: "let x = false;",
	},
	"prfix_expression_not_five": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.BANG, Literal: "!"},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedStatement: "let x = (!5);",
	},
	"prefix_expression_not_y": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.BANG, Literal: "!"},
			{Type: token.TRUE, Literal: "y"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedStatement: "let x = (!y);",
	},
	"prefix_expression_plus_five": {
		tokens: []token.Token{
			{Type: token.LET, Literal: "let"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.ASSIGN, Literal: "="},
			{Type: token.PLUS, Literal: "+"},
			{Type: token.INT, Literal: "5"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.EOF, Literal: ""},
		},
		expectedStatement: "let x = (+5);",
	},
}
