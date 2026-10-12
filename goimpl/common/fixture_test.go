package common

import (
	"bytes"
	"encoding/gob"
	"errors"
)

type exportedRecord struct {
	Name    string
	Count   int
	Tags    []string
	Weights map[string]int
}

// customSecret is a test type with unexported fields implementing GobSerializable[*customSecret].
type customSecret struct {
	token   string
	tags    []string
	weights map[string]int
}

func (c *customSecret) GobEncode() ([]byte, error) {
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
