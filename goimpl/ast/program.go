package ast

import "strings"

// Program is the top most node that will contain child statments
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	out := "\n---- Program Start ----\n"
	for _, s := range p.Statements {
		out += "\n"
		out += s.String()
		out += "\n"
	}
	out += "\n---- Program End ----\n"
	return out
}

// Top will return the last statement from the program, wont remove it
func (p Program) Top() string {
	return p.Statements[len(p.Statements)-1].String()
}

func (p Program) Dump(ident int) string {
	out := "\n---- Program Dump Start ----\n"
	out += "ast.Program{ //start of Program\n"
	out += strings.Repeat(" ", ident)
	out += "Statements: []ast.Statement{ //start of Statment\n"
	out += func() string {
		if len(p.Statements) == 0 {
			return ""
		}
		out := ""
		for _, s := range p.Statements {
			out += strings.Repeat(" ", ident+1) + s.Dump(ident+1) + ",\n"
		}
		return out
	}()
	out += strings.Repeat(" ", ident) + "} //end of Statment\n"
	out += strings.Repeat(" ", ident-1) + "} //end of Program"
	out += "\n---- Program Dump End ----\n"
	return out
}
