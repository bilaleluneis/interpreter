package common

import "fmt"

// valueReciever is a test type that has only value reciever
type valueReciever struct {
	counter int
	list    []int
}

func (v valueReciever) info() string {
	return fmt.Sprintf("valueReciever address: %p", &v)
}

func (v valueReciever) GetCopy() valueReciever {
	return valueReciever{}
}

// pointerReciever is a test type that has only pointer reciever
type pointerReciever struct {
	counter int
	list    []int
}

func (v *pointerReciever) info() string {
	return fmt.Sprintf("pointerReciever address: %p", &v)
}

func (v *pointerReciever) GetCopy() *pointerReciever {
	return &pointerReciever{}
}

// mixedReciever is a test type that has mixed value and pointer reciever
type mixedReciever struct {
	counter int
	list    []int
}

func (v mixedReciever) info() string {
	return fmt.Sprintf("mixedReciever address: %p", &v)
}
func (v *mixedReciever) GetCopy() *mixedReciever {
	return &mixedReciever{}
}
