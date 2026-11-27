package page

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		name    string
		title   string
		items   []string
		want    []string // pieces the page must contain
		notWant []string // pieces the page must not contain
	}{
		{
			name:  "a list",
			title: "Compra",
			items: []string{"pan", "leche"},
			want:  []string{"<h1>Compra</h1>", "<ul>", "<li>pan</li>", "<li>leche</li>", "</ul>"},
		},
		{
			name:    "an empty list",
			title:   "Compra",
			items:   nil,
			want:    []string{"<h1>Compra</h1>", "<p>No hay nada.</p>"},
			notWant: []string{"<ul>", "<li>"},
		},
		{
			name:    "HTML in an item is escaped",
			title:   "Compra",
			items:   []string{"<b>sal</b>"},
			want:    []string{"<li>&lt;b&gt;sal&lt;/b&gt;</li>"},
			notWant: []string{"<b>"},
		},
		{
			name:    "HTML in the title is escaped",
			title:   "<script>alert(1)</script>",
			items:   []string{"pan"},
			want:    []string{"<h1>&lt;script&gt;alert(1)&lt;/script&gt;</h1>"},
			notWant: []string{"<script>"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var page strings.Builder
			if err := Render(&page, tt.title, tt.items); err != nil {
				t.Fatalf("Render returned the error %v, want no error", err)
			}
			for _, piece := range tt.want {
				if !strings.Contains(page.String(), piece) {
					t.Errorf("the page does not contain %q:\n%s", piece, page.String())
				}
			}
			for _, piece := range tt.notWant {
				if strings.Contains(page.String(), piece) {
					t.Errorf("the page contains %q, but it should not:\n%s", piece, page.String())
				}
			}
		})
	}
}

func TestRenderKeepsTheOrder(t *testing.T) {
	var page strings.Builder
	if err := Render(&page, "Días", []string{"lunes", "martes", "miércoles"}); err != nil {
		t.Fatalf("Render returned the error %v, want no error", err)
	}
	lunes := strings.Index(page.String(), "lunes")
	martes := strings.Index(page.String(), "martes")
	miercoles := strings.Index(page.String(), "miércoles")
	if !(lunes < martes && martes < miercoles) {
		t.Errorf("the items are not in the order they were given:\n%s", page.String())
	}
}
