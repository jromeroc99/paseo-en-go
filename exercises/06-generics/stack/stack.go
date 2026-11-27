//go:build !solution

// Package stack is a last-in, first-out stack of values of any type.
//
// Write this type and these methods in this file:
//
//	type Stack[T any] struct { ... }
//
//	func (s *Stack[T]) Push(v T)
//	func (s *Stack[T]) Pop() (T, bool)
//	func (s *Stack[T]) Len() int
//
// Push puts v on top of the stack. Pop takes the value on top off the stack
// and returns it with true; on an empty stack it returns the zero value of
// T and false. Len says how many values the stack holds. The zero value of
// Stack is an empty stack, ready to use: var s Stack[int].
//
// Example: after Push(1) and Push(2), Pop returns 2, true, then 1, true,
// and then 0, false.
//
// You decide which fields Stack needs; the tests only use the methods.
package stack
