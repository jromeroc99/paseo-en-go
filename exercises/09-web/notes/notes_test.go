package notes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"testing"
)

// call sends one request to h and returns the recorded answer.
func call(h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(method, path, strings.NewReader(body)))
	return rec
}

func TestAPI(t *testing.T) {
	h := NewHandler()

	rec := call(h, "GET", "/api/notes", "")
	if got := strings.TrimSpace(rec.Body.String()); rec.Code != http.StatusOK || got != "[]" {
		t.Errorf("GET /api/notes with no notes = %d %q, want %d %q", rec.Code, got, http.StatusOK, "[]")
	}

	for i, text := range []string{"hola", "adiós"} {
		rec := call(h, "POST", "/api/notes", `{"text":"`+text+`"}`)
		if rec.Code != http.StatusCreated {
			t.Fatalf("POST /api/notes: status = %d, want %d", rec.Code, http.StatusCreated)
		}
		var got Note
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("POST /api/notes: the body %q is not a JSON note: %v", rec.Body.String(), err)
		}
		if want := (Note{ID: i + 1, Text: text}); got != want {
			t.Errorf("POST /api/notes: answered %+v, want %+v", got, want)
		}
	}

	rec = call(h, "GET", "/api/notes", "")
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("GET /api/notes: Content-Type = %q, want %q", got, "application/json")
	}
	var notes []Note
	if err := json.Unmarshal(rec.Body.Bytes(), &notes); err != nil {
		t.Fatalf("GET /api/notes: the body %q is not a JSON array: %v", rec.Body.String(), err)
	}
	want := []Note{{ID: 1, Text: "hola"}, {ID: 2, Text: "adiós"}}
	if !reflect.DeepEqual(notes, want) {
		t.Errorf("GET /api/notes after two POSTs = %+v, want %+v", notes, want)
	}
}

func TestCreateRejectsBadInput(t *testing.T) {
	for _, body := range []string{"hola", `{"text":""}`, `{}`} {
		t.Run(body, func(t *testing.T) {
			rec := call(NewHandler(), "POST", "/api/notes", body)
			if rec.Code != http.StatusBadRequest {
				t.Errorf("POST /api/notes %s: status = %d, want %d", body, rec.Code, http.StatusBadRequest)
			}
		})
	}
}

func TestPage(t *testing.T) {
	tests := []struct {
		name    string
		notes   []string
		want    []string // pieces the page must contain
		notWant []string // pieces the page must not contain
	}{
		{
			name:    "no notes",
			want:    []string{"<p>Todavía no hay notas.</p>"},
			notWant: []string{"<li>"},
		},
		{
			name:  "two notes",
			notes: []string{"hola", "adiós"},
			want:  []string{"<ul>", "<li>hola</li>", "<li>adiós</li>", "</ul>"},
		},
		{
			name:    "HTML in a note is escaped",
			notes:   []string{"<b>hola</b>"},
			want:    []string{"<li>&lt;b&gt;hola&lt;/b&gt;</li>"},
			notWant: []string{"<b>"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler()
			for _, text := range tt.notes {
				note, _ := json.Marshal(Note{Text: text})
				call(h, "POST", "/api/notes", string(note))
			}

			rec := call(h, "GET", "/", "")
			if rec.Code != http.StatusOK {
				t.Fatalf("GET /: status = %d, want %d", rec.Code, http.StatusOK)
			}
			if got := rec.Header().Get("Content-Type"); !strings.HasPrefix(got, "text/html") {
				t.Errorf("GET /: Content-Type = %q, want text/html", got)
			}
			page := rec.Body.String()
			for _, piece := range tt.want {
				if !strings.Contains(page, piece) {
					t.Errorf("the page does not contain %q:\n%s", piece, page)
				}
			}
			for _, piece := range tt.notWant {
				if strings.Contains(page, piece) {
					t.Errorf("the page contains %q, but it should not:\n%s", piece, page)
				}
			}
		})
	}
}

func TestManyRequestsAtOnce(t *testing.T) {
	h := NewHandler()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			call(h, "POST", "/api/notes", `{"text":"nota"}`)
		}()
		go func() {
			defer wg.Done()
			call(h, "GET", "/", "")
		}()
	}
	wg.Wait()

	var notes []Note
	json.Unmarshal(call(h, "GET", "/api/notes", "").Body.Bytes(), &notes)
	if len(notes) != 50 {
		t.Errorf("50 POSTs at the same time left %d notes, want 50: protect the notes with a sync.Mutex", len(notes))
	}
}
