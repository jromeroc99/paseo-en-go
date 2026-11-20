package site

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"paseo-en-go/internal/course"
	"paseo-en-go/web"
)

func testCourse() course.Course {
	return course.Course{
		Title: "Test Tour",
		Modules: []course.Module{{
			Slug:  "00-intro",
			Title: "Intro",
			Lessons: []course.Lesson{
				{Slug: "01-first", Title: "First"},
				{Slug: "02-second", Title: "Second", Exercise: &course.Exercise{Path: "00-intro/hello", Goal: "Greet"}},
			},
		}},
	}
}

func testContent() fstest.MapFS {
	return fstest.MapFS{
		"00-intro/01-first.html":      {Data: []byte(`<p>First body</p>{{code "examples/x/main.go"}}`)},
		"00-intro/02-second.html":     {Data: []byte(`<p>Second body</p>`)},
		"00-intro/examples/x/main.go": {Data: []byte(`if a < b { fmt.Println("hi") }`)},
	}
}

func newTestSite(t *testing.T, base string) *Site {
	t.Helper()
	s, err := New(testCourse(), testContent(), web.FS, base)
	if err != nil {
		t.Fatalf("New() = %v", err)
	}
	return s
}

func get(t *testing.T, h http.Handler, path string) (int, string) {
	t.Helper()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
	body, _ := io.ReadAll(rec.Result().Body)
	return rec.Code, string(body)
}

func TestPages(t *testing.T) {
	h := newTestSite(t, "/").Handler()

	tests := []struct {
		name     string
		path     string
		wantCode int
		want     []string
	}{
		{
			name:     "index lists every lesson",
			path:     "/",
			wantCode: http.StatusOK,
			want:     []string{"Test Tour", `href="/00-intro/01-first/"`, `href="/00-intro/02-second/"`},
		},
		{
			name:     "lesson shows body and escaped snippet",
			path:     "/00-intro/01-first/",
			wantCode: http.StatusOK,
			want:     []string{"<h1>First</h1>", "<p>First body</p>", "if a &lt; b", `data-done="00-intro/01-first"`, `rel="next"`},
		},
		{
			name:     "lesson shows its exercise",
			path:     "/00-intro/02-second/",
			wantCode: http.StatusOK,
			want:     []string{"exercises/00-intro/hello", "go test ./00-intro/hello", `rel="prev"`},
		},
		{
			name:     "static assets are served",
			path:     "/static/css/style.css",
			wantCode: http.StatusOK,
			want:     []string{":root"},
		},
		{
			name:     "unknown lesson is not found",
			path:     "/00-intro/99-missing/",
			wantCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, body := get(t, h, tt.path)
			if code != tt.wantCode {
				t.Fatalf("GET %s = %d, want %d", tt.path, code, tt.wantCode)
			}
			for _, w := range tt.want {
				if !strings.Contains(body, w) {
					t.Errorf("GET %s: body does not contain %q", tt.path, w)
				}
			}
		})
	}
}

func TestBasePathPrefixesLinks(t *testing.T) {
	_, body := get(t, newTestSite(t, "/paseo/").Handler(), "/")

	for _, w := range []string{`href="/paseo/static/css/style.css"`, `href="/paseo/00-intro/01-first/"`, `src="/paseo/static/js/animations.js"`} {
		if !strings.Contains(body, w) {
			t.Errorf("index does not contain %q", w)
		}
	}
}

func TestNewFailsWhenLessonBodyIsMissing(t *testing.T) {
	content := testContent()
	delete(content, "00-intro/02-second.html")

	if _, err := New(testCourse(), content, web.FS, "/"); err == nil {
		t.Fatal("New() = nil error, want error for missing lesson body")
	}
}

func TestBuildWritesEveryPage(t *testing.T) {
	dir := t.TempDir()

	if err := newTestSite(t, "/").Build(dir); err != nil {
		t.Fatalf("Build() = %v", err)
	}

	pages := map[string]string{
		"index.html":                          "Test Tour",
		"00-intro/01-first/index.html":        "First body",
		"00-intro/02-second/index.html":       "Second body",
		"static/css/style.css":                ":root",
		"static/js/progress.js":               "localStorage",
		"static/js/animations.js":             "import(",
		"static/js/animations/pixel-scene.js": "class PixelScene",
	}
	for file, want := range pages {
		data, err := os.ReadFile(filepath.Join(dir, file))
		if err != nil {
			t.Errorf("missing %s: %v", file, err)
			continue
		}
		if !strings.Contains(string(data), want) {
			t.Errorf("%s does not contain %q", file, want)
		}
	}
}
