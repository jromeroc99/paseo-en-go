package rectangle

import "testing"

func TestArea(t *testing.T) {
	tests := []struct {
		width, height int
		want          int
	}{
		{width: 3, height: 4, want: 12},
		{width: 5, height: 5, want: 25},
		{width: 0, height: 7, want: 0},
		{width: 1, height: 1, want: 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Area(tt.width, tt.height); got != tt.want {
				t.Errorf("Area(%d, %d) = %d, want %d", tt.width, tt.height, got, tt.want)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		width, height int
		want          int
	}{
		{width: 3, height: 4, want: 14},
		{width: 5, height: 5, want: 20},
		{width: 0, height: 7, want: 14},
		{width: 1, height: 1, want: 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Perimeter(tt.width, tt.height); got != tt.want {
				t.Errorf("Perimeter(%d, %d) = %d, want %d", tt.width, tt.height, got, tt.want)
			}
		})
	}
}
