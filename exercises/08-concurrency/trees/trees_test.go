package trees

import (
	"slices"
	"testing"
	"time"
)

// maxWait is how long a test waits for a channel or for Same.
const maxWait = 300 * time.Millisecond

// build returns a binary search tree with the values inserted in the order
// given, so the same values in another order give another shape.
func build(values ...int) *Tree {
	var root *Tree
	for _, v := range values {
		root = insert(root, v)
	}
	return root
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{Value: v}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func TestWalk(t *testing.T) {
	tests := []struct {
		name string
		tree *Tree
		want []int
	}{
		{name: "the empty tree", tree: nil, want: []int{}},
		{name: "a single node", tree: &Tree{Value: 7}, want: []int{7}},
		{name: "a search tree comes out sorted", tree: build(5, 3, 8, 1, 4, 9), want: []int{1, 3, 4, 5, 8, 9}},
		{
			name: "left, then the node, then right",
			tree: &Tree{Left: &Tree{Value: 9}, Value: 1, Right: &Tree{Value: 5}},
			want: []int{9, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := make(chan int)
			go Walk(tt.tree, ch)
			got := receiveAll(t, ch)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Walk sent %v, want %v", got, tt.want)
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
			t.Fatalf("after receiving %v the channel is still open; Walk must close it when it finishes", got)
		}
	}
}

func TestSame(t *testing.T) {
	tests := []struct {
		name   string
		t1, t2 *Tree
		want   bool
	}{
		{name: "two empty trees", t1: nil, t2: nil, want: true},
		{name: "the same values with different shapes", t1: build(1, 2, 3), t2: build(2, 1, 3), want: true},
		{name: "ten values in two orders", t1: build(5, 2, 8, 1, 3, 9, 7, 4, 6, 10), t2: build(1, 2, 3, 4, 5, 6, 7, 8, 9, 10), want: true},
		{name: "one different value", t1: build(1, 2, 3), t2: build(1, 2, 4), want: false},
		{name: "the second tree has one more value", t1: build(1, 2), t2: build(1, 2, 3), want: false},
		{name: "the first tree has one more value", t1: build(1, 2, 3), t2: build(1, 2), want: false},
		{name: "an empty tree and a full one", t1: nil, t2: build(1), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := make(chan bool)
			go func() { done <- Same(tt.t1, tt.t2) }()
			select {
			case got := <-done:
				if got != tt.want {
					t.Errorf("Same = %t, want %t", got, tt.want)
				}
			case <-time.After(maxWait):
				t.Fatalf("Same does not return; is it waiting for a value that never comes?")
			}
		})
	}
}
