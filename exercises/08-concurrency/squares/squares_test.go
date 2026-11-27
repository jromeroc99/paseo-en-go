package squares

import (
	"slices"
	"testing"
)

func TestSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "no numbers", nums: []int{}, want: []int{}},
		{name: "one number", nums: []int{4}, want: []int{16}},
		{name: "several numbers keep their order", nums: []int{1, 2, 3, 4, 5}, want: []int{1, 4, 9, 16, 25}},
		{name: "negative numbers", nums: []int{-3, 0, 3}, want: []int{9, 0, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Squares(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Squares(%v) = %v, want %v; does it wait for every goroutine before returning?",
					tt.nums, got, tt.want)
			}
		})
	}
}

func TestSquaresManyNumbers(t *testing.T) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	got := Squares(nums)
	if len(got) != len(nums) {
		t.Fatalf("Squares of 1000 numbers returned %d results, want 1000", len(got))
	}
	for i, square := range got {
		if square != i*i {
			t.Fatalf("Squares: result[%d] = %d, want %d; does it wait for every goroutine?",
				i, square, i*i)
		}
	}
}
