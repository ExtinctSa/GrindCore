package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ExtinctSa/final_project/internal/auth"
	"github.com/ExtinctSa/final_project/internal/database"
)

type updateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (cfg *ApiConfig) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	accessToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	userID, err := auth.ValidateJWT(accessToken, cfg.Sk)
	if err != nil {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	var request updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	if request.Email == "" || request.Password == "" {
		http.Error(w, `{"error":"email and password required"}`, http.StatusBadRequest)
		return
	}
	hpass, err := auth.HashPassword(request.Password)
	if err != nil {
		http.Error(w, `{"error":"could not update user"}`, http.StatusInternalServerError)
		return
	}

	updatedUser, err := cfg.DBQueries.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:             userID,
		Email:          request.Email,
		HashedPassword: hpass,
	})
	if err != nil {
		http.Error(w, `{"error":"could not update user"}`, http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"id":         updatedUser.ID,
		"email":      updatedUser.Email,
		"username":   updatedUser.Username,
		"created_at": updatedUser.CreatedAt,
		"updated_at": updatedUser.UpdatedAt,
	}

	json.NewEncoder(w).Encode(resp)
}
