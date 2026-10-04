package eval_test

import (
	"testing"

	"goimpl/eval"
)

func TestEnvironmentGetAndSet(t *testing.T) {
	env := eval.NewEnvironment()

	// Initial lookup should fail
	val, ok := env.Get("x")
	if ok || val != nil {
		t.Fatalf("expected variable 'x' to be undefined, got %v (ok=%t)", val, ok)
	}

	// Set variable and verify lookup
	expected := eval.Integer(10)
	setResult := env.Set("x", expected)
	if setResult != expected {
		t.Fatalf("expected Set() to return %v, got %v", expected, setResult)
	}

	val, ok = env.Get("x")
	if !ok {
		t.Fatalf("expected variable 'x' to be defined")
	}

	intVal, ok := val.(eval.Integer)
	if !ok {
		t.Fatalf("expected val to be eval.Integer, got %T", val)
	}

	if intVal != expected {
		t.Errorf("expected %d, got %d", expected, intVal)
	}
}

func TestEnvironmentOverwrite(t *testing.T) {
	env := eval.NewEnvironment()

	env.Set("x", eval.Integer(10))
	env.Set("x", eval.Integer(20))

	val, ok := env.Get("x")
	if !ok {
		t.Fatalf("expected variable 'x' to be defined")
	}

	intVal, ok := val.(eval.Integer)
	if !ok {
		t.Fatalf("expected val to be eval.Integer, got %T", val)
	}

	if intVal != eval.Integer(20) {
		t.Errorf("expected 20, got %d", intVal)
	}
}

func TestEnvironmentLexicalScoping(t *testing.T) {
	outer := eval.NewEnvironment()
	outer.Set("a", eval.Integer(1))
	outer.Set("b", eval.Integer(2))

	inner := eval.NewEnclosedEnvironment(outer)
	inner.Set("b", eval.Integer(99)) // shadow "b"
	inner.Set("c", eval.Integer(3))

	// inner should find "a" from outer
	valA, ok := inner.Get("a")
	if !ok {
		t.Fatalf("expected inner to find 'a' in outer environment")
	}
	if valA != eval.Integer(1) {
		t.Errorf("expected inner 'a' to be 1, got %v", valA)
	}

	// inner should find its own shadowed "b"
	valB, ok := inner.Get("b")
	if !ok {
		t.Fatalf("expected inner to find local 'b'")
	}
	if valB != eval.Integer(99) {
		t.Errorf("expected inner 'b' to be 99, got %v", valB)
	}

	// outer "b" should remain unchanged
	outerB, ok := outer.Get("b")
	if !ok {
		t.Fatalf("expected outer to have 'b'")
	}
	if outerB != eval.Integer(2) {
		t.Errorf("expected outer 'b' to remain 2, got %v", outerB)
	}

	// inner should find "c"
	valC, ok := inner.Get("c")
	if !ok {
		t.Fatalf("expected inner to find local 'c'")
	}
	if valC != eval.Integer(3) {
		t.Errorf("expected inner 'c' to be 3, got %v", valC)
	}

	// outer should NOT find "c"
	_, ok = outer.Get("c")
	if ok {
		t.Fatalf("outer scope should not have access to inner variable 'c'")
	}
}

func TestEnvironmentMultiLevelNesting(t *testing.T) {
	global := eval.NewEnvironment()
	global.Set("count", eval.Integer(100))

	parent := eval.NewEnclosedEnvironment(global)
	child := eval.NewEnclosedEnvironment(parent)

	// Lookup traversing two levels
	val, ok := child.Get("count")
	if !ok {
		t.Fatalf("expected child to find 'count' in grandparent environment")
	}
	if val != eval.Integer(100) {
		t.Errorf("expected 'count' to be 100, got %v", val)
	}

	// Setting in parent does not affect global
	parent.Set("count", eval.Integer(200))

	valChild, _ := child.Get("count")
	if valChild != eval.Integer(200) {
		t.Errorf("expected child to resolve to closest enclosing scope (200), got %v", valChild)
	}

	valGlobal, _ := global.Get("count")
	if valGlobal != eval.Integer(100) {
		t.Errorf("expected global 'count' to remain 100, got %v", valGlobal)
	}
}
