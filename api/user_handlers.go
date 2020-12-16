package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *API) UserRoutes(router *chi.Mux) http.Handler {
	router.Method("GET", "/home", http.HandlerFunc(a.UserHome))
	router.Method("POST", "/edit", http.HandlerFunc(a.EditProfile))
	return router
}

func (a *API) UserHome(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path); err != nil {
		return
	}
}

func (a *API) EditProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
	file, header, err := r.FormFile("profile_picture")
	if err != nil {
		log.Error(err)
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("unable to parse file"))
		return
	}
	fmt.Print("HERE")
	defer file.Close()

	filename := header.Filename
	res, err := a.Deps.AWS.S3.Upload(filename, file)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, errors.Wrapf(err, "Could not upload file"))
		return
	}

	WriteJSONPayload(w, &ServerResponse{
		Message:    "Image successfully uploaded",
		StatusCode: http.StatusOK,
		Context:    nil,
		Payload:    res,
	})
}
