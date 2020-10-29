package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (a *API) UserRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/register", http.HandlerFunc(a.Register))
	router.Method("POST", "/login", http.HandlerFunc(a.Login))

	return router
}
