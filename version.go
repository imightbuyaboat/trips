package main

import "net/http"

// Version задаётся при релизе; при сборке можно переопределить:
// go build -ldflags "-X main.Version=v1.0.0"
var Version = "v1.0.0"

// versionHandler возвращает текущую версию приложения.
func versionHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"version": Version})
}
