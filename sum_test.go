package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSumHandler(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantStatus int
		wantResult int
	}{
		{"positive", "a=2&b=3", http.StatusOK, 5},
		{"negative", "a=-5&b=2", http.StatusOK, -3},
		{"zero", "a=0&b=0", http.StatusOK, 0},
		{"missing b", "a=1", http.StatusBadRequest, 0},
		{"not a number", "a=x&b=1", http.StatusBadRequest, 0},
		{"empty", "", http.StatusBadRequest, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			sumHandler(rec, httptest.NewRequest(http.MethodGet, "/sum?"+tt.query, nil))

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantStatus != http.StatusOK {
				return
			}
			var body map[string]int
			if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
				t.Fatalf("invalid JSON: %v", err)
			}
			if body["result"] != tt.wantResult {
				t.Errorf("result = %d, want %d", body["result"], tt.wantResult)
			}
		})
	}
}
