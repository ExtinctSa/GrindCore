package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ExtinctSa/final_project/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) CheckCompletion(w http.ResponseWriter, r *http.Request) {
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

	result, err := cfg.DBQueries.CheckCompletion(r.Context(), database.CheckCompletionParams{
		HabitID: habitID,
		UserID:  user.ID,
	})
	if err != nil {
		http.Error(w, `{"error":"could not check habit completion"}`, http.StatusInternalServerError)
		return
	}

	var resp string

	if result {
		resp = `{"completed":"true"}`
	} else {
		resp = `{"completed":"false"}`
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
