package middleware

import (
	"context"
	"fmt"
	"net/http"
	"pistolistoBE/internal/common"
	"pistolistoBE/internal/config"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const clientIDKey contextKey = "id_cliente"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("valores de error")

			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		// fmt.Println(token, err, "valores token")
		if err != nil || !token.Valid {
			fmt.Println(err, "valores de error")

			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		ctx := context.WithValue(r.Context(), clientIDKey, claims["id_cliente"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
