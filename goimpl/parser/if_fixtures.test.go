package parser

import (
	"fmt"
	"goimpl/parser/internal"
	"goimpl/token"
)

type ifTestCase struct {
	tokens               []token.Token
	expectedErrMsg       string // Non-empty if an error is expected
	expectedIfExpression string // Expected string representation of the parsed IfExpression
}

var ifTests = map[string]ifTestCase{
	"simple_if": {
		tokens: []token.Token{
			{Type: token.IF, Literal: "if"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.TRUE, Literal: "true"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedIfExpression: normalize(`if (true) { x; }`),
	},
	"empty_expression_in_condition": {
		tokens: []token.Token{
			{Type: token.IF, Literal: "if"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.ErrExpectedExpression, token.RPRAN),
	},
	"missing_opening_parenthesis": {
		tokens: []token.Token{
			{Type: token.IF, Literal: "if"},
			{Type: token.TRUE, Literal: "true"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.ErrExpectedOpenPren, token.TRUE),
	},
}
