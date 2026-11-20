package temperature

import (
	"math"
	"testing"
)

func TestFahrenheit(t *testing.T) {
	tests := []struct {
		c    Celsius
		want float64
	}{
		{c: 0, want: 32},
		{c: 100, want: 212},
		{c: -40, want: -40},
		{c: 37, want: 98.6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.c.Fahrenheit(); math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("Celsius(%v).Fahrenheit() = %v, want %v", float64(tt.c), got, tt.want)
			}
		})
	}
}

func TestIsFreezing(t *testing.T) {
	tests := []struct {
		c    Celsius
		want bool
	}{
		{c: 0, want: true},
		{c: -5, want: true},
		{c: 0.5, want: false},
		{c: 20, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := tt.c.IsFreezing(); got != tt.want {
				t.Errorf("Celsius(%v).IsFreezing() = %v, want %v", float64(tt.c), got, tt.want)
			}
		})
	}
}
