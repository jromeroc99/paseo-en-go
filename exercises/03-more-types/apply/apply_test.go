package apply

import (
	"slices"
	"testing"
)

func double(n int) int { return n * 2 }

func square(n int) int { return n * n }

func TestApply(t *testing.T) {
	tests := []struct {
		nums  []int
		fName string
		f     func(int) int
		want  []int
	}{
		{nums: []int{1, 2, 3}, fName: "double", f: double, want: []int{2, 4, 6}},
		{nums: []int{-2, 3, 0}, fName: "square", f: square, want: []int{4, 9, 0}},
		{nums: []int{}, fName: "double", f: double, want: []int{}},
		{nums: nil, fName: "square", f: square, want: []int{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			original := slices.Clone(tt.nums)
			got := Apply(tt.nums, tt.f)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Apply(%v, %s) = %v, want %v", original, tt.fName, got, tt.want)
			}
			if !slices.Equal(tt.nums, original) {
				t.Errorf("Apply(%v, %s) changed its input to %v; build a new slice instead",
					original, tt.fName, tt.nums)
			}
		})
	}
}
