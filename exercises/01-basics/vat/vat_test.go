package vat

import (
	"math"
	"testing"
)

// VAT has to be a constant: if you declare it with var, this line does not
// compile and go test says "VAT (variable of type float64) is not constant".
const _ = VAT

func TestVAT(t *testing.T) {
	if VAT != 0.21 {
		t.Errorf("VAT = %v, want 0.21", VAT)
	}
}

func TestWithVAT(t *testing.T) {
	tests := []struct {
		price float64
		want  float64
	}{
		{price: 0, want: 0},
		{price: 100, want: 121},
		{price: 10, want: 12.1},
		{price: 2.5, want: 3.025},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := WithVAT(tt.price); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("WithVAT(%v) = %v, want %v", tt.price, got, tt.want)
			}
		})
	}
}
