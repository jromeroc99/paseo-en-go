package ring

import (
	"slices"
	"testing"
)

func TestValues(t *testing.T) {
	tests := []struct {
		name string
		size int
		puts []int
		want []int
	}{
		{name: "empty ring", size: 3, puts: nil, want: []int{}},
		{name: "not full yet", size: 3, puts: []int{1, 2}, want: []int{1, 2}},
		{name: "exactly full", size: 3, puts: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "the newest replaces the oldest", size: 3, puts: []int{1, 2, 3, 4}, want: []int{2, 3, 4}},
		{name: "wraps around twice", size: 3, puts: []int{1, 2, 3, 4, 5, 6, 7}, want: []int{5, 6, 7}},
		{name: "room for one", size: 1, puts: []int{1, 2, 3}, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRing(tt.size)
			for _, v := range tt.puts {
				r.Put(v)
			}
			if got := r.Values(); !slices.Equal(got, tt.want) {
				t.Errorf("NewRing(%d) after Put of %v: Values() = %v, want %v (oldest first)",
					tt.size, tt.puts, got, tt.want)
			}
		})
	}
}

func TestValuesIsACopy(t *testing.T) {
	r := NewRing(2)
	r.Put(1)
	r.Put(2)
	r.Values()[0] = 99
	if got := r.Values(); !slices.Equal(got, []int{1, 2}) {
		t.Errorf("after changing the slice Values returned, Values() = %v, want [1 2]: return a copy, not the ring's own slice", got)
	}
}
