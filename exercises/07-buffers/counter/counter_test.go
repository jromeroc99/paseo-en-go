package counter

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestCopyThroughCountingWriter(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{name: "empty", text: ""},
		{name: "short", text: "hola"},
		{name: "long text", text: strings.Repeat("abc", 20000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			cw := &CountingWriter{W: &buf}
			n, err := io.Copy(cw, strings.NewReader(tt.text))
			if err != nil {
				t.Fatalf("io.Copy into a CountingWriter returned error %v, want nil", err)
			}
			if buf.String() != tt.text {
				t.Errorf("W received %d bytes, want %d: pass p on to c.W", buf.Len(), len(tt.text))
			}
			if cw.Count != len(tt.text) || int(n) != len(tt.text) {
				t.Errorf("Count = %d and io.Copy returned %d, want %d for both", cw.Count, n, len(tt.text))
			}
		})
	}
}

func TestCountAddsUp(t *testing.T) {
	var buf bytes.Buffer
	cw := &CountingWriter{W: &buf}
	cw.Write([]byte("uno "))
	cw.Write([]byte("dos"))
	if cw.Count != 7 {
		t.Errorf("after writing \"uno \" and \"dos\", Count = %d, want 7: add to Count, do not replace it", cw.Count)
	}
}

// shortWriter accepts two bytes and then fails, like a full disk.
type shortWriter struct{}

var errDiskFull = errors.New("disk full")

func (shortWriter) Write(p []byte) (int, error) {
	return min(len(p), 2), errDiskFull
}

func TestCountOnlyWhatWasWritten(t *testing.T) {
	cw := &CountingWriter{W: shortWriter{}}
	n, err := cw.Write([]byte("hola"))
	if n != 2 || !errors.Is(err, errDiskFull) {
		t.Errorf("with a W that writes only 2 bytes, Write = (%d, %v), want (2, %v): return what W returns", n, err, errDiskFull)
	}
	if cw.Count != 2 {
		t.Errorf("with a W that writes only 2 bytes, Count = %d, want 2: count the n W returns, not len(p)", cw.Count)
	}
}
