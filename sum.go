package main

import (
	"net/http"
	"strconv"
)

// sumHandler складывает целые параметры a и b: GET /sum?a=2&b=3.
func sumHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))
	if errA != nil || errB != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "a and b must be integers"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]int{"result": a + b})
}
