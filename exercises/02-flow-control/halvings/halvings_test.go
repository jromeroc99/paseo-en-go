package halvings

import "testing"

func TestHalvings(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 0, want: 0},
		{n: 1, want: 1},
		{n: 8, want: 4},
		{n: 1000, want: 10},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Halvings(tt.n); got != tt.want {
				t.Errorf("Halvings(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
