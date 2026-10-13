# Go Language Features Reference

## Generic Methods

Go version introduced: Go 1.27.

```go
package main

type List[E any] []E
func (List[E]) Map[R any](f func(E) R) List[R] { /* ... */ }
```

## F-Bound Polymorphism

Go version introduced: Go 1.26.

```go
package main

import "fmt"

type Adder[A Adder[A]] interface {
	Add(A) A
}

type Total int

func (t Total) Add(other Total) Total {
	return t + other
}

func addTwo[A Adder[A]](left, right A) A {
	return left.Add(right)
}

func main() {
	fmt.Println(addTwo(Total(4), Total(7)))
}
```

## Range-over-func Iterators

Go version introduced: Go 1.23.

```go
package main

import "fmt"

func count(n int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	for i := range count(3) {
		fmt.Println(i)
	}
}
```