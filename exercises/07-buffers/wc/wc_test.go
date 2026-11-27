package wc

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		text string
		want Counts
	}{
		{name: "empty text", text: "", want: Counts{}},
		{name: "one line", text: "hola\n", want: Counts{Lines: 1, Words: 1, Bytes: 5}},
		{name: "two lines", text: "hola mundo\nde Go\n", want: Counts{Lines: 2, Words: 4, Bytes: 17}},
		{name: "extra spaces are not words", text: "  uno   dos  \n", want: Counts{Lines: 1, Words: 2, Bytes: 14}},
		{name: "an empty line still counts", text: "uno\n\ndos\n", want: Counts{Lines: 3, Words: 2, Bytes: 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Count(strings.NewReader(tt.text))
			if err != nil {
				t.Fatalf("Count(%q) returned error %v, want nil", tt.text, err)
			}
			if got != tt.want {
				t.Errorf("Count(%q) = %+v, want %+v", tt.text, got, tt.want)
			}
		})
	}
}

func TestCountReportsTooLongLines(t *testing.T) {
	text := strings.Repeat("a", 70000) + "\n"
	if _, err := Count(strings.NewReader(text)); err == nil {
		t.Errorf("Count with a 70000-byte line returned a nil error, want the error of scanner.Err()")
	}
}
