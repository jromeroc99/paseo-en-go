package circle

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	tests := []struct {
		radius        int
		wantArea      float64
		wantPerimeter float64
	}{
		{radius: 0, wantArea: 0, wantPerimeter: 0},
		{radius: 1, wantArea: math.Pi, wantPerimeter: 2 * math.Pi},
		{radius: 2, wantArea: 4 * math.Pi, wantPerimeter: 4 * math.Pi},
		{radius: 10, wantArea: 100 * math.Pi, wantPerimeter: 20 * math.Pi},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			area, perimeter := Circle(tt.radius)
			if math.Abs(area-tt.wantArea) > 1e-9 || math.Abs(perimeter-tt.wantPerimeter) > 1e-9 {
				t.Errorf("Circle(%d) = (%v, %v), want (%v, %v)",
					tt.radius, area, perimeter, tt.wantArea, tt.wantPerimeter)
			}
		})
	}
}
