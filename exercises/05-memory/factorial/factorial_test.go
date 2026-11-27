package factorial

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{n: 0, want: 1},
		{n: 1, want: 1},
		{n: 2, want: 2},
		{n: 3, want: 6},
		{n: 5, want: 120},
		{n: 10, want: 3628800},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Factorial(tt.n); got != tt.want {
				t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
