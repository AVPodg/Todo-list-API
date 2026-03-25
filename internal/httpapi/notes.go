package httpapi

import (
	"encoding/json"
	"net/http"
	"rest-notes-api/internal/service"
)

func createNoteHandler(svc *service.NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json body")
			return
		}

		note, err := svc.CreateNote(r.Context(), req.Title, req.Content)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusCreated, note)
	}
}

func getAllNotesHandler(svc *service.NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes, err := svc.GetAllNotes(r.Context())
		if err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeJSON(w, http.StatusOK, notes)
	}
}

func getNoteHandler(svc *service.NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		note, err := svc.GetNoteById(r.Context(), id)
		if err != nil {
			writeError(w, http.StatusNotFound, "note not found")
			return
		}

		writeJSON(w, http.StatusOK, note)
	}
}

func updateNoteHandler(svc *service.NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		var req struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}

		if err := svc.UpdateNote(r.Context(), id, req.Title, req.Content); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteNoteHandler(svc *service.NoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		if err := svc.DeleteNote(r.Context(), id); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
