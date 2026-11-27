package runes

import "testing"

func TestCountRunes(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{s: "", want: 0},
		{s: "go", want: 2},
		{s: "año", want: 3},
		{s: "corazón ♥", want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := CountRunes(tt.s); got != tt.want {
				t.Errorf("CountRunes(%q) = %d, want %d (len(%q) is %d bytes, not runes)", tt.s, got, tt.want, tt.s, len(tt.s))
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{s: "", want: ""},
		{s: "go", want: "og"},
		{s: "año", want: "oña"},
		{s: "¡olé!", want: "!élo¡"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
