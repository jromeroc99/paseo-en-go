package fibonacci

import "testing"

func TestFibonacciFirstTen(t *testing.T) {
	f := Fibonacci()
	if f == nil {
		t.Fatal("Fibonacci() = nil, want a function")
	}
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, w := range want {
		t.Run("", func(t *testing.T) {
			if got := f(); got != w {
				t.Errorf("call %d to the function from Fibonacci() = %d, want %d", i+1, got, w)
			}
		})
	}
}

func TestFibonacciGeneratorsAreIndependent(t *testing.T) {
	first, second := Fibonacci(), Fibonacci()
	if first == nil || second == nil {
		t.Fatal("Fibonacci() = nil, want a function")
	}
	for range 5 {
		first()
	}
	tests := []struct {
		name string
		next func() int
		want int
	}{
		{name: "second generator, call 1", next: second, want: 0},
		{name: "first generator, call 6", next: first, want: 5},
		{name: "second generator, call 2", next: second, want: 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.next(); got != tt.want {
				t.Errorf("%s = %d, want %d; each call to Fibonacci() needs its own variables",
					tt.name, got, tt.want)
			}
		})
	}
}
