package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ExtinctSa/final_project/internal/database"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) ListHabitsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, ok := UserFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	category := r.URL.Query().Get("category")

	var habits []database.Habit
	var err error

	switch category {
	case "":
		habits, err = cfg.DBQueries.GetAllHabits(r.Context(), user.ID)
	case "uncategorized":
		habits, err = cfg.DBQueries.ListHabitsWithoutCategory(r.Context(), user.ID)
	default:
		habits, err = cfg.DBQueries.GetHabitByCategory(
			r.Context(),
			database.GetHabitByCategoryParams{
				UserID:   user.ID,
				Category: sql.NullString{String: category, Valid: true},
			},
		)
	}

	if err != nil {
		http.Error(w, "could not fetch habits", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(habits)
}

func (cfg *ApiConfig) ListHabitsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error":"missing habit ID"}`, http.StatusBadRequest)
		return
	}

	habitID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid habit ID"}`, http.StatusBadRequest)
		return
	}

	habit, err := cfg.DBQueries.GetHabitByID(r.Context(), habitID)
	if err != nil {
		http.Error(w, `{"error":"habit not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(habit)
}
