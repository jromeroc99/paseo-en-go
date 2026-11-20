package primes

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: -3, want: false},
		{n: 0, want: false},
		{n: 1, want: false},
		{n: 4, want: false},
		{n: 25, want: false},
		{n: 2, want: true},
		{n: 3, want: true},
		{n: 17, want: true},
		{n: 97, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime(%d) = %t, want %t", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		limit int
		want  int
	}{
		{limit: 1, want: 0},
		{limit: 2, want: 1},
		{limit: 10, want: 4},
		{limit: 100, want: 25},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CountPrimes(tt.limit); got != tt.want {
				t.Errorf("CountPrimes(%d) = %d, want %d", tt.limit, got, tt.want)
			}
		})
	}
}
