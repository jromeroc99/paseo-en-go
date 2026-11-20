package describe

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		v    any
		want string
	}{
		{v: 42, want: "int: 42"},
		{v: "hi", want: "string: hi"},
		{v: true, want: "bool: true"},
		{v: nil, want: "nil"},
		{v: 3.5, want: "unknown"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Describe(tt.v); got != tt.want {
				t.Errorf("Describe(%#v) = %q, want %q", tt.v, got, tt.want)
			}
		})
	}
}
