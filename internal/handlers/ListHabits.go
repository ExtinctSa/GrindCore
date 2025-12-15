package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ExtinctSa/final_project/internal/database"
)

func (cfg *ApiConfig) ListHabitsHandler(w http.ResponseWriter, r *http.Request) {
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
