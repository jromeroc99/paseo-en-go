package player

import "testing"

func TestNewPlayer(t *testing.T) {
	tests := []string{"Ana", "Luis", ""}
	for _, name := range tests {
		t.Run("", func(t *testing.T) {
			p := NewPlayer(name)
			if p == nil {
				t.Fatalf("NewPlayer(%q) = nil, want a pointer to a Player", name)
			}
			if p.Name != name || p.Score != 0 {
				t.Errorf("NewPlayer(%q) = %+v, want {Name:%s Score:0}", name, *p, name)
			}
		})
	}
}

func TestAddPoints(t *testing.T) {
	tests := []struct {
		score, points int
		want          int
	}{
		{score: 0, points: 10, want: 10},
		{score: 10, points: 5, want: 15},
		{score: 7, points: 0, want: 7},
		{score: 3, points: -2, want: 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			p := &Player{Name: "Ana", Score: tt.score}
			AddPoints(p, tt.points)
			if p.Score != tt.want {
				t.Errorf("AddPoints(player with Score %d, %d): Score = %d, want %d",
					tt.score, tt.points, p.Score, tt.want)
			}
			if p.Name != "Ana" {
				t.Errorf("AddPoints changed the name to %q, want %q", p.Name, "Ana")
			}
		})
	}
}
