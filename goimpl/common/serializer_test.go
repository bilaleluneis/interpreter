package common

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// customSecret is a test type with unexported fields implementing GobSerializable[*customSecret].
type customSecret struct {
	token   string
	tags    []string
	weights map[string]int
}

func (c customSecret) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	shadow := struct {
		Token   string
		Tags    []string
		Weights map[string]int
	}{
		Token:   c.token,
		Tags:    c.tags,
		Weights: c.weights,
	}
	if err := enc.Encode(shadow); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *customSecret) GobDecode(data []byte) error {
	var shadow struct {
		Token   string
		Tags    []string
		Weights map[string]int
	}
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&shadow); err != nil {
		return err
	}
	c.token = shadow.Token
	c.tags = shadow.Tags
	c.weights = shadow.Weights
	return nil
}

// failingSerializable always fails during GobEncode to verify error handling.
type failingSerializable struct {
	val int
}

func (f *failingSerializable) GobEncode() ([]byte, error) {
	return nil, errors.New("simulated encode error")
}

func (f *failingSerializable) GobDecode(data []byte) error {
	return nil
}

func TestDeepCopySuccessAndIsolation(t *testing.T) {
	// Slices and maps isolation
	origSlice := []string{"alpha", "beta"}
	origMap := map[string]int{"k1": 10}

	type Container struct {
		List []string
		Dict map[string]int
	}

	c := Container{List: origSlice, Dict: origMap}
	copied, err := DeepCopy(c)
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
	copiedCh, err := DeepCopy(ch)
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
	copiedFn, err := DeepCopy(fn)
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
	secret := &customSecret{
		token:   "s3cr3t",
		tags:    []string{"admin", "prod"},
		weights: map[string]int{"cpu": 100},
	}

	// Statically verified to satisfy GobSerializable[*customSecret]
	cloned := Serialize[*customSecret](secret)

	if cloned == secret {
		t.Fatal("Serialize should return a distinct pointer")
	}
	if cloned.token != "s3cr3t" {
		t.Fatalf("expected token 's3cr3t', got %s", cloned.token)
	}

	// Mutate original slice & map
	secret.tags[0] = "guest"
	secret.weights["cpu"] = 0
	secret.token = "modified"

	if cloned.token != "s3cr3t" {
		t.Fatalf("expected cloned token to remain 's3cr3t', got %s", cloned.token)
	}
	if cloned.tags[0] != "admin" {
		t.Fatalf("expected cloned tag to remain 'admin', got %s", cloned.tags[0])
	}
	if cloned.weights["cpu"] != 100 {
		t.Fatalf("expected cloned weight to remain 100, got %d", cloned.weights["cpu"])
	}
}

func TestSerializePanicOnFailure(t *testing.T) {
	failing := &failingSerializable{val: 42}

	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected Serialize to panic when encoding fails")
		}
		errMsg := fmt.Sprint(r)
		if !strings.Contains(errMsg, "Serialize failed") {
			t.Fatalf("expected panic message to contain 'Serialize failed', got %v", r)
		}
	}()

	Serialize[*failingSerializable](failing)
}
