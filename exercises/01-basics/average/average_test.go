package average

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		a, b int
		want float64
	}{
		{a: 1, b: 2, want: 1.5},
		{a: 2, b: 2, want: 2},
		{a: -1, b: 2, want: 0.5},
		{a: 0, b: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Average(tt.a, tt.b); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Average(%d, %d) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
