package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ExtinctSa/final_project/internal/database"
)

type createHabitRequest struct {
	HabitName string `json:"habitName"`
	Frequency string `json:"frequency"`
	Category  string `json:"category"`
	UserID    string `json:"user_id"`
}

type createHabitResponse struct {
	HabitName string         `json:"habitName"`
	Frequency string         `json:"frequency"`
	Category  sql.NullString `json:"category"`
	UserID    string         `json:"user_id"`
	ID        string         `json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

func (cfg *ApiConfig) CreateHabitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := UserFromContext(r.Context())
	if !ok {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var req createHabitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}

	if req.HabitName == "" || req.Frequency == "" {
		http.Error(w, `{"error":"missing required fields"}`, http.StatusBadRequest)
		return
	}

	habit, err := cfg.DBQueries.CreateHabit(r.Context(), database.CreateHabitParams{
		Habitname: req.HabitName,
		Frequency: req.Frequency,
		Category:  toNullString(req.Category),
		UserID:    user.ID,
	})
	if err != nil {
		http.Error(w, `{"error":"could not create habit"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createHabitResponse{
		ID:        habit.ID.String(),
		HabitName: habit.Habitname,
		Frequency: habit.Frequency,
		Category:  habit.Category,
		UserID:    habit.UserID.String(),
		CreatedAt: habit.CreatedAt.Format(time.RFC3339),
		UpdatedAt: habit.UpdatedAt.Format(time.RFC3339),
	})
}
