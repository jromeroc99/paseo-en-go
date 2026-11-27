//go:build solution

// Package stack is a last-in, first-out stack of values of any type.
package stack

// Stack holds values of type T; the last one pushed is the first one out.
// Its zero value is an empty stack.
type Stack[T any] struct {
	values []T
}

// Push puts v on top of the stack.
func (s *Stack[T]) Push(v T) {
	s.values = append(s.values, v)
}

// Pop takes the top value off the stack. On an empty stack it returns the
// zero value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, false
	}
	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return top, true
}

// Len returns how many values the stack holds.
func (s *Stack[T]) Len() int {
	return len(s.values)
}
