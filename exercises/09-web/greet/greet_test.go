package greet

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		method, path string
		wantStatus   int
		wantBody     string
	}{
		{method: "GET", path: "/hello/Ana", wantStatus: http.StatusOK, wantBody: "Hola, Ana\n"},
		{method: "GET", path: "/hello/Gopher", wantStatus: http.StatusOK, wantBody: "Hola, Gopher\n"},
		{method: "GET", path: "/hello/Mar%C3%ADa%20Jos%C3%A9", wantStatus: http.StatusOK, wantBody: "Hola, María José\n"},
		{method: "GET", path: "/bye/Ana", wantStatus: http.StatusNotFound},
		{method: "POST", path: "/hello/Ana", wantStatus: http.StatusMethodNotAllowed},
	}
	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			NewHandler().ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))

			if rec.Code != tt.wantStatus {
				t.Errorf("%s %s: status = %d, want %d", tt.method, tt.path, rec.Code, tt.wantStatus)
			}
			if tt.wantBody != "" && rec.Body.String() != tt.wantBody {
				t.Errorf("%s %s: body = %q, want %q", tt.method, tt.path, rec.Body.String(), tt.wantBody)
			}
		})
	}
}
