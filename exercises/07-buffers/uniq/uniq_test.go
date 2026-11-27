package uniq

import (
	"bytes"
	"strings"
	"testing"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "empty", in: "", want: ""},
		{name: "nothing repeats", in: "a\nb\nc\n", want: "a\nb\nc\n"},
		{name: "a repeated line", in: "a\na\nb\n", want: "a\nb\n"},
		{name: "only repeats next to each other", in: "a\na\nb\na\n", want: "a\nb\na\n"},
		{name: "empty lines are lines too", in: "\n\nx\n", want: "\nx\n"},
		{name: "a long run", in: "z\nz\nz\nz\n", want: "z\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := Uniq(strings.NewReader(tt.in), &out); err != nil {
				t.Fatalf("Uniq(%q) returned error %v, want nil", tt.in, err)
			}
			if tt.want != "" && out.Len() == 0 {
				t.Fatalf("Uniq(%q) wrote nothing: did you call Flush at the end?", tt.in)
			}
			if got := out.String(); got != tt.want {
				t.Errorf("Uniq(%q) wrote %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// writeCounter counts how many times Write is called on it.
type writeCounter struct {
	calls int
}

func (c *writeCounter) Write(p []byte) (int, error) {
	c.calls++
	return len(p), nil
}

func TestUniqUsesABuffer(t *testing.T) {
	var w writeCounter
	Uniq(strings.NewReader("a\nb\nc\nd\n"), &w)
	switch {
	case w.calls == 0:
		t.Errorf("Uniq never called w.Write: did you call Flush at the end?")
	case w.calls > 1:
		t.Errorf("Uniq called w.Write %d times, want 1: write through a bufio.Writer", w.calls)
	}
}

func TestUniqReportsTooLongLines(t *testing.T) {
	in := strings.Repeat("a", 70000) + "\n"
	var out bytes.Buffer
	if err := Uniq(strings.NewReader(in), &out); err == nil {
		t.Errorf("Uniq with a 70000-byte line returned a nil error, want the error of scanner.Err()")
	}
}
