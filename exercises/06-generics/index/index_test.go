package index

import "testing"

func TestIndexInts(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		x    int
		want int
	}{
		{name: "the first one", s: []int{10, 20, 30}, x: 10, want: 0},
		{name: "the last one", s: []int{10, 20, 30}, x: 30, want: 2},
		{name: "not there", s: []int{10, 20, 30}, x: 40, want: -1},
		{name: "repeated: the first time", s: []int{7, 5, 7}, x: 7, want: 0},
		{name: "empty slice", s: []int{}, x: 1, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.s, tt.x); got != tt.want {
				t.Errorf("Index(%v, %d) = %d, want %d", tt.s, tt.x, got, tt.want)
			}
		})
	}
}

func TestIndexStrings(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		x    string
		want int
	}{
		{name: "in the middle", s: []string{"a", "b", "c"}, x: "b", want: 1},
		{name: "not there", s: []string{"a", "b", "c"}, x: "z", want: -1},
		{name: "case matters", s: []string{"Go", "go"}, x: "go", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.s, tt.x); got != tt.want {
				t.Errorf("Index(%q, %q) = %d, want %d", tt.s, tt.x, got, tt.want)
			}
		})
	}
}
