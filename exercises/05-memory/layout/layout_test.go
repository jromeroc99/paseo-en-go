package layout

import "testing"

func TestSize(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		want   int
	}{
		{name: "no fields", fields: nil, want: 0},
		{name: "one byte", fields: []int{1}, want: 1},
		{name: "bool, int64, bool", fields: []int{1, 8, 1}, want: 24},
		{name: "int64, bool, bool", fields: []int{8, 1, 1}, want: 16},
		{name: "two bytes then an int32", fields: []int{1, 1, 4}, want: 8},
		{name: "int16 after a byte", fields: []int{1, 2}, want: 4},
		{name: "padding only at the end", fields: []int{4, 1}, want: 8},
		{name: "already a multiple", fields: []int{4, 4}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Size(tt.fields); got != tt.want {
				t.Errorf("Size(%v) = %d, want %d", tt.fields, got, tt.want)
			}
		})
	}
}
