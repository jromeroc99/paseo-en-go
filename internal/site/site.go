// Package site renders the course as HTML pages. The same pages are
// served over HTTP during development and written to disk for publishing.
package site

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"paseo-en-go/internal/course"
)

// Site holds every page of the course, rendered once at startup.
type Site struct {
	pages  map[string][]byte // URL path without the base, e.g. "00-setup/01-welcome/"
	static fs.FS
}

// New renders the whole course. content holds the lesson bodies and
// examples; assets holds templates/ and static/. base is the URL prefix
// the site is published under, e.g. "/" or "/paseo-en-go/".
func New(c course.Course, content, assets fs.FS, base string) (*Site, error) {
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid course: %w", err)
	}
	static, err := fs.Sub(assets, "static")
	if err != nil {
		return nil, err
	}

	r := renderer{course: c, content: content, assets: assets, base: base}
	pages, err := r.renderAll()
	if err != nil {
		return nil, err
	}
	return &Site{pages: pages, static: static}, nil
}

// Handler serves the rendered pages and the static assets.
func (s *Site) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServerFS(s.static)))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, ok := s.pages[strings.TrimPrefix(r.URL.Path, "/")]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(page)
	})
	return mux
}

// Build writes the site as static files into dir, ready to be hosted.
func (s *Site) Build(dir string) error {
	for urlPath, page := range s.pages {
		file := filepath.Join(dir, filepath.FromSlash(urlPath), "index.html")
		if err := writeFile(file, page); err != nil {
			return err
		}
	}
	return os.CopyFS(filepath.Join(dir, "static"), s.static)
}

func writeFile(name string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(name), 0o755); err != nil {
		return err
	}
	return os.WriteFile(name, data, 0o644)
}

// renderer turns the course into HTML pages.
type renderer struct {
	course  course.Course
	content fs.FS
	assets  fs.FS
	base    string
}

// lessonPage is the data the lesson template receives.
type lessonPage struct {
	course.Ref
	Course     course.Course
	Base       string
	Body       template.HTML
	Prev, Next *course.Ref
}

func (r renderer) renderAll() (map[string][]byte, error) {
	pages := map[string][]byte{}

	index, err := r.page("index.html", map[string]any{"Course": r.course, "Base": r.base})
	if err != nil {
		return nil, err
	}
	pages[""] = index

	lessons := r.course.Lessons()
	for i, ref := range lessons {
		data := lessonPage{Ref: ref, Course: r.course, Base: r.base}
		if i > 0 {
			data.Prev = &lessons[i-1]
		}
		if i < len(lessons)-1 {
			data.Next = &lessons[i+1]
		}
		if data.Body, err = r.body(ref); err != nil {
			return nil, err
		}
		if pages[ref.ID()+"/"], err = r.page("lesson.html", data); err != nil {
			return nil, err
		}
	}
	return pages, nil
}

// page renders one template from templates/ inside the shared layout.
func (r renderer) page(name string, data any) ([]byte, error) {
	tmpl, err := template.ParseFS(r.assets, "templates/layout.html", "templates/"+name)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, "layout", data); err != nil {
		return nil, fmt.Errorf("render %s: %w", name, err)
	}
	return buf.Bytes(), nil
}

// body renders a lesson body. Inside it, {{code "examples/x/main.go"}}
// embeds a file from the lesson's module folder as an escaped code block,
// and {{codeLines "examples/x/main.go" 1 17}} embeds only those lines, so a
// program too long for one card is read across several while it is still
// compiled whole.
func (r renderer) body(ref course.Ref) (template.HTML, error) {
	src, err := fs.ReadFile(r.content, ref.BodyPath())
	if err != nil {
		return "", fmt.Errorf("lesson %q: %w", ref.ID(), err)
	}

	read := func(file string) (string, error) {
		snippet, err := fs.ReadFile(r.content, path.Join(ref.Module.Slug, file))
		return string(snippet), err
	}
	code := func(file string) (template.HTML, error) {
		snippet, err := read(file)
		if err != nil {
			return "", err
		}
		return codeBlock(snippet), nil
	}
	codeLines := func(file string, first, last int) (template.HTML, error) {
		snippet, err := read(file)
		if err != nil {
			return "", err
		}
		lines := strings.SplitAfter(strings.TrimSuffix(snippet, "\n"), "\n")
		if first < 1 || last < first || last > len(lines) {
			return "", fmt.Errorf("codeLines %s %d %d: the file has lines 1 to %d", file, first, last, len(lines))
		}
		return codeBlock(strings.Join(lines[first-1:last], "")), nil
	}

	funcs := template.FuncMap{"code": code, "codeLines": codeLines}
	tmpl, err := template.New(ref.ID()).Funcs(funcs).Parse(string(src))
	if err != nil {
		return "", fmt.Errorf("lesson %q: %w", ref.ID(), err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		return "", fmt.Errorf("lesson %q: %w", ref.ID(), err)
	}
	// Lesson bodies are authored by us, not by users, so they are trusted HTML.
	return template.HTML(buf.String()), nil
}

// codeBlock wraps Go source in an escaped code block.
func codeBlock(src string) template.HTML {
	return template.HTML(`<pre><code class="language-go">` + html.EscapeString(src) + `</code></pre>`)
}
