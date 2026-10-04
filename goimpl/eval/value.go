package eval

import (
	"fmt"
	"goimpl/ast"
	"strings"
)

// ValueType identifies the runtime type of a Value.
type ValueType string

const (
	IntegerType  ValueType = "INTEGER"
	BooleanType  ValueType = "BOOLEAN"
	NullType     ValueType = "NULL"
	ReturnType   ValueType = "RETURN"
	ErrorType    ValueType = "ERROR"
	FunctionType ValueType = "FUNCTION"
)

// Value represents an evaluated runtime value in Monkey.
type Value interface {
	Type() ValueType
	Inspect() string
}

// Integer wraps int64 as an immutable runtime value.
type Integer int64

func (i Integer) Type() ValueType { return IntegerType }
func (i Integer) Inspect() string { return fmt.Sprintf("%d", i) }
func (i Integer) String() string  { return fmt.Sprintf("Integer(%d)", i) }

// Boolean wraps bool as an immutable runtime value.
type Boolean bool

func (b Boolean) Type() ValueType { return BooleanType }
func (b Boolean) Inspect() string { return fmt.Sprintf("%t", b) }
func (b Boolean) String() string  { return fmt.Sprintf("Boolean(%t)", b) }

// Error represents a runtime evaluation error.
type Error string

func (e Error) Type() ValueType { return ErrorType }
func (e Error) Inspect() string { return "ERROR: " + string(e) }
func (e Error) String() string  { return fmt.Sprintf("Error(%s)", string(e)) }

// Error implements the built-in error interface.
func (e Error) Error() string { return string(e) }

// Null represents the absence of a value.
type Null struct{}

func (n Null) Type() ValueType { return NullType }
func (n Null) Inspect() string { return "null" }
func (n Null) String() string  { return "null" }

// Preallocated singletons and constants for common immutable values.
const (
	True  Boolean = true
	False Boolean = false
)

var (
	Nil = Null{}
)

// NativeBoolToBoolean maps a Go bool to a Boolean singleton.
func NativeBoolToBoolean(input bool) Boolean {
	if input {
		return True
	}
	return False
}

// ReturnValue wraps an inner Value when returning from a block or function.
type ReturnValue struct {
	Value Value
}

func (r ReturnValue) Type() ValueType { return ReturnType }
func (r ReturnValue) Inspect() string {
	if r.Value == nil {
		return "null"
	}
	return r.Value.Inspect()
}
func (r ReturnValue) String() string {
	if r.Value == nil {
		return "ReturnValue(null)"
	}
	return fmt.Sprintf("ReturnValue(%s)", r.Value.Inspect())
}

// Function represents a user-defined function literal with its captured environment.
type Function struct {
	Parameters []ast.Identifier
	Body       ast.Block
	Env        *Environment
}

func (f Function) Type() ValueType { return FunctionType }
func (f Function) Inspect() string {
	params := make([]string, 0, len(f.Parameters))
	for _, p := range f.Parameters {
		params = append(params, p.Value)
	}
	return fmt.Sprintf("fn(%s) {\n%s\n}", strings.Join(params, ", "), f.Body.String())
}
func (f Function) String() string {
	params := make([]string, 0, len(f.Parameters))
	for _, p := range f.Parameters {
		params = append(params, p.Value)
	}
	return fmt.Sprintf("Function(%s)", strings.Join(params, ", "))
}
