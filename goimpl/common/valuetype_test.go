package common

import (
	"testing"
)

// Testing the CopyOf function with Type that has only value reciever
func TestValueReciever(t *testing.T) {
	vr := valueReciever{}
	cp := CopyOf(vr)
	if vr.info() == cp.info() {
		t.Fatal("CopyOf should return a new instance")
	}
}

//----------------------------------------------------------------------------

// Testing the CopyOf function with Type that has only pointer reciever
func TestPointerReciever(t *testing.T) {
	pr := &pointerReciever{}
	cp := CopyOf(pr)
	if pr.info() == cp.info() {
		t.Fatal("CopyOf should return a new instance")
	}
}

//----------------------------------------------------------------------------

// Testing the CopyOf function with Type that has mixed V/P reciever
func TestMixedReciever(t *testing.T) {
	mr := mixedReciever{}
	cp := CopyOf(&mr)
	if mr.info() == cp.info() {
		t.Fatal("CopyOf should return a new instance")
	}
}

//----------------------------------------------------------------------------
