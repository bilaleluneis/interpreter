package common

import "testing"

func TestPtrInstanceOf(t *testing.T) {
	pt := PtrInstanceOf[int]()
	if pt == nil {
		t.Fatal("PtrInstanceOf should return a non-nil pointer")
	}
	if *pt != 0 {
		t.Fatal("PtrInstanceOf should return a pointer to zero value")
	}
}

func TestPtrOf(t *testing.T) {
	i := 42
	pt := PtrOf(i)
	if *pt != i {
		t.Fatal("PtrOf should return a pointer to the value")
	}
}
