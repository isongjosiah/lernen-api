package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

func (a *API) CourseRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/create", http.HandlerFunc(a.CreateCourse))
	router.Method("POST", "/delete", http.HandlerFunc(a.DeleteCourse))

	return router
}

//CreateCourse creates a course
func (a *API) CreateCourse(w http.ResponseWriter, r *http.Request) {

}

//DeleteCourse delete a course
func (a *API) DeleteCourse(w http.ResponseWriter, r *http.Request) {

}
