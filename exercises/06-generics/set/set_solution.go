//go:build solution

// Package set is a generic set, plus Map and Filter for slices of any type.
package set

// Set holds each value of type T at most once.
type Set[T comparable] struct {
	items map[T]bool
}

// NewSet returns an empty set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: map[T]bool{}}
}

// Add puts v in the set; adding it twice keeps one copy.
func (s *Set[T]) Add(v T) {
	s.items[v] = true
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	return s.items[v]
}

// Len returns how many different values the set holds.
func (s *Set[T]) Len() int {
	return len(s.items)
}

// Map returns f applied to every element of s, in the same order.
func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

// Filter returns the elements of s for which keep is true, in order.
func Filter[T any](s []T, keep func(T) bool) []T {
	var result []T
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}
