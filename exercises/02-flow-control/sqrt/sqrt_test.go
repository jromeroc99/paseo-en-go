package sqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []float64{0, 1, 2, 4, 9, 100, 0.25, 2e6}
	for _, x := range tests {
		t.Run("", func(t *testing.T) {
			got, want := Sqrt(x), math.Sqrt(x)
			if math.Abs(got-want) > 1e-9 {
				t.Errorf("Sqrt(%v) = %v, want %v", x, got, want)
			}
		})
	}
}
