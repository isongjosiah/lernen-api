package api

import (
	"context"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//TODO(josiah): restructure the middleware properly. i.e carry out the token authentication elegantly
func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthJwtVerify(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		res := &ServerResponse{}
		header := r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			WriteErrorResponse(w, http.StatusForbidden, "Authorization failed")
			return
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) { return []byte(os.Getenv("SECRET")), nil })

		if err != nil {
			WriteJSONPayload(w, res)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(r.Context(), "userID", claims["userID"])

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
