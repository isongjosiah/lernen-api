package api

import (
	"context"
	"net/http"
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
			WriteErrorResponse(w, http.StatusForbidden, "You are unable to access this page")
			return
		}

		email, err := VerifyToken("", authorization)

		if err != nil {
			WriteErrorResponse(w, http.StatusForbidden, "You do not have the authorization to view this page, Please sign in again.")
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
