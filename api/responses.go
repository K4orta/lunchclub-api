package api

import (
	"net/http"
)

func sendError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError) // 500
	w.Write([]byte(`{"error": "` + msg + `"}`))
}
