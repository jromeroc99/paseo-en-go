package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{n: 3, want: "Fizz"},
		{n: 5, want: "Buzz"},
		{n: 15, want: "FizzBuzz"},
		{n: 7, want: "7"},
		{n: 30, want: "FizzBuzz"},
		{n: 1, want: "1"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FizzBuzz(tt.n); got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q, want %q", tt.n, got, tt.want)
			}
		})
	}
}
