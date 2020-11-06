package api

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/isongjosiah/lernen-api/dal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

//AuthRoutes sets up the authentication handlers
func (a *API) AuthRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/login", http.HandlerFunc(a.Login))
	router.Method("POST", "/register", http.HandlerFunc(a.Register))

	return router
}

//Register is the handler for the path /register
func (a *API) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	//read the information from the body. TODO(josiah): check if you need to define middlewares to set the content-type to "application/json"
	err := decodeJSONBody(nil, r.Body, &user)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	// check for empty fields
	if user.Firstname == "" || user.Lastname == "" || user.Email == "" || user.Username == "" || user.Password == "" {
		WriteErrorResponse(w, http.StatusBadRequest, errors.New("some required fields are empty. Please fill all fields"))
		return
	}
	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	fmt.Println("DEBUG1")
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	user.Password = string(hash)
	fmt.Println("DEBUG2")
	fmt.Printf("%T", &user)
	// add the user to the database
	status, err := a.Deps.DAL.UserDAL.Add(&user)
	if err != nil {
		WriteErrorResponse(w, status, err)
		return
	}

	WriteJSONPayload(w, &ServerResponse{
		Message:    "User successfully registered",
		StatusCode: status,
		Payload:    user,
	})
}

//Login is the handler for the path /login
func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello you have requested: %s\n", r.URL.Path); err != nil {
		return
	}
}
