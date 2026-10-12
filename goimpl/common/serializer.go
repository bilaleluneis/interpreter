package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

// GobSerializable is an interface constraint that requires implementing both
// gob.GobEncoder and gob.GobDecoder.
// Custom types with non-exported (private) fields can implement these methods
// to participate in safe deep-copying and serialization via encoding/gob.
type GobSerializable[T any] interface {
	*T
	gob.GobEncoder
	gob.GobDecoder
}

// deepCopy creates an independent, deep copy of val using encoding/gob.
// It preserves value semantics and avoids accidental shared state across pointers,
// slices, or maps.
// Unlike shallow copies, the returned instance shares no reference-bearing memory
// with the input.
// If val cannot be encoded or decoded by gob (for example if val contains channels,
// functions, or unregistered interface types), deepCopy does NOT return val; instead
// it returns the zero value of T and a non-nil error detailing the exact encoding
// or decoding failure reason.
func deepCopy[T any](val T) (T, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	encodedValue := any(val)
	rawValue := reflect.ValueOf(val)
	if rawValue.IsValid() && rawValue.Kind() != reflect.Ptr {
		addressableCopy := reflect.New(rawValue.Type())
		addressableCopy.Elem().Set(rawValue)
		encodedValue = addressableCopy.Interface()
	}
	if err := enc.Encode(encodedValue); err != nil {
		var zero T
		return zero, fmt.Errorf("failed to encode: %w", err)
	}
	var copied T
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(&copied); err != nil {
		var zero T
		return zero, fmt.Errorf("failed to decode: %w", err)
	}
	return copied, nil
}

// Serialize creates an independent deep copy of a value whose type implements GobSerializable[T].
// It is designed specifically for custom types with non-exported fields that define custom
// GobEncode and GobDecode logic to control their binary representation.
// If encoding or decoding fails, Serialize returns an Err Option with the failure reason.
// usage example:
//
//	type MyType struct {
//	    field int
//	}
//	func (m *MyType) GobEncode() ([]byte, error) { ... }
//	func (m *MyType) GobDecode(data []byte) error { ... }
//	val := MyType{field: 42}
//	copiedOption := Serialize[MyType](val)
func Serialize[T any, _ GobSerializable[T]](val T) Option[T] {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&val); err != nil {
		return Err[T](fmt.Errorf("failed to encode: %w", err))
	}

	var copied T
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(&copied); err != nil {
		return Err[T](fmt.Errorf("failed to decode: %w", err))
	}

	return Option[T]{value: copied, hasValue: true}
}
