package countdown

import (
	"io"
	"os"
	"testing"
)

// printed runs f and returns everything it wrote to the standard output.
func printed(t *testing.T, f func()) string {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	stdout := os.Stdout
	os.Stdout = writer
	f()
	os.Stdout = stdout
	writer.Close()
	out, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

func TestCountdown(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{n: 0, want: "Go!\n"},
		{n: 1, want: "1\nGo!\n"},
		{n: 3, want: "3\n2\n1\nGo!\n"},
		{n: 5, want: "5\n4\n3\n2\n1\nGo!\n"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := printed(t, func() { Countdown(tt.n) })
			if got != tt.want {
				t.Errorf("Countdown(%d) prints\n%s\nwant\n%s", tt.n, got, tt.want)
			}
		})
	}
}
