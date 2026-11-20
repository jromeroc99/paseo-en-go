package rot13

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

func TestRead(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "Lbh penpxrq gur pbqr!", want: "You cracked the code!"},
		{in: "abc xyz", want: "nop klm"},
		{in: "ABC XYZ", want: "NOP KLM"},
		{in: "123, 456?", want: "123, 456?"},
		{in: "", want: ""},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := io.ReadAll(Rot13Reader{R: strings.NewReader(tt.in)})
			if string(got) != tt.want || err != nil {
				t.Errorf("reading Rot13Reader over %q = (%q, %v), want (%q, <nil>)", tt.in, got, err, tt.want)
			}
		})
	}
}

func TestReadOneByteAtATime(t *testing.T) {
	in := "Lbh penpxrq gur pbqr!"
	want := "You cracked the code!"
	r := Rot13Reader{R: iotest.OneByteReader(strings.NewReader(in))}
	got, err := io.ReadAll(r)
	if string(got) != want || err != nil {
		t.Errorf("reading Rot13Reader over %q one byte at a time = (%q, %v), want (%q, <nil>)", in, got, err, want)
	}
}

func TestReadPassesEOF(t *testing.T) {
	r := Rot13Reader{R: strings.NewReader("")}
	n, err := r.Read(make([]byte, 8))
	if n != 0 || err != io.EOF {
		t.Errorf("Read on an empty source = (%d, %v), want (0, EOF)", n, err)
	}
}
