package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (a *API) UserRoutes(router *chi.Mux) http.Handler {
	return router
}
