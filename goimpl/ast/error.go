package ast

type Error struct {
	Message string
}

func (Error) statmentNode()          {}
func (e Error) TokenLiteral() string { return e.Message }
func (e Error) String() string {
	if e.Message != "" {
		return "Error{" + e.Message + "}"
	}
	return e.Message
}
