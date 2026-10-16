package main

import (
	"fmt"
	"net/http"
)

// helloHandler отвечает приветствием с путём запроса.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s\n", r.URL.Path)
}

// healthHandler возвращает {"status":"ok"} для проверки живости сервиса.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}
