package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"upload-service/pkg/common"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.APIResponse{Error: "Not authenticated"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.APIResponse{Error: "Not authenticated"})
			return
		}

		claims, err := common.ValidateJWT(tokenParts[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(common.APIResponse{Error: "Not authenticated"})
			return
		}

		// Add user info to request context
		ctx := context.WithValue(r.Context(), common.UserIDContextKey, claims.UserID)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
