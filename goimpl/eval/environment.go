package eval

// Environment tracks variable bindings and supports lexical scoping via outer environments.
type Environment struct {
	store map[string]Value
	outer *Environment
}

// NewEnvironment creates a new top-level environment.
func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Value),
		outer: nil,
	}
}

// NewEnclosedEnvironment creates a new child environment with outer as its parent scope.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	return &Environment{
		store: make(map[string]Value),
		outer: outer,
	}
}

// Get looks up a variable name, traversing parent environments if not found locally.
func (e *Environment) Get(name string) (Value, bool) {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return val, ok
}

// Set binds a name to a value in the current local environment.
func (e *Environment) Set(name string, val Value) Value {
	e.store[name] = val
	return val
}
