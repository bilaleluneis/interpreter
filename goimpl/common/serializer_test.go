package common

import (
	"strings"
	"testing"
)

func TestDeepCopySuccessAndIsolation(t *testing.T) {
	// Slices and maps isolation
	origSlice := []string{"alpha", "beta"}
	origMap := map[string]int{"k1": 10}

	type Container struct {
		List []string
		Dict map[string]int
	}

	c := Container{List: origSlice, Dict: origMap}
	copied, err := deepCopy(c)
	if err != nil {
		t.Fatalf("expected DeepCopy to succeed, got %v", err)
	}

	// Mutate original
	origSlice[0] = "mutated"
	origMap["k1"] = 999

	if copied.List[0] != "alpha" {
		t.Fatalf("expected copied slice to remain 'alpha', got %s", copied.List[0])
	}
	if copied.Dict["k1"] != 10 {
		t.Fatalf("expected copied map to remain 10, got %d", copied.Dict["k1"])
	}

	// Mutate copied
	copied.List[0] = "tampered"
	copied.Dict["k1"] = 777

	if c.List[0] != "mutated" {
		t.Fatalf("expected original list to be unaffected by mutating copy")
	}
}

func TestDeepCopyFailureOnUnsupportedTypes(t *testing.T) {
	// Channels cannot be encoded by gob
	ch := make(chan int, 1)
	copiedCh, err := deepCopy(ch)
	if err == nil {
		t.Fatal("expected DeepCopy to fail on channel")
	}
	if !strings.Contains(err.Error(), "failed to encode") {
		t.Fatalf("expected error to specify 'failed to encode', got: %v", err)
	}
	if copiedCh != nil {
		t.Fatalf("expected zero value returned on error, got %v", copiedCh)
	}

	// Functions cannot be encoded by gob
	fn := func() string { return "test" }
	copiedFn, err := deepCopy(fn)
	if err == nil {
		t.Fatal("expected DeepCopy to fail on function")
	}
	if !strings.Contains(err.Error(), "failed to encode") {
		t.Fatalf("expected error to specify 'failed to encode', got: %v", err)
	}
	if copiedFn != nil {
		t.Fatal("expected nil function returned on error")
	}
}

func TestSerializeWithCustomUnexportedFields(t *testing.T) {
	secret := customSecret{
		token:   "s3cr3t",
		tags:    []string{"admin", "prod"},
		weights: map[string]int{"cpu": 100},
	}

	// Statically verified to satisfy GobSerializable[*customSecret]
	cloned := Serialize(secret)
	if !cloned.IsSome() {
		t.Fatalf("expected Serialize to return Some, got %v", cloned.Error())
	}
	clonedValue := cloned.Unwrap()

	if clonedValue.token != "s3cr3t" ||
		len(clonedValue.tags) != 2 ||
		clonedValue.weights["cpu"] != 100 {
		t.Fatalf("Serialize returned unexpected cloned value: %+v", clonedValue)
	}
	if clonedValue.token != "s3cr3t" {
		t.Fatalf("expected token 's3cr3t', got %s", clonedValue.token)
	}

	// Mutate original slice & map
	secret.tags[0] = "guest"
	secret.weights["cpu"] = 0
	secret.token = "modified"

	if clonedValue.token != "s3cr3t" {
		t.Fatalf("expected cloned token to remain 's3cr3t', got %s", clonedValue.token)
	}
	if clonedValue.tags[0] != "admin" {
		t.Fatalf("expected cloned tag to remain 'admin', got %s", clonedValue.tags[0])
	}
	if clonedValue.weights["cpu"] != 100 {
		t.Fatalf("expected cloned weight to remain 100, got %d", clonedValue.weights["cpu"])
	}
}

func TestSerializeFailureReturnsNone(t *testing.T) {
	failing := &failingSerializable{val: 42}

	serialized := Serialize(*failing)
	if !serialized.IsNone() {
		t.Fatalf("expected Serialize to return None on encode failure, got %v", serialized)
	}
	if serialized.Error() == nil {
		t.Fatal("expected Serialize to return an error on failure")
	}
	if !strings.Contains(serialized.Error().Error(), "failed to encode") {
		t.Fatalf("expected encode failure message, got %v", serialized.Error())
	}
}
