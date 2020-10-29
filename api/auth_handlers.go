package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

//AuthRoutes sets up the authentication handlers
func (a *API) AuthRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/register", http.HandlerFunc(a.Login))
	router.Method("POST", "/register", http.HandlerFunc(a.Register))

	return router
}

//Register is the handler for the path /register
func (a *API) Register(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello you have requested: %s\n", r.URL.Path); err != nil {
		return
	}
}

//Login is the handler for the path /login
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello you have requested: %s\n", r.URL.Path); err != nil {
		return
	}
}
