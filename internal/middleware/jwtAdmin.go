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

type contextKeyAdmin string

const adminIDKey contextKeyAdmin = "id_administrativo"

func AdminJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			fmt.Println("sin header", authHeader)
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		claims := jwt.MapClaims{}
		fmt.Println(claims, "claves")
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})
		fmt.Println(token, err, "claves despues")
		if err != nil || !token.Valid {
			fmt.Println("token invalido", token)
			common.ErrorResponse(w, http.StatusUnauthorized, common.HTTP_UNAUTHORIZED, common.ERR_UNAUTHORIZED, nil)
			return
		}

		// if role, ok := claims["rol"].(uint); !ok || role != 1 {
		// 	common.ErrorResponse(w, http.StatusForbidden, common.HTTP_FORBIDDEN, "Acceso restringido a administradores", nil)
		// 	return
		// }

		ctx := context.WithValue(r.Context(), adminIDKey, claims["id_administrativo"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
