package handler

import (
	"context"
	"net/http"

	"github.com/ExtinctSa/final_project/internal/auth"
	"github.com/ExtinctSa/final_project/internal/database"
)

type contextKey string

const userKey contextKey = "lowkenuinely"

func (cfg *ApiConfig) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		userID, err := auth.ValidateJWT(tokenString, cfg.Sk)
		if err != nil {
			http.Error(w, "Unathorized: Invalid Token", http.StatusUnauthorized)
			return
		}

		user, err := cfg.DBQueries.GetUserByID(r.Context(), userID)
		if err != nil {
			http.Error(w, "Unauthorized: User not found", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserFromContext(ctx context.Context) (database.GetUserByIDRow, bool) {
	user, ok := ctx.Value(userKey).(database.GetUserByIDRow)
	return user, ok
}
