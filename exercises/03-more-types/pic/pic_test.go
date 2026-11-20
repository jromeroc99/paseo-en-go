package pic

import "testing"

func TestPicSize(t *testing.T) {
	tests := []struct {
		dx, dy int
	}{
		{dx: 0, dy: 0},
		{dx: 3, dy: 2},
		{dx: 1, dy: 4},
		{dx: 256, dy: 256},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Pic(tt.dx, tt.dy)
			if len(got) != tt.dy {
				t.Fatalf("len(Pic(%d, %d)) = %d rows, want %d", tt.dx, tt.dy, len(got), tt.dy)
			}
			for y, row := range got {
				if len(row) != tt.dx {
					t.Fatalf("len(Pic(%d, %d)[%d]) = %d columns, want %d", tt.dx, tt.dy, y, len(row), tt.dx)
				}
			}
		})
	}
}

func TestPicValues(t *testing.T) {
	tests := []struct {
		dx, dy int
		want   [][]uint8
	}{
		{dx: 3, dy: 2, want: [][]uint8{{0, 0, 1}, {0, 1, 1}}},
		{dx: 2, dy: 3, want: [][]uint8{{0, 0}, {0, 1}, {1, 1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Pic(tt.dx, tt.dy)
			if !equal(got, tt.want) {
				t.Errorf("Pic(%d, %d) = %v, want %v", tt.dx, tt.dy, got, tt.want)
			}
		})
	}
}

func TestPicLargeCells(t *testing.T) {
	got := Pic(256, 256)
	if len(got) != 256 || len(got[0]) != 256 {
		t.Fatalf("Pic(256, 256) must have 256 rows of 256 columns")
	}
	tests := []struct {
		x, y int
		want uint8
	}{
		{x: 0, y: 0, want: 0},
		{x: 50, y: 100, want: 75},
		{x: 0, y: 255, want: 127},
		{x: 255, y: 255, want: 255},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if cell := got[tt.y][tt.x]; cell != tt.want {
				t.Errorf("Pic(256, 256)[%d][%d] = %d, want %d", tt.y, tt.x, cell, tt.want)
			}
		})
	}
}

func TestPicRowsAreIndependent(t *testing.T) {
	got := Pic(3, 2)
	if len(got) != 2 || len(got[0]) != 3 || len(got[1]) != 3 {
		t.Fatalf("Pic(3, 2) must have 2 rows of 3 columns")
	}
	got[0][0] = 99
	if got[1][0] != 0 {
		t.Errorf("after changing Pic(3, 2)[0][0], [1][0] = %d, want 0; create each row with its own make", got[1][0])
	}
}

// equal reports whether two grids have the same rows and cells.
func equal(a, b [][]uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for y := range a {
		if len(a[y]) != len(b[y]) {
			return false
		}
		for x := range a[y] {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}
