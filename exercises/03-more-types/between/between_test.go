package between

import (
	"slices"
	"testing"
)

func TestBetween(t *testing.T) {
	tests := []struct {
		from, to int
		want     []int
	}{
		{from: 1, to: 5, want: []int{1, 2, 3, 4, 5}},
		{from: 3, to: 3, want: []int{3}},
		{from: -2, to: 1, want: []int{-2, -1, 0, 1}},
		{from: 5, to: 1, want: []int{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Between(tt.from, tt.to)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Between(%d, %d) = %v, want %v", tt.from, tt.to, got, tt.want)
			}
		})
	}
}
