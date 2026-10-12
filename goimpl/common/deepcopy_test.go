package common

import (
	"reflect"
	"testing"
)

func TestDeepCopyValueAndPointerInputs(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{
			name: "exported struct value",
			run: func(t *testing.T) {
				original := exportedRecord{
					Name:    "alpha",
					Count:   7,
					Tags:    []string{"one", "two"},
					Weights: map[string]int{"cpu": 100},
				}

				copied, err := deepCopy(original)
				if err != nil {
					t.Fatalf("expected deepCopy to succeed for exported value, got %v", err)
				}

				if !reflect.DeepEqual(copied, original) {
					t.Fatalf("copy not matching original, got %#v and %#v", copied, original)
				}

				original.Name = "mutated"
				original.Tags[0] = "changed"
				original.Weights["cpu"] = 0

				if copied.Name != "alpha" {
					t.Fatalf("expected copied name to remain alpha, got %s", copied.Name)
				}
				if copied.Tags[0] != "one" {
					t.Fatalf("expected copied tag to remain one, got %s", copied.Tags[0])
				}
				if copied.Weights["cpu"] != 100 {
					t.Fatalf("expected copied weight to remain 100, got %d", copied.Weights["cpu"])
				}
			},
		},
		{
			name: "exported struct pointer",
			run: func(t *testing.T) {
				original := &exportedRecord{
					Name:    "beta",
					Count:   11,
					Tags:    []string{"north", "south"},
					Weights: map[string]int{"mem": 2048},
				}

				copied, err := deepCopy(original)
				if err != nil {
					t.Fatalf("expected deepCopy to succeed for exported pointer, got %v", err)
				}
				if copied == nil {
					t.Fatal("expected copied pointer to be non-nil")
				}
				if copied == original {
					t.Fatal("expected deepCopy to return a distinct pointer")
				}
				if !reflect.DeepEqual(*copied, *original) {
					t.Fatalf("copied ptr not matching orig, got %#v and %#v", copied, original)
				}

				original.Name = "mutated"
				original.Tags[0] = "changed"
				original.Weights["mem"] = 4096

				if copied.Name != "beta" {
					t.Fatalf("expected copied name to remain beta, got %s", copied.Name)
				}
				if copied.Tags[0] != "north" {
					t.Fatalf("expected copied tag to remain north, got %s", copied.Tags[0])
				}
				if copied.Weights["mem"] != 2048 {
					t.Fatalf("expected copied weight to remain 2048, got %d", copied.Weights["mem"])
				}
			},
		},
		{
			name: "custom gob value",
			run: func(t *testing.T) {
				original := customSecret{
					token:   "s3cr3t",
					tags:    []string{"admin", "prod"},
					weights: map[string]int{"cpu": 100},
				}

				copied, err := deepCopy(original)
				if err != nil {
					t.Fatalf("expected deepCopy to succeed for custom gob value, got %v", err)
				}

				if !reflect.DeepEqual(copied, original) {
					t.Fatalf("copied value not matching orig, got %#v and %#v", copied, original)
				}

				original.token = "mutated"
				original.tags[0] = "guest"
				original.weights["cpu"] = 0

				if copied.token != "s3cr3t" {
					t.Fatalf("expected copied token to remain s3cr3t, got %s", copied.token)
				}
				if copied.tags[0] != "admin" {
					t.Fatalf("expected copied tag to remain admin, got %s", copied.tags[0])
				}
				if copied.weights["cpu"] != 100 {
					t.Fatalf("expected copied weight to remain 100, got %d", copied.weights["cpu"])
				}
			},
		},
		{
			name: "custom gob pointer",
			run: func(t *testing.T) {
				original := &customSecret{
					token:   "opaque",
					tags:    []string{"private", "stable"},
					weights: map[string]int{"gpu": 2},
				}

				copied, err := deepCopy(original)
				if err != nil {
					t.Fatalf("expected deepCopy to succeed for custom gob pointer, got %v", err)
				}
				if copied == nil {
					t.Fatal("expected copied pointer to be non-nil")
				}
				if copied == original {
					t.Fatal("expected deepCopy to return a distinct pointer")
				}
				if !reflect.DeepEqual(*copied, *original) {
					t.Fatalf("copied pointer not matching orig, got %#v and %#v", copied, original)
				}

				original.token = "changed"
				original.tags[0] = "mutated"
				original.weights["gpu"] = 8

				if copied.token != "opaque" {
					t.Fatalf("expected copied token to remain opaque, got %s", copied.token)
				}
				if copied.tags[0] != "private" {
					t.Fatalf("expected copied tag to remain private, got %s", copied.tags[0])
				}
				if copied.weights["gpu"] != 2 {
					t.Fatalf("expected copied weight to remain 2, got %d", copied.weights["gpu"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}
