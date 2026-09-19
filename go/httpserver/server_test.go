package httpserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	// httptest se asli server chalaye bina handler test hota hai
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, httptest.NewRequest("GET", "/health", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("code = %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"status":"ok"`) {
		t.Errorf("body = %s", rec.Body.String())
	}
}

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	NewMux().ServeHTTP(rec, httptest.NewRequest("GET", "/hello/tanishk", nil))
	if rec.Body.String() != "namaste tanishk" {
		t.Errorf("body = %s", rec.Body.String())
	}
}
