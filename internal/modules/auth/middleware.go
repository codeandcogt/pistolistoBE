package auth

import (
	"context"
	"net/http"
	"pistolistoBE/internal/config"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.JwtSecret)

type contextKey string

const clientIDKey contextKey = "id_cliente"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token requerido", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), clientIDKey, claims["id_cliente"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
