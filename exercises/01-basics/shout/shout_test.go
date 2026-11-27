package shout

import "testing"

func TestShout(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "hola", want: "HOLA!"},
		{s: "Go mola", want: "GO MOLA!"},
		{s: "YA", want: "YA!"},
		{s: "", want: "!"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := Shout(tt.s); got != tt.want {
				t.Errorf("Shout(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
