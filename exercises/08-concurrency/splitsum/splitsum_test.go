package splitsum

import (
	"testing"
	"time"
)

// maxWait is how long a test waits for a channel before it gives up.
const maxWait = 300 * time.Millisecond

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "no numbers add up to zero", nums: []int{}, want: 0},
		{name: "one number", nums: []int{5}, want: 5},
		{name: "several numbers", nums: []int{7, 2, 8}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int)
			go Sum(tt.nums, ch)
			select {
			case got := <-ch:
				if got != tt.want {
					t.Errorf("Sum(%v) sent %d, want %d", tt.nums, got, tt.want)
				}
			case <-time.After(maxWait):
				t.Fatalf("Sum(%v) sent nothing on the channel", tt.nums)
			}
		})
	}
}

func TestSplitSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "the tour's example", nums: []int{7, 2, 8, -9, 4, 0}, want: 12},
		{name: "an odd length", nums: []int{1, 2, 3}, want: 6},
		{name: "one number", nums: []int{9}, want: 9},
		{name: "no numbers", nums: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan int)
			go func() { done <- SplitSum(tt.nums) }()
			select {
			case got := <-done:
				if got != tt.want {
					t.Errorf("SplitSum(%v) = %d, want %d", tt.nums, got, tt.want)
				}
			case <-time.After(maxWait):
				t.Fatalf("SplitSum(%v) does not return; does it receive both totals from the channel?", tt.nums)
			}
		})
	}
}
