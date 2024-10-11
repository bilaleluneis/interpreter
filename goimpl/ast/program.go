package ast

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

// top will return the last statement from the program, wont remove it
func (p Program) top() string {
	return p.Statements[len(p.Statements)-1].String()
}
