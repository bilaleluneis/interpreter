# Go Best Practices Reference

## Favor Iterators and Range-over-Func

When a solution naturally walks a sequence, prefer the iterator pattern and `range` over functions instead of exposing manual index loops or callback-style traversal in the API.

Primitive loop:

```go
package main

import "fmt"

func printEvenNumbers(values []int) {
	for i := 0; i < len(values); i++ {
		if values[i]%2 == 0 {
			fmt.Println(values[i])
		}
	}
}
```

Iterator and range-over-func solution:

```go
package main

import "fmt"

func evenNumbers(values []int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for _, value := range values {
			if value%2 != 0 {
				continue
			}
			if !yield(value) {
				return
			}
		}
	}
}

func main() {
	for value := range evenNumbers([]int{1, 2, 3, 4, 5}) {
		fmt.Println(value)
	}
}
```

Use this pattern eagerly when suggesting new solutions, especially when it keeps traversal lazy, composable, and easier to read.