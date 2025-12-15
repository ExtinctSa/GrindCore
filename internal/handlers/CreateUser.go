package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ExtinctSa/final_project/internal/auth"
	"github.com/ExtinctSa/final_project/internal/database"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type createUserResponse struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func (cfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil || req.Email == "" || req.Password == "" || req.Username == "" {
		http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
		return
	}
	//hashes inputted password before storing to database
	hpass, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(
			w,
			`{"error":"could not create user","details":"`+err.Error()+`"}`,
			http.StatusInternalServerError,
		)
		return
	}
	user, err := cfg.DBQueries.CreateUserByEmail(r.Context(), database.CreateUserByEmailParams{
		Email:          req.Email,
		HashedPassword: hpass,
		Username:       req.Username,
	})
	if err != nil {
		http.Error(
			w,
			`{"error":"could not create user","details":"`+err.Error()+`"}`,
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createUserResponse{
		ID:        user.ID.String(),
		CreatedAt: user.CreatedAt.UTC().Format(time.RFC3339),
		Username:  user.Username,
		Email:     user.Email,
	})
}
