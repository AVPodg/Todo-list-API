package httpapi

import (
	"net/http"
	"time"
)

func health(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC().Format(time.RFC3339)
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"time":   now,
	})
}
