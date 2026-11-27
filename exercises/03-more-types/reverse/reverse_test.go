package reverse

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		a    [5]int
		want [5]int
	}{
		{a: [5]int{1, 2, 3, 4, 5}, want: [5]int{5, 4, 3, 2, 1}},
		{a: [5]int{7, 7, 7, 7, 7}, want: [5]int{7, 7, 7, 7, 7}},
		{a: [5]int{0, 0, 9, 0, 0}, want: [5]int{0, 0, 9, 0, 0}},
		{a: [5]int{-1, 0, 1, 2, 3}, want: [5]int{3, 2, 1, 0, -1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Reverse(tt.a); got != tt.want {
				t.Errorf("Reverse(%v) = %v, want %v", tt.a, got, tt.want)
			}
		})
	}
}

func TestReverseLeavesTheOriginalAlone(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	Reverse(a)
	if a != [5]int{1, 2, 3, 4, 5} {
		t.Errorf("after Reverse(a), a = %v, want it unchanged: [1 2 3 4 5]", a)
	}
}
