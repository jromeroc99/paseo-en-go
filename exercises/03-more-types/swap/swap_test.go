package swap

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		a, b int
	}{
		{a: 1, b: 2},
		{a: -3, b: 7},
		{a: 0, b: 9},
		{a: 5, b: 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			a, b := tt.a, tt.b
			Swap(&a, &b)
			if a != tt.b || b != tt.a {
				t.Errorf("after Swap(&a, &b) with a=%d, b=%d: a=%d, b=%d, want a=%d, b=%d",
					tt.a, tt.b, a, b, tt.b, tt.a)
			}
		})
	}
}
