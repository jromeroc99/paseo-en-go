package request

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Request
	}{
		{
			name:  "request line only",
			input: "GET / HTTP/1.1\r\n\r\n",
			want:  Request{Method: "GET", Path: "/", Headers: map[string]string{}},
		},
		{
			name:  "with headers",
			input: "GET /notes HTTP/1.1\r\nHost: localhost:8080\r\nAccept: */*\r\n\r\n",
			want: Request{Method: "GET", Path: "/notes", Headers: map[string]string{
				"Host":   "localhost:8080",
				"Accept": "*/*",
			}},
		},
		{
			name:  "the body is not a header",
			input: "POST /notes HTTP/1.1\r\nContent-Length: 4\r\n\r\nhola",
			want: Request{Method: "POST", Path: "/notes", Headers: map[string]string{
				"Content-Length": "4",
			}},
		},
		{
			name:  "lines ending in \\n only",
			input: "DELETE /notes/3 HTTP/1.1\nHost: localhost\n\n",
			want: Request{Method: "DELETE", Path: "/notes/3", Headers: map[string]string{
				"Host": "localhost",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("Parse(%q) returned the error %v, want no error", tt.input, err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse(%q)\n got %+v\nwant %+v", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseRejectsBadRequests(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "empty input", input: ""},
		{name: "request line with two parts", input: "GET /\r\n\r\n"},
		{name: "request line with four parts", input: "GET / HTTP/1.1 extra\r\n\r\n"},
		{name: "header without colon", input: "GET / HTTP/1.1\r\nHost localhost\r\n\r\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Parse(strings.NewReader(tt.input)); err == nil {
				t.Errorf("Parse(%q) returned no error, want one", tt.input)
			}
		})
	}
}
