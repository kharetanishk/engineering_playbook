// Generics (Go 1.18+): ek hi function alag alag types ke liye.
package main

import "fmt"

// Number constraint: sirf ye types allowed
type Number interface{ ~int | ~int64 | ~float64 }

func Sum[T Number](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}

func Map[T, U any](xs []T, f func(T) U) []U {
	out := make([]U, 0, len(xs)) // capacity pehle se de di, realloc nahi hoga
	for _, x := range xs {
		out = append(out, f(x))
	}
	return out
}

// Stack[T]: generic data structure
type Stack[T any] struct{ items []T }

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}), Sum([]float64{1.5, 2.5}))
	fmt.Println(Map([]int{1, 2, 3}, func(i int) string { return fmt.Sprint("n", i) }))

	s := Stack[string]{}
	s.Push("a")
	s.Push("b")
	v, _ := s.Pop()
	fmt.Println(v) // b
}
