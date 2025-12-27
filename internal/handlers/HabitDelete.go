package handler

import (
	"net/http"

	"github.com/ExtinctSa/final_project/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) DeleteHabit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := UserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	habitIDstr := r.PathValue("habitID")
	habitID, err := uuid.Parse(habitIDstr)
	if err != nil {
		http.Error(w, `{"error":"invalid habit ID"}`, http.StatusBadRequest)
		return
	}

	err = cfg.DBQueries.DeleteHabit(r.Context(), database.DeleteHabitParams{
		UserID: user.ID,
		ID:     habitID,
	})
	if err != nil {
		http.Error(w, `{"error":"could not delete habit"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
