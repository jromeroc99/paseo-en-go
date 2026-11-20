package hello

import "testing"

func TestGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Gopher", want: "Hello, Gopher!"},
		{name: "Ana", want: "Hello, Ana!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greeting(tt.name); got != tt.want {
				t.Errorf("Greeting(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
