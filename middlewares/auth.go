package middlewares

import (
	"context"
	"net/http"
	"os"

	"github.com/Alejandrocuartas/geophoto/helpers"
	"github.com/Alejandrocuartas/geophoto/types"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		jwtSecret := os.Getenv("JWT")
		userId, _ := helpers.VerifyToken(token, jwtSecret)
		ctx := context.WithValue(r.Context(), types.Auth, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
