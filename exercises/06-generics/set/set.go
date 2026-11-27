//go:build !solution

// Package set is a generic set, plus Map and Filter for slices of any type.
//
// Write this type and these functions in this file:
//
//	type Set[T comparable] struct { ... }
//
//	func NewSet[T comparable]() *Set[T]
//	func (s *Set[T]) Add(v T)
//	func (s *Set[T]) Has(v T) bool
//	func (s *Set[T]) Len() int
//
//	func Map[T, U any](s []T, f func(T) U) []U
//	func Filter[T any](s []T, keep func(T) bool) []T
//
// NewSet returns an empty set. Add puts v in the set; adding a value that is
// already there changes nothing. Has says whether v is in the set, and Len
// how many different values it holds.
//
// Map returns a new slice with f applied to every element of s, in order.
// Filter returns a new slice with only the elements of s for which keep
// returns true, in order. Neither of them changes s.
//
// Example: Map([]int{1, 2}, strconv.Itoa) is ["1" "2"], and
// Filter([]int{1, 2, 3, 4}, isEven) is [2 4].
//
// You decide which fields Set needs; the tests only use the methods.
package set
