package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 3, Height: 4}, want: 12},
		{shape: Rectangle{Width: 0, Height: 5}, want: 0},
		{shape: Circle{Radius: 1}, want: math.Pi},
		{shape: Circle{Radius: 2}, want: 4 * math.Pi},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.shape.Area(); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("%#v.Area() = %v, want %v", tt.shape, got, tt.want)
			}
		})
	}
}

func TestTotalArea(t *testing.T) {
	tests := []struct {
		shapes []Shape
		want   float64
	}{
		{shapes: nil, want: 0},
		{shapes: []Shape{Rectangle{Width: 2, Height: 5}}, want: 10},
		{shapes: []Shape{Rectangle{Width: 3, Height: 4}, Circle{Radius: 1}}, want: 12 + math.Pi},
		{shapes: []Shape{Circle{Radius: 1}, Circle{Radius: 1}, Rectangle{Width: 1, Height: 1}}, want: 2*math.Pi + 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TotalArea(tt.shapes); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("TotalArea(%v) = %v, want %v", tt.shapes, got, tt.want)
			}
		})
	}
}
