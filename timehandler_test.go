package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	timeHandler(rec, httptest.NewRequest(http.MethodGet, "/time", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if _, err := time.Parse(time.RFC3339, body["time"]); err != nil {
		t.Errorf("bad time %q: %v", body["time"], err)
	}
}

func TestTimeHandlerMethodNotAllowed(t *testing.T) {
	rec := httptest.NewRecorder()
	timeHandler(rec, httptest.NewRequest(http.MethodPost, "/time", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", rec.Code)
	}
}
