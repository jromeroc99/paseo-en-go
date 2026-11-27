package table

import (
	"bytes"
	"testing"
)

func TestWriteTable(t *testing.T) {
	tests := []struct {
		name string
		rows [][]string
		want string
	}{
		{
			name: "no rows",
			rows: nil,
			want: "",
		},
		{
			name: "one cell",
			rows: [][]string{{"hola"}},
			want: "hola\n",
		},
		{
			name: "two columns line up",
			rows: [][]string{{"Nombre", "Edad"}, {"Ana", "31"}, {"Luis", "7"}},
			want: "Nombre | Edad\nAna    | 31\nLuis   | 7\n",
		},
		{
			name: "the widest cell is not in the first row",
			rows: [][]string{{"x", "y", "z"}, {"uno", "dos", "tres"}},
			want: "x   | y   | z\nuno | dos | tres\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			WriteTable(&buf, tt.rows)
			if got := buf.String(); got != tt.want {
				t.Errorf("WriteTable(%q) wrote\n%q\nwant\n%q", tt.rows, got, tt.want)
			}
		})
	}
}

func TestWriteTableAppends(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("Tabla:\n")
	WriteTable(&buf, [][]string{{"a", "b"}})
	want := "Tabla:\na | b\n"
	if got := buf.String(); got != want {
		t.Errorf("with \"Tabla:\\n\" already in buf, buf = %q, want %q: append to buf, keep what it held", got, want)
	}
}
