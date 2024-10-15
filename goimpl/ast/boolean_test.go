package ast

import (
	"goimpl/token"
	"testing"
)

func TestBoolean(t *testing.T) {
	tests := []struct {
		name            string
		booleanAst      Boolean
		expectedLiteral string
		expectedString  string
		expectedDump    string
	}{
		{
			name: "true boolean",
			booleanAst: Boolean{
				Tok:   token.Token{Type: token.TRUE, Literal: "true"},
				Value: true,
			},
			expectedLiteral: "true",
			expectedString:  "true",
			expectedDump:    `ast.Boolean{Tok: true, Value: true}`,
		},
		{
			name: "false boolean",
			booleanAst: Boolean{
				Tok:   token.Token{Type: token.FALSE, Literal: "false"},
				Value: false,
			},
			expectedLiteral: "false",
			expectedString:  "false",
			expectedDump:    `ast.Boolean{Tok: false, Value: false}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actualLiteral := tt.booleanAst.TokenLiteral()
			expectedLiteral := tt.expectedLiteral
			if actualLiteral != expectedLiteral {
				t.Errorf("boolean.TokenLiteral() wrong. got=%q, want=%q", actualLiteral, expectedLiteral)
			}

			actualString := tt.booleanAst.String()
			expectedString := tt.expectedString
			if actualString != expectedString {
				t.Errorf("boolean.String() wrong. got=%q, want=%q", actualString, expectedString)
			}

			actualAst := TrimDump(tt.booleanAst)
			expectedAst := TrimString(tt.expectedDump)
			if actualAst != expectedAst {
				t.Errorf("boolean.Dump() wrong. got=%q, want=%q", actualAst, expectedAst)
			}
		})
	}
}
