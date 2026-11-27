package set

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestSet(t *testing.T) {
	tests := []struct {
		name    string
		adds    []string
		wantLen int
	}{
		{name: "empty", adds: nil, wantLen: 0},
		{name: "no repeats", adds: []string{"rojo", "verde"}, wantLen: 2},
		{name: "a repeat counts once", adds: []string{"rojo", "verde", "rojo"}, wantLen: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet[string]()
			for _, v := range tt.adds {
				s.Add(v)
			}
			if got := s.Len(); got != tt.wantLen {
				t.Errorf("after adding %q, Len() = %d, want %d", tt.adds, got, tt.wantLen)
			}
			for _, v := range tt.adds {
				if !s.Has(v) {
					t.Errorf("after adding %q, Has(%q) = false, want true", tt.adds, v)
				}
			}
			if s.Has("azul") {
				t.Errorf(`after adding %q, Has("azul") = true, want false`, tt.adds)
			}
		})
	}
}

func TestSetOfInts(t *testing.T) {
	s := NewSet[int]()
	s.Add(3)
	if !s.Has(3) || s.Has(4) {
		t.Errorf("Set[int] after Add(3): Has(3) = %t, Has(4) = %t, want true, false", s.Has(3), s.Has(4))
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []string
	}{
		{name: "several", s: []int{1, 2, 3}, want: []string{"1", "2", "3"}},
		{name: "empty slice", s: []int{}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.s, strconv.Itoa)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Map(%v, strconv.Itoa) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestMapChangesType(t *testing.T) {
	words := []string{"go", "genéricos"}
	got := Map(words, strings.ToUpper)
	want := []string{"GO", "GENÉRICOS"}
	if !slices.Equal(got, want) {
		t.Errorf("Map(%q, strings.ToUpper) = %q, want %q", words, got, want)
	}
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{name: "keeps the even ones", s: []int{1, 2, 3, 4}, want: []int{2, 4}},
		{name: "none passes", s: []int{1, 3}, want: []int{}},
		{name: "empty slice", s: []int{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := slices.Clone(tt.s)
			got := Filter(tt.s, isEven)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Filter(%v, isEven) = %v, want %v", tt.s, got, tt.want)
			}
			if !slices.Equal(tt.s, original) {
				t.Errorf("Filter changed its input: it is now %v, want %v", tt.s, original)
			}
		})
	}
}
