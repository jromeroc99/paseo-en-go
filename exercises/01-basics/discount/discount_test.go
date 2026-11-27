package discount

import (
	"math"
	"testing"
)

func TestDiscount(t *testing.T) {
	tests := []struct {
		name           string
		price, percent float64
		want           float64
	}{
		{name: "no discount", price: 19.99, percent: 0, want: 19.99},
		{name: "a quarter off", price: 20, percent: 25, want: 15},
		{name: "the decimals stay", price: 10, percent: 15, want: 8.5},
		{name: "half price", price: 3, percent: 50, want: 1.5},
		{name: "free", price: 42, percent: 100, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Discount(tt.price, tt.percent); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Discount(%v, %v) = %v, want %v", tt.price, tt.percent, got, tt.want)
			}
		})
	}
}
