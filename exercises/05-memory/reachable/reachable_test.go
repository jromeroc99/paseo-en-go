package reachable

import (
	"slices"
	"testing"
)

func TestReachable(t *testing.T) {
	tests := []struct {
		name  string
		refs  map[string][]string
		roots []string
		want  []string
	}{
		{
			name:  "no roots, everything is garbage",
			refs:  map[string][]string{"a": {"b"}},
			roots: nil,
			want:  []string{},
		},
		{
			name:  "a root with no pointers",
			refs:  map[string][]string{},
			roots: []string{"a"},
			want:  []string{"a"},
		},
		{
			name:  "follow a chain",
			refs:  map[string][]string{"a": {"b"}, "b": {"c"}},
			roots: []string{"a"},
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "nothing points to c",
			refs:  map[string][]string{"a": {"b"}, "c": {"a"}},
			roots: []string{"a"},
			want:  []string{"a", "b"},
		},
		{
			name:  "a cycle is followed once",
			refs:  map[string][]string{"a": {"b"}, "b": {"a"}},
			roots: []string{"a"},
			want:  []string{"a", "b"},
		},
		{
			name:  "a cycle nobody reaches is garbage",
			refs:  map[string][]string{"x": {"y"}, "y": {"x"}, "r": {}},
			roots: []string{"r"},
			want:  []string{"r"},
		},
		{
			name:  "two roots sharing an object, sorted and without repeats",
			refs:  map[string][]string{"z": {"m"}, "b": {"m"}},
			roots: []string{"z", "b"},
			want:  []string{"b", "m", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reachable(tt.refs, tt.roots)
			if got == nil || !slices.Equal(got, tt.want) {
				t.Errorf("Reachable(%v, %v) = %#v, want %v", tt.refs, tt.roots, got, tt.want)
			}
		})
	}
}
