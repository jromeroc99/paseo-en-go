package inner

import (
	"slices"
	"testing"
)

func TestInner(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{nums: []int{1, 2, 3, 4, 5}, want: []int{2, 3, 4}},
		{nums: []int{10, 20, 30}, want: []int{20}},
		{nums: []int{10, 20}, want: []int{}},
		{nums: []int{10}, want: []int{}},
		{nums: []int{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Inner(tt.nums); !slices.Equal(got, tt.want) {
				t.Errorf("Inner(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestInnerSharesTheArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got := Inner(nums)
	got[0] = 99
	if nums[1] != 99 {
		t.Errorf("after Inner(nums)[0] = 99, nums = %v, want nums[1] to be 99 too: "+
			"Inner should slice nums, not copy it", nums)
	}
}
