package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ExtinctSa/final_project/internal/database"
	"github.com/google/uuid"
)

type updateHabitRequest struct {
	HabitID   uuid.UUID `json:"id"`
	HabitName *string   `json:"habitName"`
	Frequency *string   `json:"frequency"`
	Category  *string   `json:"category"`
}

func toNullStringPtr(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func (cfg *ApiConfig) UpdateHabit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := UserFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req updateHabitRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if req.HabitID == uuid.Nil {
		http.Error(w, `{"error":"id is required"}`, http.StatusBadRequest)
		return
	}

	if req.HabitName == nil && req.Frequency == nil && req.Category == nil {
		http.Error(w, `{"error":"no fields to update"}`, http.StatusBadRequest)
		return
	}

	updatedHabit, err := cfg.DBQueries.UpdateHabit(
		r.Context(),
		database.UpdateHabitParams{
			ID:        req.HabitID,
			UserID:    user.ID,
			Habitname: toNullStringPtr(req.HabitName),
			Frequency: toNullStringPtr(req.Frequency),
			Category:  toNullStringPtr(req.Category),
		},
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, `{"error":"habit not found"}`, http.StatusNotFound)
			return
		}
		http.Error(w, `{"error":"could not update habit"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedHabit)
}
