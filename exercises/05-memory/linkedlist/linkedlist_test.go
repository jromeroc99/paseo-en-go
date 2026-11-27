package linkedlist

import (
	"slices"
	"testing"
)

func TestPushAndValues(t *testing.T) {
	tests := []struct {
		name   string
		pushes []int
		want   []int
	}{
		{name: "empty list", pushes: nil, want: []int{}},
		{name: "one node", pushes: []int{7}, want: []int{7}},
		{name: "the last push is the head", pushes: []int{3, 2, 1}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *Node
			for _, v := range tt.pushes {
				head = Push(head, v)
			}
			got := Values(head)
			if got == nil || !slices.Equal(got, tt.want) {
				t.Errorf("after pushing %v, Values(head) = %#v, want %v", tt.pushes, got, tt.want)
			}
		})
	}
}

func TestPushLinksTheOldHead(t *testing.T) {
	first := Push(nil, 1)
	second := Push(first, 2)
	if second.Next != first {
		t.Errorf("Push(first, 2).Next should point to first, the old head")
	}
	if first.Next != nil {
		t.Errorf("Push(nil, 1).Next = %v, want nil: it is the last node", first.Next)
	}
}
