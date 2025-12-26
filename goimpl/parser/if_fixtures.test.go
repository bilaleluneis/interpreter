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
		expectedErrMsg: internal.ErrEmptyExpression,
	},
	"missing_opening_parenthesis": {
		tokens: []token.Token{
			{Type: token.IF, Literal: "if"},
			{Type: token.TRUE, Literal: "true"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedErrMsg: fmt.Sprintf(internal.ErrExpectedOpenPren, token.TRUE),
	},
	// TODO: need to unifiy statement formmating and spacing rules
	// otherwise tests will be hard to maintain
	// as we have to match exact string representation
	// with spacing and tabs , etc !!!
	"simple_if_statement": {
		tokens: []token.Token{
			{Type: token.IF, Literal: "if"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.TRUE, Literal: "true"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RETURN, Literal: "return"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedIfExpression: normalize(`if(true){
			return x;
			}`),
	},
	"if_with_else": {
		tokens: []token.Token{
			{Type: token.IF, Literal: "if"},
			{Type: token.LPRAN, Literal: "("},
			{Type: token.TRUE, Literal: "true"},
			{Type: token.RPRAN, Literal: ")"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RETURN, Literal: "return"},
			{Type: token.IDENTIFIER, Literal: "x"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.ELSE, Literal: "else"},
			{Type: token.LBRACE, Literal: "{"},
			{Type: token.RETURN, Literal: "return"},
			{Type: token.IDENTIFIER, Literal: "y"},
			{Type: token.SEMICOLON, Literal: ";"},
			{Type: token.RBRACE, Literal: "}"},
			{Type: token.EOF, Literal: ""},
		},
		expectedIfExpression: normalize(`if(true){
				return x;
				} else {
				return y;
				}`),
	},
}
