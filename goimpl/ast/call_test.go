package ast

import "testing"

func TestCall(t *testing.T) {
	tests := []struct {
		name string
		call Call
		want string
	}{
		{
			name: "function no args",
			call: Call{
				Function: Identifier{Value: "doSomething"},
				Args:     []Expression{},
			},
			want: "doSomething()",
		},
		{
			name: "function with args",
			call: Call{
				Function: Identifier{Value: "add"},
				Args: []Expression{
					IntegerLiteral{Value: 1},
					IntegerLiteral{Value: 2},
					IntegerLiteral{Value: 3},
				},
			},
			want: "add(1, 2, 3)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.call.String()
			if got != tc.want {
				t.Fatalf("Call.String() = %q, want %q", got, tc.want)
			}
		})
	}
}
