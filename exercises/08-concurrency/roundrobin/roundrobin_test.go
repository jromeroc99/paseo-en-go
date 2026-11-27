package roundrobin

import (
	"slices"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		name    string
		tasks   []string
		threads int
		want    [][]string
	}{
		{
			name:    "a single thread runs everything",
			tasks:   []string{"a", "b", "c"},
			threads: 1,
			want:    [][]string{{"a", "b", "c"}},
		},
		{
			name:    "two threads take turns",
			tasks:   []string{"a", "b", "c", "d", "e"},
			threads: 2,
			want:    [][]string{{"a", "c", "e"}, {"b", "d"}},
		},
		{
			name:    "as many threads as tasks",
			tasks:   []string{"a", "b", "c"},
			threads: 3,
			want:    [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name:    "more threads than tasks",
			tasks:   []string{"a", "b"},
			threads: 4,
			want:    [][]string{{"a"}, {"b"}, {}, {}},
		},
		{
			name:    "no tasks",
			tasks:   nil,
			threads: 2,
			want:    [][]string{{}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RoundRobin(tt.tasks, tt.threads)
			if len(got) != len(tt.want) {
				t.Fatalf("RoundRobin(%v, %d) returned %d lists, want %d (one per thread)",
					tt.tasks, tt.threads, len(got), tt.threads)
			}
			for thread := range tt.want {
				if !slices.Equal(got[thread], tt.want[thread]) {
					t.Errorf("RoundRobin(%v, %d): thread %d got %v, want %v",
						tt.tasks, tt.threads, thread, got[thread], tt.want[thread])
				}
			}
		})
	}
}
