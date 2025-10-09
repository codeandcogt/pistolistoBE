package middleware

import (
	"context"
	"net/http"
	"pistolistoBE/internal/common"
	"pistolistoBE/internal/config"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKeyAdmin string

const adminIDKey contextKeyAdmin = "id_administrativo"

func AdminJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		ctx := context.WithValue(r.Context(), adminIDKey, claims["id_administrativo"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
