package sum

import "testing"

func TestSumTo(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 10, want: 55},
		{n: 100, want: 5050},
		{n: -5, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SumTo(tt.n); got != tt.want {
				t.Errorf("SumTo(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
