package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ExtinctSa/final_project/internal/auth"
	"github.com/ExtinctSa/final_project/internal/database"
	"github.com/google/uuid"
)

type createUserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type createUserLoginResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
}

func (cfg *ApiConfig) UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request createUserLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"error":"invalid JSON"}`, http.StatusBadRequest)
		return
	}

	user, err := cfg.DBQueries.GetUserByUsername(r.Context(), request.Username)
	if err != nil {
		http.Error(w, `{"error":"Incorrect username or password"}`, http.StatusUnauthorized)
		return
	}
	match, err := auth.CheckPasswordHash(request.Password, user.HashedPassword)
	if err != nil || !match {
		http.Error(w, `{"error":"Incorrect username or password"}`, http.StatusUnauthorized)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.Sk, time.Hour)
	if err != nil {
		http.Error(w, `{"error":"could not create token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		http.Error(w, `{"error":"could not create refresh token"}`, http.StatusInternalServerError)
		return
	}

	refreshExp := time.Now().Add(60 * 24 * time.Hour)

	_, err = cfg.DBQueries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
		ExpiresAt: refreshExp,
	})
	if err != nil {
		http.Error(w, `{"error":"could not save refresh token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createUserLoginResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Username:  user.Username,
		Email:     user.Email,
		Token:     accessToken,
	})
}
