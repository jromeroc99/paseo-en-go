package season

import "testing"

func TestSeason(t *testing.T) {
	tests := []struct {
		month int
		want  string
	}{
		{month: 12, want: "winter"},
		{month: 1, want: "winter"},
		{month: 2, want: "winter"},
		{month: 3, want: "spring"},
		{month: 4, want: "spring"},
		{month: 5, want: "spring"},
		{month: 6, want: "summer"},
		{month: 7, want: "summer"},
		{month: 8, want: "summer"},
		{month: 9, want: "autumn"},
		{month: 10, want: "autumn"},
		{month: 11, want: "autumn"},
		{month: 0, want: "unknown"},
		{month: 13, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Season(tt.month); got != tt.want {
				t.Errorf("Season(%d) = %q, want %q", tt.month, got, tt.want)
			}
		})
	}
}
