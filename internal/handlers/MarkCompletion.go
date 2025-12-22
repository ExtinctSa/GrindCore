package handler

import (
	"net/http"

	"github.com/ExtinctSa/final_project/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) MarkCompletion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := UserFromContext(r.Context())
	if !ok {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	habitIDstr := r.PathValue("id")
	if habitIDstr == "" {
		http.Error(w, `{"error":"habit id is required"}`, http.StatusBadRequest)
		return
	}
	habitID, err := uuid.Parse(habitIDstr)
	if err != nil {
		http.Error(w, `{"error":"invalid habit id"}`, http.StatusBadRequest)
		return
	}

	err = cfg.DBQueries.MarkCompletion(r.Context(), database.MarkCompletionParams{
		HabitID: habitID,
		UserID:  user.ID,
	})
	if err != nil {
		http.Error(w, `{"error":"could not mark completion"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
