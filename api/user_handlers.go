package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *API) UserRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/home", http.HandlerFunc(a.UserHome))
	router.Method("POST", "/edit", http.HandlerFunc(a.EditProfilePicture))
	return router
}

func (a *API) UserHome(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path); err != nil {
		return
	}
}

func (a *API) EditProfilePicture(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("profile_picture")
	if err != nil {
		log.Error(err)
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("unable to parse file"))
		return
	}
	defer file.Close()

	filename := header.Filename
	upRes, err := a.Deps.AWS.S3.Upload(filename, file)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, errors.Wrapf(err, "Could not upload file"))
		return
	}
	//save the output in the database and handle error
	dummyUsername := ""
	if err := a.Deps.DAL.UserDAL.EditPicture(dummyUsername, upRes); err != nil {
		// delete the file from AWS if there was an error saving information to DB
		delRes, aerr := a.Deps.AWS.S3.Delete("nothing")
		fmt.Println(aerr)
		if aerr != nil {
			log.Errorf("could not delete file from s3 bucket: %s", aerr)
			WriteErrorResponse(w, http.StatusInternalServerError, errors.Wrapf(aerr, "Could not upload file"))
			return
		}
		log.Info(delRes)

		// notify client of the error
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	WriteJSONPayload(w, &ServerResponse{
		Message:    "Image successfully uploaded",
		StatusCode: http.StatusOK,
		Context:    nil,
		Payload:    upRes,
	})
}
