package todos

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

// listTodos asks h for the list and decodes it.
func listTodos(t *testing.T, h http.Handler) []Todo {
	t.Helper()
	rec := call(h, "GET", "/todos", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("GET /todos: status = %d, want %d", rec.Code, http.StatusOK)
	}
	var todos []Todo
	if err := json.Unmarshal(rec.Body.Bytes(), &todos); err != nil {
		t.Fatalf("GET /todos: the body %q is not a JSON array: %v", rec.Body.String(), err)
	}
	return todos
}

func TestEmptyList(t *testing.T) {
	rec := call(NewHandler(), "GET", "/todos", "")
	if rec.Code != http.StatusOK {
		t.Errorf("GET /todos: status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "[]" {
		t.Errorf("GET /todos on an empty list: body = %q, want %q (a nil slice is encoded as null)", got, "[]")
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("GET /todos: Content-Type = %q, want %q", got, "application/json")
	}
}

func TestCreate(t *testing.T) {
	h := NewHandler()
	for i, title := range []string{"comprar pan", "llamar a Ana"} {
		rec := call(h, "POST", "/todos", `{"title":"`+title+`"}`)
		if rec.Code != http.StatusCreated {
			t.Fatalf("POST /todos: status = %d, want %d", rec.Code, http.StatusCreated)
		}
		if got := rec.Header().Get("Content-Type"); got != "application/json" {
			t.Errorf("POST /todos: Content-Type = %q, want %q", got, "application/json")
		}
		var got Todo
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatalf("POST /todos: the body %q is not a JSON task: %v", rec.Body.String(), err)
		}
		if want := (Todo{ID: i + 1, Title: title}); got != want {
			t.Errorf("POST /todos: answered %+v, want %+v", got, want)
		}
	}

	want := []Todo{{ID: 1, Title: "comprar pan"}, {ID: 2, Title: "llamar a Ana"}}
	if got := listTodos(t, h); !reflect.DeepEqual(got, want) {
		t.Errorf("GET /todos after two POSTs = %+v, want %+v", got, want)
	}
}

func TestCreateRejectsBadInput(t *testing.T) {
	tests := []struct {
		name, body string
	}{
		{name: "not JSON", body: "comprar pan"},
		{name: "empty title", body: `{"title":""}`},
		{name: "no title", body: `{}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandler()
			rec := call(h, "POST", "/todos", tt.body)
			if rec.Code != http.StatusBadRequest {
				t.Errorf("POST /todos %s: status = %d, want %d", tt.body, rec.Code, http.StatusBadRequest)
			}
			if got := listTodos(t, h); len(got) != 0 {
				t.Errorf("a rejected POST added %+v to the list", got)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		path       string
		wantStatus int
		wantLeft   []Todo
	}{
		{path: "/todos/1", wantStatus: http.StatusNoContent, wantLeft: []Todo{{ID: 2, Title: "dos"}}},
		{path: "/todos/2", wantStatus: http.StatusNoContent, wantLeft: []Todo{{ID: 1, Title: "uno"}}},
		{path: "/todos/7", wantStatus: http.StatusNotFound, wantLeft: []Todo{{ID: 1, Title: "uno"}, {ID: 2, Title: "dos"}}},
		{path: "/todos/uno", wantStatus: http.StatusBadRequest, wantLeft: []Todo{{ID: 1, Title: "uno"}, {ID: 2, Title: "dos"}}},
	}
	for _, tt := range tests {
		t.Run("DELETE "+tt.path, func(t *testing.T) {
			h := NewHandler()
			call(h, "POST", "/todos", `{"title":"uno"}`)
			call(h, "POST", "/todos", `{"title":"dos"}`)

			rec := call(h, "DELETE", tt.path, "")
			if rec.Code != tt.wantStatus {
				t.Errorf("DELETE %s: status = %d, want %d", tt.path, rec.Code, tt.wantStatus)
			}
			if got := listTodos(t, h); !reflect.DeepEqual(got, tt.wantLeft) {
				t.Errorf("after DELETE %s the list is %+v, want %+v", tt.path, got, tt.wantLeft)
			}
		})
	}
}

func TestIDsAreNotReused(t *testing.T) {
	h := NewHandler()
	call(h, "POST", "/todos", `{"title":"uno"}`)
	call(h, "DELETE", "/todos/1", "")
	call(h, "POST", "/todos", `{"title":"dos"}`)

	want := []Todo{{ID: 2, Title: "dos"}}
	if got := listTodos(t, h); !reflect.DeepEqual(got, want) {
		t.Errorf("POST, DELETE /todos/1, POST: list = %+v, want %+v (an ID is never given twice)", got, want)
	}
}

func TestManyRequestsAtOnce(t *testing.T) {
	h := NewHandler()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			call(h, "POST", "/todos", `{"title":"tarea"}`)
		}()
	}
	wg.Wait()

	todos := listTodos(t, h)
	if len(todos) != 50 {
		t.Fatalf("50 POSTs at the same time left %d tasks, want 50: protect the list with a sync.Mutex", len(todos))
	}
	seen := map[int]bool{}
	for _, todo := range todos {
		if seen[todo.ID] {
			t.Errorf("two tasks got the ID %d: protect the counter with the same sync.Mutex", todo.ID)
		}
		seen[todo.ID] = true
	}
}
