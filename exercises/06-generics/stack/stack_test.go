package stack

import "testing"

func TestPopReturnsTheLastPushed(t *testing.T) {
	tests := []struct {
		name   string
		pushes []int
		want   []int // what each Pop returns, in order
	}{
		{name: "one", pushes: []int{7}, want: []int{7}},
		{name: "last in, first out", pushes: []int{1, 2, 3}, want: []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s Stack[int]
			for _, v := range tt.pushes {
				s.Push(v)
			}
			for _, want := range tt.want {
				got, ok := s.Pop()
				if !ok || got != want {
					t.Fatalf("after pushing %v, Pop() = %d, %t, want %d, true", tt.pushes, got, ok, want)
				}
			}
		})
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	var s Stack[string]
	got, ok := s.Pop()
	if ok || got != "" {
		t.Errorf(`on an empty stack, Pop() = %q, %t, want "", false (the zero value of string)`, got, ok)
	}

	s.Push("hola")
	s.Pop()
	got, ok = s.Pop()
	if ok || got != "" {
		t.Errorf(`after emptying the stack, Pop() = %q, %t, want "", false`, got, ok)
	}
}

func TestLen(t *testing.T) {
	var s Stack[float64]
	if got := s.Len(); got != 0 {
		t.Errorf("on a new stack, Len() = %d, want 0", got)
	}
	s.Push(1.5)
	s.Push(2.5)
	if got := s.Len(); got != 2 {
		t.Errorf("after two Push, Len() = %d, want 2", got)
	}
	s.Pop()
	if got := s.Len(); got != 1 {
		t.Errorf("after two Push and one Pop, Len() = %d, want 1", got)
	}
}
