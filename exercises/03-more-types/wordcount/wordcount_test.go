package wordcount

import (
	"maps"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		s    string
		want map[string]int
	}{
		{s: "", want: map[string]int{}},
		{s: "I am learning Go", want: map[string]int{"I": 1, "am": 1, "learning": 1, "Go": 1}},
		{s: "the quick the lazy the", want: map[string]int{"the": 3, "quick": 1, "lazy": 1}},
		{s: "  hola   mundo  hola ", want: map[string]int{"hola": 2, "mundo": 1}},
		{s: "Go go", want: map[string]int{"Go": 1, "go": 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := WordCount(tt.s)
			if !maps.Equal(got, tt.want) {
				t.Errorf("WordCount(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
