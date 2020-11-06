package api

import (
	"context"
	"errors"
	"github.com/isongjosiah/lernen-api/common"
	"net/http"
	"os"
	"strings"
)

func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// AuthenticateJWT validates that the jwt token sent along with a user requests was signed by the API
func AuthenticateJWT(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		authorization = strings.TrimSpace(authorization)

		if authorization == "" {
			WriteErrorResponse(w, http.StatusForbidden, errors.New("You are unable to access this page"))
			return
		}
		jwtSecret := os.Getenv("TOKEN_SECRET") // see if there is a better way to do this josiah
		email, err := VerifyToken(jwtSecret, authorization)

		if err != nil {
			WriteErrorResponse(w, http.StatusForbidden, errors.New("You do not have the authorization to view this page, Please sign in again."))
			return
		}

		ctx := context.WithValue(r.Context(), common.ContextKey("email"), email)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
