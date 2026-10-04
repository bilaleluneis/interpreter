package eval_test

import (
	"fmt"
	"testing"

	"goimpl/ast"
	"goimpl/eval"
	"goimpl/token"
)

func TestValueTypesAndInspection(t *testing.T) {
	tests := []struct {
		name            string
		val             eval.Value
		expectedType    eval.ValueType
		expectedInspect string
		expectedString  string
	}{
		{
			name:            "positive integer",
			val:             eval.Integer(42),
			expectedType:    eval.IntegerType,
			expectedInspect: "42",
			expectedString:  "Integer(42)",
		},
		{
			name:            "negative integer",
			val:             eval.Integer(-99),
			expectedType:    eval.IntegerType,
			expectedInspect: "-99",
			expectedString:  "Integer(-99)",
		},
		{
			name:            "zero integer",
			val:             eval.Integer(0),
			expectedType:    eval.IntegerType,
			expectedInspect: "0",
			expectedString:  "Integer(0)",
		},
		{
			name:            "boolean true",
			val:             eval.True,
			expectedType:    eval.BooleanType,
			expectedInspect: "true",
			expectedString:  "Boolean(true)",
		},
		{
			name:            "boolean false",
			val:             eval.False,
			expectedType:    eval.BooleanType,
			expectedInspect: "false",
			expectedString:  "Boolean(false)",
		},
		{
			name:            "null singleton",
			val:             eval.Nil,
			expectedType:    eval.NullType,
			expectedInspect: "null",
			expectedString:  "null",
		},
		{
			name:            "error value",
			val:             eval.Error("type mismatch: INTEGER + BOOLEAN"),
			expectedType:    eval.ErrorType,
			expectedInspect: "ERROR: type mismatch: INTEGER + BOOLEAN",
			expectedString:  "Error(type mismatch: INTEGER + BOOLEAN)",
		},
		{
			name: "return value wrapping integer",
			val: eval.ReturnValue{
				Value: eval.Integer(10),
			},
			expectedType:    eval.ReturnType,
			expectedInspect: "10",
			expectedString:  "ReturnValue(10)",
		},
		{
			name: "return value wrapping nil",
			val: eval.ReturnValue{
				Value: nil,
			},
			expectedType:    eval.ReturnType,
			expectedInspect: "null",
			expectedString:  "ReturnValue(null)",
		},
		{
			name: "function value",
			val: eval.Function{
				Parameters: []ast.Identifier{
					{Tok: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
					{Tok: token.Token{Type: token.IDENTIFIER, Literal: "y"}, Value: "y"},
				},
				Body: ast.Block{
					Tok: token.Token{Type: token.LBRACE, Literal: "{"},
					Statements: []ast.Statement{
						ast.ExpressionStatement{
							Tok: token.Token{Type: token.IDENTIFIER, Literal: "x"},
							Exprssn: ast.InfixExpression{
								Tok:      token.Token{Type: token.ASTER, Literal: "*"},
								Operator: "*",
								Left:     ast.Identifier{Tok: token.Token{Type: token.IDENTIFIER, Literal: "x"}, Value: "x"},
								Right:    ast.Identifier{Tok: token.Token{Type: token.IDENTIFIER, Literal: "y"}, Value: "y"},
							},
						},
					},
				},
				Env: eval.NewEnvironment(),
			},
			expectedType:    eval.FunctionType,
			expectedInspect: "fn(x, y) {\n{\n(x * y)\n}\n}",
			expectedString:  "Function(x, y)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.val.Type() != tt.expectedType {
				t.Fatalf("expected type %q, got %q", tt.expectedType, tt.val.Type())
			}

			if tt.val.Inspect() != tt.expectedInspect {
				t.Errorf("expected inspect %q, got %q", tt.expectedInspect, tt.val.Inspect())
			}

			stringer, ok := tt.val.(fmt.Stringer)
			if !ok {
				t.Fatalf("expected value %T to implement fmt.Stringer", tt.val)
			}

			if stringer.String() != tt.expectedString {
				t.Errorf("expected String() %q, got %q", tt.expectedString, stringer.String())
			}
		})
	}
}

func TestNativeBoolToBoolean(t *testing.T) {
	if got := eval.NativeBoolToBoolean(true); got != eval.True {
		t.Errorf("expected eval.True, got %v", got)
	}

	if got := eval.NativeBoolToBoolean(false); got != eval.False {
		t.Errorf("expected eval.False, got %v", got)
	}
}

func TestErrorImplementsErrorInterface(t *testing.T) {
	var err error = eval.Error("division by zero")
	if err.Error() != "division by zero" {
		t.Errorf("expected error message %q, got %q", "division by zero", err.Error())
	}
}

func TestFunctionParametersStringEmpty(t *testing.T) {
	fn := eval.Function{
		Parameters: []ast.Identifier{},
		Body: ast.Block{
			Tok:        token.Token{Type: token.LBRACE, Literal: "{"},
			Statements: []ast.Statement{},
		},
		Env: eval.NewEnvironment(),
	}

	if fn.Type() != eval.FunctionType {
		t.Fatalf("expected FunctionType, got %q", fn.Type())
	}

	if fn.String() != "Function()" {
		t.Errorf("expected 'Function()', got %q", fn.String())
	}
}
