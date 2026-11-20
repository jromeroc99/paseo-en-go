package evens

import (
	"slices"
	"testing"
)

func TestEvens(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nums: nil, want: []int{}},
		{nums: []int{1, 2, 3, 4}, want: []int{2, 4}},
		{nums: []int{1, 3}, want: []int{}},
		{nums: []int{-2, 0, 5}, want: []int{-2, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			original := slices.Clone(tt.nums)
			got := Evens(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Evens(%v) = %v, want %v", original, got, tt.want)
			}
			if !slices.Equal(tt.nums, original) {
				t.Errorf("Evens(%v) changed its input to %v; build a new slice instead", original, tt.nums)
			}
		})
	}
}
