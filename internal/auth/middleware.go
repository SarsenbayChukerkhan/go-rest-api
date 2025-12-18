package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "missing token", 401)
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid token format", 401)
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
