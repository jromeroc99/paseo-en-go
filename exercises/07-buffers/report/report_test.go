package report

import (
	"bytes"
	"errors"
	"testing"
)

func TestWriteReport(t *testing.T) {
	tests := []struct {
		name   string
		scores []Score
		want   string
	}{
		{
			name:   "no scores",
			scores: nil,
			want:   "Informe\nTotal: 0\n",
		},
		{
			name:   "two scores",
			scores: []Score{{"Ana", 31}, {"Luis", 7}},
			want:   "Informe\nAna: 31\nLuis: 7\nTotal: 38\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			if err := WriteReport(&buf, tt.scores); err != nil {
				t.Fatalf("WriteReport(%v) returned error %v, want nil", tt.scores, err)
			}
			got := buf.String()
			if got == "" {
				t.Fatalf("WriteReport(%v) wrote nothing: did you call Flush at the end?", tt.scores)
			}
			if got != tt.want {
				t.Errorf("WriteReport(%v) wrote\n%q\nwant\n%q", tt.scores, got, tt.want)
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

func TestWriteReportWritesOnce(t *testing.T) {
	var w writeCounter
	WriteReport(&w, []Score{{"Ana", 31}, {"Luis", 7}, {"Eva", 12}})
	switch {
	case w.calls == 0:
		t.Errorf("WriteReport never called w.Write: did you call Flush at the end?")
	case w.calls > 1:
		t.Errorf("WriteReport called w.Write %d times, want 1: write into the bufio.Writer, not into w", w.calls)
	}
}

// brokenWriter fails every write, like a full disk.
type brokenWriter struct{}

var errDiskFull = errors.New("disk full")

func (brokenWriter) Write(p []byte) (int, error) {
	return 0, errDiskFull
}

func TestWriteReportReturnsTheError(t *testing.T) {
	err := WriteReport(brokenWriter{}, []Score{{"Ana", 31}})
	if !errors.Is(err, errDiskFull) {
		t.Errorf("with a Writer that always fails, WriteReport returned %v, want %v: return the error of Flush", err, errDiskFull)
	}
}
