package minimal

import (
	"goimpl/lexer"
	"goimpl/parser/pratt"
	"testing"
)

// Test parsing:
// let five = 5;
// let ten = 10;
// let mul = fn(x, y) { x * y };
// let add = fn(x, y) { x + y };
// let result = if(five < ten) { mul(five, ten) } else { add(five, ten };
func TestParsing(t *testing.T) {
	input := `	let five = 5;
				let ten = 10;
				let mul = fn(x, y) { x * y };
				let add = fn(x, y) { x + y };
				let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
			`

	lexr := lexer.NewLazyLexer(input)
	p := pratt.New(&lexr)
	prog := p.ParseProgram()

	if len(prog.Statements) != 5 {
		t.Fatalf("program has wrong number of statements. got=%d", len(prog.Statements))
	}

	expected := []string{
		"let five = 5;",
		"let ten = 10;",
		// see ast.Fun String() for formatting
		"let mul = fn(x, y) {\n(x * y)\n};",
		// see ast.Fun String() for formatting
		"let add = fn(x, y) {\n(x + y)\n};",
		// see ast.IfExpression String() for formatting
		"let result = if((five < ten)) {\nmul(five, ten)\n} else {\nadd(five, ten)\n};",
	}

	for i, stmt := range prog.Statements {
		t.Run(expected[i], func(t *testing.T) {
			actual := stmt.String()
			expect := expected[i]
			if actual != expect {
				t.Errorf("statement %d wrong. got=%q, want=%q", i+1, actual, expect)
			}
		})
	}
}
