package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// timeHandler возвращает текущее UTC-время в формате RFC3339 (только GET).
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"time": time.Now().UTC().Format(time.RFC3339),
	})
}
