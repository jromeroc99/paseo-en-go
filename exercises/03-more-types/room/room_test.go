package room

import "testing"

func TestRoom(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	tests := []struct {
		name string
		s    []int
		want int
	}{
		{name: "nil slice", s: nil, want: 0},
		{name: "full slice", s: []int{1, 2, 3}, want: 0},
		{name: "made with room", s: make([]int, 2, 5), want: 3},
		{name: "cut from the middle of an array", s: array[1:3], want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Room(tt.s); got != tt.want {
				t.Errorf("Room(s) with len %d and cap %d = %d, want %d",
					len(tt.s), cap(tt.s), got, tt.want)
			}
		})
	}
}

func TestExtend(t *testing.T) {
	tests := []struct {
		n       int
		wantLen int
	}{
		{n: 0, wantLen: 2},
		{n: 2, wantLen: 4},
		{n: 3, wantLen: 5},
		{n: 4, wantLen: 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			s := make([]int, 2, 5)
			got := Extend(s, tt.n)
			if len(got) != tt.wantLen {
				t.Errorf("Extend(s, %d) with len 2 and cap 5 has len %d, want %d",
					tt.n, len(got), tt.wantLen)
			}
			if cap(got) != 5 {
				t.Errorf("Extend(s, %d) has cap %d, want 5: reslice s, do not make a new slice",
					tt.n, cap(got))
			}
		})
	}
}
