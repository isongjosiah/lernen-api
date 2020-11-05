package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func (a *API) UserRoutes(router *chi.Mux) http.Handler {
	router.Method("GET", "/home", http.HandlerFunc(a.UserHome))
	return router
}

func (a *API)UserHome(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path); err != nil {
		return
	}
}