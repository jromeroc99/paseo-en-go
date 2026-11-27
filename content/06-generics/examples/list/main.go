package main

import "fmt"

// List is a linked list of values of any type; a nil *List is empty.
type List[T any] struct {
	next *List[T]
	val  T
}

// Push returns a new list with v in front of l.
func (l *List[T]) Push(v T) *List[T] {
	return &List[T]{next: l, val: v}
}

// Values returns every value, from the front to the end.
func (l *List[T]) Values() []T {
	var values []T
	for n := l; n != nil; n = n.next {
		values = append(values, n.val)
	}
	return values
}

func main() {
	var nums *List[int]
	nums = nums.Push(3).Push(2).Push(1)
	fmt.Println(nums.Values())

	var words *List[string]
	words = words.Push("mundo").Push("hola")
	fmt.Println(words.Values())
}
