package httpapi

import (
	"net/http"
	"rest-notes-api/internal/service"
)

func NewHandler(svc *service.NoteService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", health)

	mux.HandleFunc("GET /echo", echo)

	mux.HandleFunc("GET /notes", getAllNotesHandler(svc))
	mux.HandleFunc("GET /notes/{id}", getNoteHandler(svc))
	mux.HandleFunc("POST /notes", createNoteHandler(svc))
	mux.HandleFunc("DELETE /notes/{id}", deleteNoteHandler(svc))
	mux.HandleFunc("PUT /notes/{id}", updateNoteHandler(svc))

	return WithPanicRecovery(WithLogger(mux))
}
