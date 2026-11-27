package sumints

import "testing"

func TestSumInts(t *testing.T) {
	tests := []struct {
		values []any
		want   int
	}{
		{values: nil, want: 0},
		{values: []any{1, 2, 3}, want: 6},
		{values: []any{1, "2", 3.0, true, 4}, want: 5},
		{values: []any{"hola", 2.5, nil}, want: 0},
		{values: []any{-3, 10}, want: 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SumInts(tt.values); got != tt.want {
				t.Errorf("SumInts(%#v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}
