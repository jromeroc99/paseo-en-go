package generate

import (
	"slices"
	"testing"
	"time"
)

// maxWait is how long a test waits for the channel to be closed.
const maxWait = 300 * time.Millisecond

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{name: "nothing to send", n: 0, want: []int{}},
		{name: "only the first number", n: 1, want: []int{0}},
		{name: "the first two numbers", n: 2, want: []int{0, 1}},
		{name: "ten numbers", n: 10, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int)
			go Fibonacci(tt.n, ch)
			got := receiveAll(t, ch)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Fibonacci(%d) sent %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// receiveAll reads ch until it is closed, and fails the test if it is
// still open after maxWait.
func receiveAll(t *testing.T, ch chan int) []int {
	t.Helper()
	got := []int{}
	timeout := time.After(maxWait)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return got
			}
			got = append(got, v)
		case <-timeout:
			t.Fatalf("after receiving %v the channel is still open; close it with close(ch) when you finish", got)
		}
	}
}
