package httpapi

import "net/http"

func echo(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		writeError(w, http.StatusBadRequest, "msg is required")
		return
	}

	writeJSON(w, 200, map[string]string{
		"echo": msg,
	})
}
