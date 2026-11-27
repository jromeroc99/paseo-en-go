package sum

import "testing"

func TestSumInts(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "several", nums: []int{1, 2, 3}, want: 6},
		{name: "with negatives", nums: []int{10, -4}, want: 6},
		{name: "empty slice", nums: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.nums); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSumFloats(t *testing.T) {
	tests := []struct {
		name string
		nums []float64
		want float64
	}{
		{name: "decimals", nums: []float64{0.5, 1.5}, want: 2},
		{name: "just one", nums: []float64{2.25}, want: 2.25},
		{name: "empty slice", nums: []float64{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.nums); got != tt.want {
				t.Errorf("Sum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// Euros has float64 as its underlying type, so it satisfies ~float64.
type Euros float64

func TestSumOwnType(t *testing.T) {
	prices := []Euros{1.5, 2.5, 3}
	if got := Sum(prices); got != 7 {
		t.Errorf("Sum(%v) = %v, want 7 (does Number have ~ before float64?)", prices, got)
	}
}
