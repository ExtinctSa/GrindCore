package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ExtinctSa/final_project/internal/auth"
)

func (cfg *ApiConfig) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get refresh token from Authorization: Bearer <token>
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, `{"error":"invalid or missing authorization header"}`, http.StatusUnauthorized)
		return
	}

	//Load refresh token from DB
	dbToken, err := cfg.DBQueries.GetRefreshToken(r.Context(), refreshToken)
	if err != nil {
		http.Error(w, `{"error":"invalid refresh token"}`, http.StatusUnauthorized)
		return
	}

	//Check expiration
	if time.Now().After(dbToken.ExpiresAt) {
		http.Error(w, `{"error":"refresh token expired"}`, http.StatusUnauthorized)
		return
	}

	//Check if revoked
	if dbToken.RevokedAt.Valid {
		http.Error(w, `{"error":"refresh token revoked"}`, http.StatusUnauthorized)
		return
	}

	//Issue new access token
	token, err := auth.MakeJWT(dbToken.UserID.UUID, cfg.Sk, time.Hour)
	if err != nil {
		http.Error(w, `{"error":"could not create token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (cfg *ApiConfig) RefreshTokenRevokeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Extract refresh token from Authorization header
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, `{"error":"missing or invalid authorization header"}`, http.StatusUnauthorized)
		return
	}

	//Mark the refresh token as revoked in DB
	err = cfg.DBQueries.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		http.Error(w, `{"error":"could not revoke token"}`, http.StatusInternalServerError)
		return
	}

	//Success response
	w.WriteHeader(http.StatusNoContent)
}
