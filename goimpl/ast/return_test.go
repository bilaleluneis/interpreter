package ast

import (
	"goimpl/token"
	"testing"
)

func TestReturn(t *testing.T) {
	tests := []struct {
		input          string
		returnAst      Return
		expectedString string
		expectedDump   string
	}{
		{
			input: "return 42;",
			returnAst: Return{
				Tok: token.Token{Type: token.RETURN, Literal: "return"},
				Value: &IntegerLiteral{
					Tok:   token.Token{Type: token.INT, Literal: "42"},
					Value: 42,
				},
			},
			expectedString: "return 42;",
			expectedDump: `ast.Return{
							Tok: token.Token{Type:token.RETURN, Literal: return},
							Value: ast.IntegerLiteral{
								Tok: token.Token{Type:token.INT, Literal: 42},
								Value: 42,
								},
							}`,
		},
		{
			input: "return;",
			returnAst: Return{
				Tok:   token.Token{Type: token.RETURN, Literal: "return"},
				Value: nil,
			},
			expectedString: "return",
			expectedDump:   `ast.Return{Tok: token.Token{Type:token.RETURN, Literal:"return"}, Value: nil,}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			actualString := tt.returnAst.String()
			expectedString := tt.expectedString
			if actualString != expectedString {
				t.Errorf("return.String() wrong. got=%q, want=%q", actualString, expectedString)
			}

			actualAst := TrimDump(tt.returnAst)
			expecteAst := TrimString(tt.expectedDump)
			if actualAst != expecteAst {
				t.Errorf("return.Dump() wrong. got=%q, want=%q", actualAst, expecteAst)
			}
		})
	}
}
